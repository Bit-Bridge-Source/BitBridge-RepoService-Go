package rest

import "github.com/gofiber/fiber/v2"

type FiberRouterAdapter struct {
	App *fiber.App
}

func (f *FiberRouterAdapter) GET(path string, handler HandlerFunc) {
	f.App.Get(path, func(ctx *fiber.Ctx) error {
		handler(&FiberContextAdapter{Ctx: ctx})
		return nil
	})
}

func (f *FiberRouterAdapter) POST(path string, handler HandlerFunc) {
	f.App.Post(path, func(ctx *fiber.Ctx) error {
		handler(&FiberContextAdapter{Ctx: ctx})
		return nil
	})
}

func (f *FiberRouterAdapter) PUT(path string, handler HandlerFunc) {
	f.App.Put(path, func(ctx *fiber.Ctx) error {
		handler(&FiberContextAdapter{Ctx: ctx})
		return nil
	})
}

func (f *FiberRouterAdapter) DELETE(path string, handler HandlerFunc) {
	f.App.Delete(path, func(ctx *fiber.Ctx) error {
		handler(&FiberContextAdapter{Ctx: ctx})
		return nil
	})
}
