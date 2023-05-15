package rabbit

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/streadway/amqp"
)

type mq struct {
	conn      *amqp.Connection
	url       string
	channel   *amqp.Channel
	consumers []*consumer
}

type consumer struct {
	queueName   string
	consumerTag string
	done        chan error
}

func NewMq() MQ {
	return &mq{}
}

func (m *mq) Connect(url string) error {
	var err error

	m.conn, err = amqp.Dial(url)
	if err != nil {
		return fmt.Errorf("failed to connect to RabbitMQ: %s", err)
	}

	m.channel, err = m.conn.Channel()
	if err != nil {
		return fmt.Errorf("failed to open a channel: %s", err)
	}

	m.url = url

	return nil
}

func (m *mq) CreateQueue(name string) error {
	_, err := m.channel.QueueDeclare(
		name,  // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare a queue: %s", err)
	}

	return nil
}

func (m *mq) Send(queueName string, message string) error {
	ch, err := m.conn.Channel()
	if err != nil {
		// if the channel is closed, try to reconnect to the server
		if err == amqp.ErrClosed {
			if err := m.Connect(m.url); err != nil {
				return err
			}
			ch, err = m.conn.Channel()
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}
	defer ch.Close()

	err = ch.Publish(
		"",
		queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (m *mq) Receive(queueName string) (string, error) {
	msgs, err := m.channel.Consume(
		queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return "", fmt.Errorf("failed to register a consumer: %s", err)
	}

	for msg := range msgs {
		return string(msg.Body), nil
	}

	return "", nil
}

func (m *mq) AddConsumer(num int, queueName string, consumerName string, callback func([]byte)) error {
	for i := 0; i < num; i++ {
		c := &consumer{
			queueName:   queueName,
			consumerTag: consumerName,
			done:        make(chan error),
		}

		delivery, err := m.channel.Consume(
			queueName,
			consumerName+uuid.NewString(),
			false,
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			return fmt.Errorf("failed to register a consumer: %s", err)
		}

		go func() {
			for d := range delivery {
				callback(d.Body)
				d.Ack(false)
			}
			c.done <- nil
		}()

		m.consumers = append(m.consumers, c)
	}

	return nil
}

func (m *mq) Close() error {
	if m.conn != nil {
		return m.conn.Close()
	}
	return nil
}
