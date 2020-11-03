package api

import (
	"github.com/GopherChat/gopher-server/module/context"
	"github.com/gofiber/fiber/v2"
)

type Context = context.Context

func (a *API) Handler(h func(*Context) error) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return h(context.FromRequest(c))
	}
}
