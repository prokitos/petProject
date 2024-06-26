package services

import (
	"encoding/json"
	"log"
	"math/rand"
	"module/internal/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/streadway/amqp"
)

func SendcarInsert(c *fiber.Ctx) error {

	var curCar models.CarToRM
	if err := c.BodyParser(&curCar); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	curCar.Types = "insert"
	DatabaseProducing(curCar)

	return c.SendStatus(fiber.StatusAccepted)
}

func SendcarShow(c *fiber.Ctx) error {

	var curCar models.CarToRM
	curCar.Id = c.Query("id", "")
	curCar.RegNum = c.Query("regNum")
	curCar.Mark = c.Query("mark")
	curCar.Model = c.Query("model")
	curCar.Year = c.Query("year")
	curCar.Owner.Name = c.Query("name")
	curCar.Owner.Surname = c.Query("surname")
	curCar.Owner.Patronymic = c.Query("patronymic")

	curCar.Types = "show"
	DatabaseProducing(curCar)

	return c.SendStatus(fiber.StatusAccepted)
}

func SendcarUpdate(c *fiber.Ctx) error {

	var curCar models.CarToRM
	if err := c.BodyParser(&curCar); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	curCar.Types = "update"
	DatabaseProducing(curCar)

	return c.SendStatus(fiber.StatusAccepted)
}

func SendcarDelete(c *fiber.Ctx) error {

	var curCar models.CarToRM
	curCar.Id = c.Query("id", "")

	curCar.Types = "delete"
	DatabaseProducing(curCar)

	return c.SendStatus(fiber.StatusAccepted)
}

func DatabaseProducing(curCar models.CarToRM) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	handleError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	amqpChannel, err := conn.Channel()
	handleError(err, "Can't create a amqpChannel")

	defer amqpChannel.Close()

	queue, err := amqpChannel.QueueDeclare("carService", true, false, false, false, nil)
	handleError(err, "Could not declare `database` queue")

	rand.Seed(time.Now().UnixNano())

	addTask := curCar
	body, err := json.Marshal(addTask)
	if err != nil {
		handleError(err, "Error encoding JSON")
	}

	err = amqpChannel.Publish("", queue.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         body,
	})

	if err != nil {
		log.Fatalf("Error publishing message: %s", err)
	}

	log.Printf("AddTask: %s %s", addTask.Id, addTask.Mark)
}

func handleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}

}
