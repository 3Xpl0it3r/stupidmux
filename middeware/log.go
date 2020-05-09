package middeware

import (
	"github.com/3Xpl0it3r/stupidmux/context"
)
import "github.com/sirupsen/logrus"


func LogMiddleWare(ctx *context.Context){
	logrus.WithField("method", ctx.Method).WithField("url", ctx.Request.URL.String()).Println("LogMiddleWare")
}