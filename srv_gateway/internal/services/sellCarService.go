package services

import (
	"context"
	"encoding/json"
	"module/internal/models"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	amqp "github.com/rabbitmq/amqp091-go"
	log "github.com/sirupsen/logrus"
)

func DatabaseSellProducing(c *fiber.Ctx, curSell models.SellingToRM) error {
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

	addTask := curSell
	body, err := json.Marshal(addTask)
	corrId := "67"

	err = ch.PublishWithContext(ctx,
		"",               // exchange
		"carSellService", // routing key
		false,            // mandatory
		false,            // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: corrId,
			ReplyTo:       q.Name,
			Body:          body,
		})
	handleError(err, "Failed to publish a message")

	for d := range msgs {
		if corrId == d.CorrelationId {

			addTasks := &models.ResponseSell{}
			json.Unmarshal(d.Body, addTasks)

			return returnSellResponse(c, addTasks)

		}
	}

	return models.ResponseBadRequest()
}

func returnSellResponse(c *fiber.Ctx, res *models.ResponseSell) error {

	if res.Description == models.ResponseSellGoodExecute().Description || res.Description == models.ResponseSellBadExecute().Description {
		return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"status": res.Description, "code": res.Code})
	}

	if res.Description == models.ResponseSellGoodShow([]models.Selling{}).Description || res.Description == models.ResponseSellBadShow().Description || models.ResponseSellNotForSale().Description == res.Description {
		return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"status": res.Description, "code": res.Code, "data": res.Sells})
	}
	return c.SendStatus(fiber.StatusAccepted)

}

func SendSellcarInsert(c *fiber.Ctx) error {

	var curSell models.SellingToRM

	if err := c.BodyParser(&curSell); err != nil {
		log.Debug("body parse error")
		return models.ResponseBadRequest()
	}

	curSell.Types = "SellInsert"

	return DatabaseSellProducing(c, curSell)
}

func SendSellcarShow(c *fiber.Ctx) error {

	var curSell models.SellingToRM

	curSell.Id, _ = strconv.Atoi(c.Query("id"))
	curSell.CarId, _ = strconv.Atoi(c.Query("car_id"))
	curSell.PeopleId, _ = strconv.Atoi(c.Query("car_id"))

	curSell.Types = "SellShow"
	return DatabaseSellProducing(c, curSell)
}

func SendSellcarUpdate(c *fiber.Ctx) error {

	var curSell models.SellingToRM

	if err := c.BodyParser(&curSell); err != nil {
		log.Debug("body parse error")
		return models.ResponseBadRequest()
	}

	curSell.Types = "SellUpdate"

	return DatabaseSellProducing(c, curSell)
}

func SendSellcarDelete(c *fiber.Ctx) error {

	var curSell models.SellingToRM

	curSell.Id, _ = strconv.Atoi(c.Query("id"))

	curSell.Types = "SellDelete"
	return DatabaseSellProducing(c, curSell)
}
