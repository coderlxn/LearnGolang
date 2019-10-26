package main

import (
	"flag"
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/streadway/amqp"
)

var (
	uri = flag.String("uri", "amqp://guest:guest@localhost:5672/", "AMQP URI")
	exchangeName = flag.String("exchange", "test-exchange", "Durable AMQP exchange name")
	exchangeType = flag.String("exchange-type", "direct", "Exchange type - direct|fanout|topic|x-custom")
	routingKey   = flag.String("key", "test-key", "AMQP routing key")
	body         = flag.String("body", "foobar", "Body of message")
	reliable     = flag.Bool("reliable", true, "Wait for the publisher confirmation before exiting")
)

func init()  {
	flag.Parse()
}

func confirmOne(confirms <-chan amqp.Confirmation)  {
	log.Printf("waiting for confirmation of one publishing")

	if confirmed := <-confirms; confirmed.Ack {
		log.Printf("confirmed delivery with delivery tag: %d", confirmed.DeliveryTag)
	} else {
		log.Printf("failed delivery of delivery tag: %d", confirmed.DeliveryTag)
	}
}

func publish(amqpURI, exchange, exchangeType, routingKey, body string, reliable bool) error  {
	log.Printf("dialing %q", amqpURI)
	connection, err := amqp.Dial(amqpURI)
	if err != nil {
		return fmt.Errorf("dial: %s", err)
	}
	defer connection.Close()

	log.Printf("got Connection, getting Channel")
	channel, err := connection.Channel()
	if err != nil {
		return fmt.Errorf("channel: %s", err)
	}

	log.Printf("got Channel, declaring %q Exchange (%q)", exchangeType, exchange)
	if err := channel.ExchangeDeclare(
		exchange,
		exchangeType,
		true,
		false,
		false,
		false,
		nil); err != nil {
			return fmt.Errorf("exchange declare: %s", err)
	}

	if reliable {
		log.Printf("enabling publishing confirms.")
		if err := channel.Confirm(false); err != nil {
			return fmt.Errorf("channel could not be put into confirm mode: %s", err)
		}

		confirms := channel.NotifyPublish(make(chan amqp.Confirmation, 1))
		defer confirmOne(confirms)
	}

	log.Printf("declared Exchange, publishing %dB body (%q)", len(body), body)
	if err = channel.Publish(
		exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			Headers:amqp.Table{},
			ContentType:"text/plain",
			ContentEncoding:"",
			Body:[]byte(body),
			DeliveryMode:amqp.Transient,
			Priority:0,
		},
		); err != nil {
			return fmt.Errorf("exchange Publish: %s", err)
	}

	return nil
}

func main()  {
	if err := publish(*uri, *exchangeName, *exchangeType, *routingKey, *body, *reliable); err != nil {
		log.Fatalf("%s", err)
	}
	log.Printf("published %dB OK", len(*body))
}