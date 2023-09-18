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
