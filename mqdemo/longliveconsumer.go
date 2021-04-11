package main

import (
	"flag"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"time"
)

var (
	uri          = flag.String("uri", "amqp://appmaker:maker123@192.168.1.121:5672/", "AMQP URI")
	exchange     = flag.String("exchange", "test-exchange", "Durable, non-auto-deleted AMQP exchange name")
	exchangeType = flag.String("exchange-type", "direct", "Exchange type - direct|fanout|topic|x-custom")
	queue        = flag.String("queue", "test-queue", "Ephemeral AMQP queue name")
	bindingKey   = flag.String("key", "test-key", "AMQP binding key")
	consumerTag  = flag.String("consumer-tag", "simple-consumer", "AMQP consumer tag (should not be blank)")
	lifetime     = flag.Duration("lifetime", 5* time.Second, "lifetime of process before shutdown (0s=infinite)")
)
var forever chan bool

func init()  {
	flag.Parse()
}

type Consumer struct {
	conn	*amqp.Connection
	channel *amqp.Channel
	tag		string
	done	chan error
}

func NewConsumer(amqpURI, exchange, exchangeType, queueName, key, ctag string) (*Consumer, error)  {
	c := &Consumer{
		conn:	nil,
		channel:nil,
		tag:	ctag,
		done:	make(chan  error),
	}

	var err error
	log.Printf("dialing %q", amqpURI)
	c.conn, err = amqp.Dial(amqpURI)
	if err != nil {
		return nil, fmt.Errorf("dial: %s", err)
	}

	go func() {
		fmt.Printf("closing: %s", <-c.conn.NotifyClose(make(chan *amqp.Error)))
	}()

	log.Printf("got Connection, getting Channel")
	c.channel, err = c.conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("channel: %s", err)
	}

	log.Printf("got Channel, declaring Exchange (%q)", exchange)
	if err = c.channel.ExchangeDeclare(
		exchange,
		exchangeType,
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		return nil, fmt.Errorf("exchange Declare: %s", err)
	}

	log.Printf("declared Exchange, declaring Queue %q", queueName)
	queue, err := c.channel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil)
	if err != nil {
		return nil, fmt.Errorf("queue Declare: %s", err)
	}

	log.Printf("declared Queue (%q %d message, %d consumers), binging to Exchange (key %q)",
		queue.Name, queue.Messages, queue.Consumers, key)
	if err = c.channel.QueueBind(
		queue.Name,
		key,
		exchange,
		false,
		nil); err != nil {
		return nil, fmt.Errorf("queue Bind: %s", err)
	}

	log.Printf("Queue bind to Exchanged, starting Consume (consumer tag %q)", c.tag)
	deliveries, err := c.channel.Consume(
		queue.Name,
		c.tag,
		false,
		false,
		false,
		false,
		nil,)
	if err != nil {
		return nil, fmt.Errorf("queue Consume: %s", err)
	}
	go handle(deliveries, c.done)

	return c, nil
}

func (c *Consumer) Shutdown() error {
	if err := c.channel.Cancel(c.tag, true); err != nil {
		return fmt.Errorf("consumer cancel failed: %s", err)
	}

	if err := c.conn.Close(); err != nil {
		return fmt.Errorf("AMQP connection close error: %s", err)
	}

	defer log.Printf("AMQP shutdown OK")

	return <-c.done
}

func handle(deliveries <-chan amqp.Delivery, done chan error)  {
	for d := range deliveries {
		log.Printf("got %dB delivery: [%v] %q", len(d.Body), d.DeliveryTag, d.Body)
		d.Ack(false)
	}
	log.Printf("handel: deliveries channel closed")
	forever <- true
	log.Printf("handel: assign forever value")
	done <- nil
	log.Printf("handel: finished")
}

func main()  {
	_, err := NewConsumer(*uri, *exchange, *exchangeType, *queue, *bindingKey, *consumerTag)
	if err != nil {
		log.Fatal("%s", err)
	}

	forever = make(chan bool)


	log.Println("program exited with: ", <- forever)

	//if *lifetime > 0 {
	//	log.Printf("running for %s", *lifetime)
	//	time.Sleep(*lifetime)
	//} else {
	//	log.Printf("running forever")
	//	select {
	//
	//	}
	//}
	//log.Printf("shutting down")
	//if err := c.Shutdown(); err != nil {
	//	log.Fatal("error during shutdown: %s", err)
	//}
}