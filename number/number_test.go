package number

import (
	"testing"
	"time"
)

func TestRandomInt(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(RandomInt(10, 10000))

		time.Sleep(time.Second)
	}
}

func TestRandomFloat(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(RandomFloat(-0.1, 0.1))
		t.Log(RandomFloat(-0.3, 0.3))

		time.Sleep(time.Second)
	}
}
