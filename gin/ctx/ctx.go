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

		e.(*Flow).ctx = c

		e.Do()
	}
}
