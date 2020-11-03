package api

import "github.com/gofiber/fiber/v2"

type API struct {
	// logger *glog.Logger
	// config config.Config
}

func Init(r fiber.Router) *API {
	api := &API{}

	root := r.Group("/v1")

	api.UsersRouter(root.Group("/users"))

	return api
}
