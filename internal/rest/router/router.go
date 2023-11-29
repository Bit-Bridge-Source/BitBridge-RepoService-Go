package router

import "github.com/Bit-Bridge-Source/BitBridge-RepoService-Go/internal/rest/server"

type HandlerFunc func(ctx server.HTTPContext)

type Router interface {
	GET(path string, handler HandlerFunc)
	POST(path string, handler HandlerFunc)
	PUT(path string, handler HandlerFunc)
	DELETE(path string, handler HandlerFunc)
}
