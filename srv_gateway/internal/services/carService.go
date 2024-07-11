package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"module/internal/generpc"
	"module/internal/models"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
)

func registerSend(c *fiber.Ctx, car models.CarToRM) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	conn, err := grpc.Dial("localhost:8006", grpc.WithInsecure())
	if err != nil {
		return c.SendString("connecting error")
	}
	defer conn.Close()
	client := generpc.NewEnrichmentClient(conn)

	// var sendedData *generpc.CarRequest
	// sendedData.Mark = car.Mark
	// sendedData.Year = car.Year
	// sendedData.Price = int64(car.Price)
	// sendedData.Color = car.Color
	// sendedData.MaxSpeed = int64(car.MaxSpeed)
	// sendedData.SeatsNum = int64(car.SeatsNum)
	// sendedData.Status = car.Status

	response, err := client.CarEnricht(ctx, &generpc.CarRequest{})
	if err != nil {
		fmt.Println(err)
		fmt.Println("too long. context time expired. more than 1 second.")
		return c.SendString("long execution")
	}

	fmt.Println(response.Color)

	return c.SendString("uraa")
}

func SendcarInsert(c *fiber.Ctx) error {

	var curCar models.CarToRM

	if err := c.BodyParser(&curCar); err != nil {
		return models.ResponseBadRequest()
	}

	curCar.Types = "insert"
	fmt.Println(curCar)

	return registerSend(c, curCar)
	// return DatabaseProducing(c, curCar)
}

func SendcarShow(c *fiber.Ctx) error {

	var curCar models.CarToRM

	curCar.Id, _ = strconv.Atoi(c.Query("id"))
	curCar.Mark = c.Query("mark", "")
	curCar.Year = c.Query("year", "")
	curCar.Price, _ = strconv.Atoi(c.Query("price", ""))
	curCar.Color = c.Query("color", "")
	curCar.MaxSpeed, _ = strconv.Atoi(c.Query("max_speed", ""))
	curCar.SeatsNum, _ = strconv.Atoi(c.Query("seat_num", ""))
	curCar.Status = c.Query("status", "")

	curCar.Types = "show"
	return DatabaseProducing(c, curCar)
}

func SendcarUpdate(c *fiber.Ctx) error {

	var curCar models.CarToRM

	if err := c.BodyParser(&curCar); err != nil {
		return models.ResponseBadRequest()
	}

	curCar.Types = "update"
	fmt.Println(curCar)

	return DatabaseProducing(c, curCar)
}

func SendcarDelete(c *fiber.Ctx) error {

	var curCar models.CarToRM

	curCar.Id, _ = strconv.Atoi(c.Query("id"))
	curCar.Mark = c.Query("mark", "")
	curCar.Year = c.Query("year", "")
	curCar.Price, _ = strconv.Atoi(c.Query("price", ""))
	curCar.Color = c.Query("color", "")
	curCar.MaxSpeed, _ = strconv.Atoi(c.Query("max_speed", ""))
	curCar.SeatsNum, _ = strconv.Atoi(c.Query("seat_num", ""))
	curCar.Status = c.Query("status", "")

	curCar.Types = "delete"
	return DatabaseProducing(c, curCar)
}

func DatabaseProducing(c *fiber.Ctx, curCar models.CarToRM) error {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	handleError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	handleError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"",    // name
		true,  // durable
		false, // delete when unused
		true,  // exclusive
		false, // noWait
		nil,   // arguments
	)
	handleError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	handleError(err, "Failed to register a consumer")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	addTask := curCar
	body, err := json.Marshal(addTask)

	corrId := "66"

	err = ch.PublishWithContext(ctx,
		"",           // exchange
		"carService", // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: corrId,
			ReplyTo:       q.Name,
			Body:          body,
		})
	handleError(err, "Failed to publish a message")

	for d := range msgs {
		if corrId == d.CorrelationId {

			addTasks := &models.ResponseCar{}
			json.Unmarshal(d.Body, addTasks)

			return returnResponse(c, addTasks)

		}
	}

	return models.ResponseBadRequest()
}

func returnResponse(c *fiber.Ctx, res *models.ResponseCar) error {

	if res.Description == models.ResponseCarGoodCreate().Description || res.Description == models.ResponseCarBadCreate().Description {
		return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"status": res.Description, "code": res.Code})
	}

	if res.Description == models.ResponseCarGoodShow([]models.Car{}).Description {
		return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"status": res.Description, "code": res.Code, "data": res.Cars})
	}
	if res.Description == models.ResponseCarBadShow().Description {
		return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"status": res.Description, "code": res.Code})
	}

	if res.Description == models.ResponseCarBadDelete().Description || res.Description == models.ResponseCarGoodDelete().Description {
		return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"status": res.Description, "code": res.Code})
	}

	if res.Description == models.ResponseCarBadUpdate().Description || res.Description == models.ResponseCarGoodUpdate().Description {
		return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"status": res.Description, "code": res.Code})
	}

	return c.SendStatus(fiber.StatusAccepted)

}

func handleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}

}
