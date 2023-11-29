package rest

import (
	"github.com/Bit-Bridge-Source/BitBridge-RepoService-Go/internal/rest/router"
	"github.com/Bit-Bridge-Source/BitBridge-RepoService-Go/internal/rest/server"
	"github.com/gofiber/fiber/v2"
)

type FiberRouterAdapter struct {
	App *fiber.App
}

func (f *FiberRouterAdapter) GET(path string, handler router.HandlerFunc) {
	f.App.Get(path, func(ctx *fiber.Ctx) error {
		handler(&server.FiberContextAdapter{Ctx: ctx})
		return nil
	})
}

func (f *FiberRouterAdapter) POST(path string, handler router.HandlerFunc) {
	f.App.Post(path, func(ctx *fiber.Ctx) error {
		handler(&server.FiberContextAdapter{Ctx: ctx})
		return nil
	})
}

func (f *FiberRouterAdapter) PUT(path string, handler router.HandlerFunc) {
	f.App.Put(path, func(ctx *fiber.Ctx) error {
		handler(&server.FiberContextAdapter{Ctx: ctx})
		return nil
	})
}

func (f *FiberRouterAdapter) DELETE(path string, handler router.HandlerFunc) {
	f.App.Delete(path, func(ctx *fiber.Ctx) error {
		handler(&server.FiberContextAdapter{Ctx: ctx})
		return nil
	})
}
