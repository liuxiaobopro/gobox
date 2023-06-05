package ctx

import (
	"encoding/json"
	"reflect"

	"github.com/gin-gonic/gin"
)

func Use(e ICtx) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctl := slave(e).(ICtx)
		initFlow(c, ctl)

		if err := ctl.FlowHandle(); err != nil {
			ctl.ReturnJson(err)
			return
		}
		if err := ctl.FlowValidate(); err != nil {
			ctl.ReturnJson(err)
			return
		}
		if err := ctl.FlowLogic(); err != nil {
			ctl.ReturnJson(err)
			return
		}
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
