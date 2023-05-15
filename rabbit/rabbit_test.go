package rabbit

import (
	"testing"
	"time"
)

func Test_mq_Connect(t *testing.T) {
	var mqClient = NewMq()
	if err := mqClient.Connect("amqp://admin:admin@192.168.3.74:5672/"); err != nil {
		t.Error("Connect error")
	}

	queueName := "test"

	if err := mqClient.CreateQueue(queueName); err != nil {
		t.Error("CreateQueue error")
		return
	}

	go func() {
		if err := mqClient.AddConsumer(10, queueName, "lxb1", func(data []byte) {
			t.Log("Receive success, data:", string(data))
		}); err != nil {
			t.Errorf("AddConsumer error, err: %s", err)
		}
	}()

	for i := 0; i < 50; i++ {
		if err := mqClient.Send(queueName, "liuxiaobo"); err != nil {
			t.Errorf("Send error, err: %s", err)
			return
		}

		t.Log("Send success, data: liuxiaobo")

		time.Sleep(2 * time.Second)
	}

	if err := mqClient.Close(); err != nil {
		t.Error("Close error")
	}
}
