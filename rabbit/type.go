package rabbit

type MQ interface {
	Connect(url string) error
	CreateQueue(name string) error
	AddConsumer(num int, queueName string, consumerName string, callback func([]byte)) error
	Send(queueName string, message string) error
	Receive(queueName string) (string, error)
	Close() error
}
