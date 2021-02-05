package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

const (
	TraceIdName = "trace-id"
)

func main() {
	s := g.Server()
	s.Group("/group", func(group *ghttp.RouterGroup) {
		group.Middleware(func(r *ghttp.Request) {
			r.SetCtxVar(TraceIdName, "123123123")
			r.Middleware.Next()
		})
		group.ALL("/", func(r *ghttp.Request) {
			r.Response.Write(r.GetCtxVar(TraceIdName))
		})
	})
	s.SetPort(8091)
	s.Run()
}
