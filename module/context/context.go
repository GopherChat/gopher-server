package context

import (
	"github.com/GopherChat/gopher-server/app"
	"github.com/GopherChat/gopher-server/model"
	"github.com/GopherChat/gopher-server/module/klog"
	"github.com/gofiber/fiber/v2"
)

var ctxKey string = "context"

type Context struct {
	*fiber.Ctx

	app    app.App
	logger *klog.Logger

	IsSigned bool
	User     *model.User
}

func FromRequest(c *fiber.Ctx) *Context {
	ctx, _ := c.Locals(ctxKey).(*Context)
	return ctx
}

func Contexter(a app.App, l *klog.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := &Context{
			Ctx:    c,
			app:    a,
			logger: l,
		}

		c.Locals(ctxKey, ctx)

		return c.Next()
	}
}
