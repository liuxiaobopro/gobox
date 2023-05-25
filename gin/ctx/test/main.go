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
		Code: 0,
		Msg:  d.GetAuthor(),
	})
}

func main() {
	r := gin.Default()

	r.GET("/ping", ctx.Use(new(demo)))

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
