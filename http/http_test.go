package http

import (
	"testing"
)

func TestClient_Get(t *testing.T) {
	client := Client{
		Url: "http://localhost:8080/goodsInfo",
	}

	resp, err := client.Get()
	if err != nil {
		t.Error(err)
	}

	t.Log(string(resp))
}

func TestClient_Post(t *testing.T) {
	// TODO: test post
}
