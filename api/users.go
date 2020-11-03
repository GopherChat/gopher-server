package api

import "github.com/gofiber/fiber/v2"

func (a *API) UsersRouter(r fiber.Router) {
	r.Post("/signup", a.Handler(register))
}

func register(ctx *Context) error {
	// bind request

	// call app

	// return result
	return ctx.SendStatus(fiber.StatusCreated)
}
