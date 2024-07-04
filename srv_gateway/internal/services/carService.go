package services

import (
	"context"
	"encoding/json"
	"log"
	"module/internal/models"
	"time"

	"github.com/gofiber/fiber/v2"
	amqp "github.com/rabbitmq/amqp091-go"
)

func firstTest() models.CarToRM {

	var curCar models.CarToRM
	var NewDevices []models.AdditionalDevices
	var dev1 models.AdditionalDevices
	var dev2 models.AdditionalDevices
	dev1.DeviceName = "ParkMaster"
	dev2.DeviceName = "DVR"
	NewDevices = append(NewDevices, dev1)
	NewDevices = append(NewDevices, dev2)

	var CarEngine models.CarEngine
	CarEngine.EngineCapacity = 450
	CarEngine.EnginePower = 120

	var Owner models.People
	Owner.Email = "owenrMail"
	Owner.Name = "ivan"
	Owner.Surname = "zhilov"

	var AllOwners []models.People
	var own1 models.People
	var own2 models.People
	own1.Email = "firster"
	own1.Name = "nikita"
	own1.Surname = "nikitov"
	own2.Email = "tesst"
	own2.Name = "jora"
	own2.Surname = "sinikov"
	AllOwners = append(AllOwners, own1)
	AllOwners = append(AllOwners, own2)

	curCar.Mark = "Lada"
	curCar.Year = "2010"
	curCar.Price = 2450000
	curCar.Color = "red"
	curCar.MaxSpeed = 140
	curCar.SeatsNum = 4
	curCar.Engine = CarEngine
	curCar.Devices = NewDevices
	curCar.OwnerList = AllOwners

	return curCar
}

func secondTest() models.CarToRM {

	var curCar models.CarToRM
	var NewDevices []models.AdditionalDevices
	var dev1 models.AdditionalDevices
	var dev2 models.AdditionalDevices
	dev1.DeviceName = "Fignya"
	dev2.DeviceName = "DVR"
	NewDevices = append(NewDevices, dev1)
	NewDevices = append(NewDevices, dev2)

	var CarEngine models.CarEngine
	CarEngine.EngineCapacity = 999
	CarEngine.EnginePower = 111

	var Owner models.People
	Owner.Email = "sdg436t"
	Owner.Name = "shulua"
	Owner.Surname = "dashokv"

	var AllOwners []models.People
	var own1 models.People
	var own2 models.People
	own1.Email = "qqwe"
	own1.Name = "qwrqw"
	own1.Surname = "qwwqr"
	own2.Email = "poipoi"
	own2.Name = "poipio"
	own2.Surname = "oipipi"
	AllOwners = append(AllOwners, own1)
	AllOwners = append(AllOwners, own2)

	curCar.Mark = "Ladina"
	curCar.Year = "2011"
	curCar.Price = 2450000
	curCar.Color = "reds"
	curCar.MaxSpeed = 166
	curCar.SeatsNum = 4
	curCar.Engine = CarEngine
	curCar.Devices = NewDevices
	curCar.OwnerList = AllOwners

	return curCar
}

func SendcarInsert(c *fiber.Ctx) error {

	var curCar models.CarToRM

	// тестовые данные
	curCar = firstTest()

	curCar.Types = "insert"
	return DatabaseProducing(c, curCar)
}

func SendcarShow(c *fiber.Ctx) error {

	var curCar models.CarToRM

	// тестовые данные
	curCar.Mark = "Lada"

	curCar.Types = "show"
	return DatabaseProducing(c, curCar)
}

func SendcarUpdate(c *fiber.Ctx) error {

	var curCar models.CarToRM

	// тестовые данные
	curCar.Color = "green"
	curCar.MaxSpeed = 155
	curCar.Id = 2

	curCar.Types = "update"
	return DatabaseProducing(c, curCar)
}

func SendcarDelete(c *fiber.Ctx) error {

	var curCar models.CarToRM

	// тестовые данные
	curCar.Color = "red"
	curCar.MaxSpeed = 140
	curCar.Id = 3

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

	if res.Description == models.ResponseCarGoodCreate().Description {
		return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"status": res.Description, "code": res.Code})
	}
	if res.Description == models.ResponseCarBadCreate().Description {
		return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"status": res.Description, "code": res.Code})
	}

	if res.Description == models.ResponseCarGoodShow([]models.Car{}).Description {
		return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"status": res.Description, "code": res.Code, "data": res.Cars})
	}
	if res.Description == models.ResponseCarBadShow().Description {
		return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"status": res.Description, "code": res.Code})
	}

	if res.Description == models.ResponseCarBadDelete().Description {
		return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"status": res.Description, "code": res.Code})
	}
	if res.Description == models.ResponseCarGoodDelete().Description {
		return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"status": res.Description, "code": res.Code})
	}

	if res.Description == models.ResponseCarBadUpdate().Description {
		return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"status": res.Description, "code": res.Code})
	}
	if res.Description == models.ResponseCarGoodUpdate().Description {
		return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"status": res.Description, "code": res.Code})
	}

	return c.SendStatus(fiber.StatusAccepted)

}

func handleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}

}
