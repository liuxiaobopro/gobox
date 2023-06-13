package main

import (
	"fmt"

	"github.com/liuxiaobopro/gobox/gin/ctx"
	"github.com/liuxiaobopro/gobox/reply"

	"github.com/gin-gonic/gin"
)

type demo struct {
	ctx.Flow
}

func (d *demo) Handle() {
	fmt.Println("handle")
	d.LogInfof("handle")
}

func (d *demo) Validate() {
	fmt.Println("validate")
	d.LogInfof("validate")
}

func (d *demo) Logic() {
	fmt.Println("logic")
	d.LogInfof("logic")

	d.ReturnJson(&reply.T{
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
