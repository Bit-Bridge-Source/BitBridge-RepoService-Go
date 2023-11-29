package rest

import "github.com/gofiber/fiber/v2"

type HTTPContext interface {
	GetParam(key string) string
	BindJSON(obj interface{}) error
	JSON(code int, obj interface{})
}

type FiberContextAdapter struct {
	Ctx *fiber.Ctx
}

func (f *FiberContextAdapter) GetParam(key string) string {
	return f.Ctx.Params(key)
}

func (f *FiberContextAdapter) BindJSON(obj interface{}) error {
	return f.Ctx.BodyParser(obj)
}

func (f *FiberContextAdapter) JSON(code int, obj interface{}) {
	f.Ctx.Status(code).JSON(obj)
}
