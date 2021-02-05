package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

type Controller struct{}

func (c *Controller) Index(r *ghttp.Request) {
	r.Response.Write("this is index")
}

func (c *Controller) Show(r *ghttp.Request) {
	r.Response.Write("this is show")
}

type Order struct{}

func (o *Order) List(r *ghttp.Request) {
	r.Response.Write("list")
}

func (o *Order) Post(r *ghttp.Request)  {
	r.Response.Write("Post")
}

func main() {
	s := g.Server()
	// 函数注册
	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.Write("hello world")
	})

	// 对象注册
	s.BindObject("/obj", new(Controller))
	s.BindObject("/{.struct}-{.method}", new(Order))
	// 	REST	ful对象注册
	s.BindObjectRest("/order", new(Order))

	// 分组路由
	group := s.Group("/group")
	group.GET("/order", func(r *ghttp.Request){
		r.Response.Write("this is get method")
	})


	s.SetPort(8081)
	s.Run()
}
