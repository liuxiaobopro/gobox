package ctx

import (
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	replyx "github.com/liuxiaobopro/gobox/reply"
)

func Use(e ICtx) gin.HandlerFunc {
	return func(c *gin.Context) {
		//TODO: panic recover
		defer func() {
			if err := recover(); err != nil {
				e.PrintErrorf("flow panic: %v", err)
				e.SetHttpCode(http.StatusInternalServerError)
				e.ReturnJson(replyx.InternalErrT)
				c.Abort()
			}
		}()

		ctl := slave(e).(ICtx)
		initFlow(c, ctl)

		ctl.Handle()
		ctl.Validate()
		ctl.Logic()
	}
}

func slave(src interface{}) interface{} {
	typ := reflect.TypeOf(src)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
		dst := reflect.New(typ).Elem()
		b, _ := json.Marshal(src)
		json.Unmarshal(b, dst.Addr().Interface())
		return dst.Addr().Interface()
	} else {
		dst := reflect.New(typ).Elem()
		b, _ := json.Marshal(src)
		json.Unmarshal(b, dst.Addr().Interface())
		return dst.Interface()
	}
}

func initFlow(ctx *gin.Context, flow ICtx) {
	flow.setCtx(ctx)
	flow.initLog()
}
