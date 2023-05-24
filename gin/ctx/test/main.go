package main

import (
	"github.com/liuxiaobopro/gobox/gin/ctx"
	"github.com/liuxiaobopro/gobox/reply"

	"github.com/gin-gonic/gin"
)

type demo struct {
	ctx.Flow
}

func (d *demo) Do() {
	d.ReturnJson(200, &reply.T{
		Code: 66,
		Msg:  "hello world",
	})
}

func main() {
	r := gin.Default()

	var d = &demo{}
	r.GET("/ping", ctx.Use(d))

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
