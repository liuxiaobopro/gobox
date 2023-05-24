package ctx

import (
	"github.com/gin-gonic/gin"
)

func Use(e CtxInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				e.Error(err.(error))
			}
		}()

		// q: interface conversion: ctx.CtxInterface is *main.demo, not *ctx.Flow
		// a: e.(*Flow).ctx = c
		// q: 不对
		// a: 对
		// q: 为什么
		// a: 因为 e 是 *demo 类型，而 *demo 类型实现了 CtxInterface 接口，所以 e 是 CtxInterface 类型

		e.(*Flow).ctx = c

		e.Do()
	}
}
