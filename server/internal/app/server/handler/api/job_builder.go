package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/pouyam79i/simple_quera/server/internal/app/server/config"
	amqp "github.com/rabbitmq/amqp091-go"
)

// TODO: send msg to jb properly
// TODO: connect once!
func SendDataToJobBuilder(data config.ClientCode) bool {

	queryParams := url.Values{
		"code":     {data.CodeX.Code},
		"language": {data.CodeX.Language},
		"input":    {data.CodeX.Input},
	}
	payload := []byte(queryParams.Encode())

	jobMSG := config.JB_MSG{
		Data:  string(payload),
		Token: data.Token,
	}

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	conn.IsClosed()
	if err != nil {
		return false
	}

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()
	if err != nil {
		return false
	}

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err != nil {
		return false
	}

	body, err := json.Marshal(jobMSG)
	failOnError(err, "Failed to build json from struct")
	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", body)
	return err == nil
}

// TODO: tell user that you failed to complete requested file execution!
func failOnError(err error, msg string) {
	if err != nil {
		// log.Panicf("%s: %s", msg, err)
		fmt.Println("Failed to send msg!,\nmsg: ", msg, "\nerr: ", err)
	}
}
