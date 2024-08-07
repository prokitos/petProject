package server

import (
	"context"
	"encoding/json"
	"module/internal/models"
	"module/internal/services"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	log "github.com/sirupsen/logrus"
)

var RMQaddress string = ""

func handleError(err error, msg string) {
	if err != nil {
		log.Error(msg, err)
	}

}

func CarConsuming() {

	conn, err := amqp.Dial("amqp://guest:guest@" + RMQaddress + "/")
	handleError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	handleError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"carService", // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	handleError(err, "Failed to declare a queue")

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	handleError(err, "Failed to set QoS")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	handleError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		for d := range msgs {

			addTask := &models.CarToRM{}
			err := json.Unmarshal(d.Body, addTask)
			if err != nil {
				log.Error("unmarshal error")
				return
			}

			resp := GatewayCar(addTask)
			temp, err := json.Marshal(resp)

			err = ch.PublishWithContext(ctx,
				"",        // exchange
				d.ReplyTo, // routing key
				false,     // mandatory
				false,     // immediate
				amqp.Publishing{
					ContentType:   "text/plain",
					CorrelationId: d.CorrelationId,
					Body:          temp,
				})
			handleError(err, "Failed to publish a message")

			d.Ack(false)
		}
	}()

	log.Printf(" [*] Awaiting RPC requests")
	<-forever
}

func GatewayCar(car *models.CarToRM) models.ResponseCar {

	types := car.Types

	var tempCar models.Car
	tempCar = car.Car

	switch types {
	case "insert":
		return services.CarInsert(tempCar)
	case "delete":
		return services.CarDelete(tempCar)
	case "update":
		return services.CarUpdate(tempCar)
	case "show":
		return services.CarShow(tempCar)
	}

	return models.ResponseCarUnsupported()
}

func CarSellConsuming() {

	conn, err := amqp.Dial("amqp://guest:guest@" + RMQaddress + "/")
	log.Info(RMQaddress)
	handleError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	handleError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"carSellService", // name
		true,             // durable
		false,            // delete when unused
		false,            // exclusive
		false,            // no-wait
		nil,              // arguments
	)
	handleError(err, "Failed to declare a queue")

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	handleError(err, "Failed to set QoS")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	handleError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		for d := range msgs {

			addTask := &models.SellingToRM{}
			err := json.Unmarshal(d.Body, addTask)
			if err != nil {
				log.Error("unmarshal error")
				return
			}

			resp := GatewaySellCar(addTask)
			temp, err := json.Marshal(resp)

			err = ch.PublishWithContext(ctx,
				"",        // exchange
				d.ReplyTo, // routing key
				false,     // mandatory
				false,     // immediate
				amqp.Publishing{
					ContentType:   "text/plain",
					CorrelationId: d.CorrelationId,
					Body:          temp,
				})
			handleError(err, "Failed to publish a message")

			d.Ack(false)
		}
	}()

	log.Printf(" [*] Awaiting RPC requests")
	<-forever
}

func GatewaySellCar(sell *models.SellingToRM) models.ResponseSell {

	types := sell.Types

	switch types {
	case "SellInsert":
		return services.CarSellInsert(*sell)
	case "SellDelete":
		return services.CarSellDelete(*sell)
	case "SellUpdate":
		return services.CarSellUpdate(*sell)
	case "SellShow":
		return services.CarSellShow(*sell)
	}

	return models.ResponseSellBadExecute()
}
