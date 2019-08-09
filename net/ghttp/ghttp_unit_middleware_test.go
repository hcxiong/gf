// Copyright 2018 gf Author(https://github.com/gogf/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package ghttp_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/test/gtest"
)

func Test_BindMiddleware_Basic1(t *testing.T) {
	p := ports.PopRand()
	s := g.Server(p)
	s.BindHandler("/test/test", func(r *ghttp.Request) {
		r.Response.Write("test")
	})
	s.BindMiddleware("/test", func(r *ghttp.Request) {
		r.Response.Write("1")
		r.Middleware.Next()
		r.Response.Write("2")
	}, func(r *ghttp.Request) {
		r.Response.Write("3")
		r.Middleware.Next()
		r.Response.Write("4")
	})
	s.BindMiddleware("/test/:name", func(r *ghttp.Request) {
		r.Response.Write("5")
		r.Middleware.Next()
		r.Response.Write("6")
	}, func(r *ghttp.Request) {
		r.Response.Write("7")
		r.Middleware.Next()
		r.Response.Write("8")
	})
	s.SetPort(p)
	s.SetDumpRouteMap(false)
	s.Start()
	defer s.Shutdown()

	// 等待启动完成
	time.Sleep(200 * time.Millisecond)
	gtest.Case(t, func() {
		client := ghttp.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		gtest.Assert(client.GetContent("/"), "Not Found")
		gtest.Assert(client.GetContent("/test"), "1342")
		gtest.Assert(client.GetContent("/test/test"), "57test86")
	})
}

func Test_BindMiddleware_Basic2(t *testing.T) {
	p := ports.PopRand()
	s := g.Server(p)
	s.BindHandler("/test/test", func(r *ghttp.Request) {
		r.Response.Write("test")
	})
	s.BindMiddleware("PUT:/test", func(r *ghttp.Request) {
		r.Response.Write("1")
		r.Middleware.Next()
		r.Response.Write("2")
	}, func(r *ghttp.Request) {
		r.Response.Write("3")
		r.Middleware.Next()
		r.Response.Write("4")
	})
	s.BindMiddleware("POST:/test/:name", func(r *ghttp.Request) {
		r.Response.Write("5")
		r.Middleware.Next()
		r.Response.Write("6")
	}, func(r *ghttp.Request) {
		r.Response.Write("7")
		r.Middleware.Next()
		r.Response.Write("8")
	})
	s.SetPort(p)
	s.SetDumpRouteMap(false)
	s.Start()
	defer s.Shutdown()

	// 等待启动完成
	time.Sleep(200 * time.Millisecond)
	gtest.Case(t, func() {
		client := ghttp.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		gtest.Assert(client.GetContent("/"), "Not Found")
		gtest.Assert(client.GetContent("/test"), "Not Found")
		gtest.Assert(client.PutContent("/test"), "1342")
		gtest.Assert(client.PostContent("/test"), "Not Found")
		gtest.Assert(client.GetContent("/test/test"), "test")
		gtest.Assert(client.PutContent("/test/test"), "test")
		gtest.Assert(client.PostContent("/test/test"), "57test86")
	})
}

func Test_AddMiddleware_Basic1(t *testing.T) {
	p := ports.PopRand()
	s := g.Server(p)
	s.BindHandler("/test/test", func(r *ghttp.Request) {
		r.Response.Write("test")
	})
	s.AddMiddleware(func(r *ghttp.Request) {
		r.Response.Write("1")
		r.Middleware.Next()
		r.Response.Write("2")
	})
	s.AddMiddleware(func(r *ghttp.Request) {
		r.Response.Write("3")
		r.Middleware.Next()
		r.Response.Write("4")
	})
	s.SetPort(p)
	s.SetDumpRouteMap(false)
	s.Start()
	defer s.Shutdown()

	// 等待启动完成
	time.Sleep(200 * time.Millisecond)
	gtest.Case(t, func() {
		client := ghttp.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		gtest.Assert(client.GetContent("/"), "1342")
		gtest.Assert(client.GetContent("/test/test"), "13test42")
	})
}

func Test_AddMiddleware_Basic2(t *testing.T) {
	p := ports.PopRand()
	s := g.Server(p)
	s.BindHandler("PUT:/test/test", func(r *ghttp.Request) {
		r.Response.Write("test")
	})
	s.AddMiddleware(func(r *ghttp.Request) {
		r.Response.Write("1")
		r.Middleware.Next()
		r.Response.Write("2")
	})
	s.AddMiddleware(func(r *ghttp.Request) {
		r.Response.Write("3")
		r.Middleware.Next()
		r.Response.Write("4")
	})
	s.SetPort(p)
	s.SetDumpRouteMap(false)
	s.Start()
	defer s.Shutdown()

	// 等待启动完成
	time.Sleep(200 * time.Millisecond)
	gtest.Case(t, func() {
		client := ghttp.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		gtest.Assert(client.GetContent("/"), "1342")
		gtest.Assert(client.PutContent("/"), "1342")
		gtest.Assert(client.GetContent("/test/test"), "1342")
		gtest.Assert(client.PutContent("/test/test"), "13test42")
	})
}

func Test_AddMiddleware_Basic3(t *testing.T) {
	p := ports.PopRand()
	s := g.Server(p)
	s.BindHandler("/test/test", func(r *ghttp.Request) {
		r.Response.Write("test")
	})
	s.AddMiddleware(func(r *ghttp.Request) {
		r.Response.Write("1")
		r.Middleware.Next()
	})
	s.AddMiddleware(func(r *ghttp.Request) {
		r.Middleware.Next()
		r.Response.Write("2")
	})
	s.SetPort(p)
	s.SetDumpRouteMap(false)
	s.Start()
	defer s.Shutdown()

	// 等待启动完成
	time.Sleep(200 * time.Millisecond)
	gtest.Case(t, func() {
		client := ghttp.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		gtest.Assert(client.GetContent("/"), "12")
		gtest.Assert(client.GetContent("/test/test"), "1test2")
	})
}

func Test_AddMiddleware_Basic4(t *testing.T) {
	p := ports.PopRand()
	s := g.Server(p)
	s.BindHandler("/test/test", func(r *ghttp.Request) {
		r.Response.Write("test")
	})
	s.AddMiddleware(func(r *ghttp.Request) {
		r.Middleware.Next()
		r.Response.Write("1")
	})
	s.AddMiddleware(func(r *ghttp.Request) {
		r.Response.Write("2")
		r.Middleware.Next()
	})
	s.SetPort(p)
	s.SetDumpRouteMap(false)
	s.Start()
	defer s.Shutdown()

	// 等待启动完成
	time.Sleep(200 * time.Millisecond)
	gtest.Case(t, func() {
		client := ghttp.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		gtest.Assert(client.GetContent("/"), "21")
		gtest.Assert(client.GetContent("/test/test"), "2test1")
	})
}

func Test_AddMiddleware_Basic5(t *testing.T) {
	p := ports.PopRand()
	s := g.Server(p)
	s.BindHandler("/test/test", func(r *ghttp.Request) {
		r.Response.Write("test")
	})
	s.AddMiddleware(func(r *ghttp.Request) {
		r.Response.Write("1")
		r.Middleware.Next()
	})
	s.AddMiddleware(func(r *ghttp.Request) {
		r.Response.Write("2")
		r.Middleware.Next()
	})
	s.SetPort(p)
	s.SetDumpRouteMap(false)
	s.Start()
	defer s.Shutdown()

	// 等待启动完成
	time.Sleep(200 * time.Millisecond)
	gtest.Case(t, func() {
		client := ghttp.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		gtest.Assert(client.GetContent("/"), "12")
		gtest.Assert(client.GetContent("/test/test"), "12test")
	})
}

type ObjectMiddleware struct{}

func (o *ObjectMiddleware) Init(r *ghttp.Request) {
	r.Response.Write("100")
}

func (o *ObjectMiddleware) Shut(r *ghttp.Request) {
	r.Response.Write("200")
}

func (o *ObjectMiddleware) Index(r *ghttp.Request) {
	r.Response.Write("Object Index")
}

func (o *ObjectMiddleware) Show(r *ghttp.Request) {
	r.Response.Write("Object Show")
}

func (o *ObjectMiddleware) Info(r *ghttp.Request) {
	r.Response.Write("Object Info")
}

func Test_AddMiddleware_Basic6(t *testing.T) {
	p := ports.PopRand()
	s := g.Server(p)
	s.BindObject("/", new(ObjectMiddleware))
	s.AddMiddleware(func(r *ghttp.Request) {
		r.Response.Write("1")
		r.Middleware.Next()
		r.Response.Write("2")
	})
	s.AddMiddleware(func(r *ghttp.Request) {
		r.Response.Write("3")
		r.Middleware.Next()
		r.Response.Write("4")
	})
	s.SetPort(p)
	s.SetDumpRouteMap(false)
	s.Start()
	defer s.Shutdown()

	// 等待启动完成
	time.Sleep(200 * time.Millisecond)
	gtest.Case(t, func() {
		client := ghttp.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		gtest.Assert(client.GetContent("/"), "13100Object Index20042")
		gtest.Assert(client.GetContent("/init"), "1342")
		gtest.Assert(client.GetContent("/shut"), "1342")
		gtest.Assert(client.GetContent("/index"), "13100Object Index20042")
		gtest.Assert(client.GetContent("/show"), "13100Object Show20042")
		gtest.Assert(client.GetContent("/none-exist"), "1342")
	})
}

func Test_Hook_Middleware_Basic1(t *testing.T) {
	p := ports.PopRand()
	s := g.Server(p)
	s.BindHandler("/test/test", func(r *ghttp.Request) {
		r.Response.Write("test")
	})
	s.BindHookHandler("/*", ghttp.HOOK_BEFORE_SERVE, func(r *ghttp.Request) {
		r.Response.Write("a")
	})
	s.BindHookHandler("/*", ghttp.HOOK_AFTER_SERVE, func(r *ghttp.Request) {
		r.Response.Write("b")
	})
	s.BindHookHandler("/*", ghttp.HOOK_BEFORE_SERVE, func(r *ghttp.Request) {
		r.Response.Write("c")
	})
	s.BindHookHandler("/*", ghttp.HOOK_AFTER_SERVE, func(r *ghttp.Request) {
		r.Response.Write("d")
	})
	s.AddMiddleware(func(r *ghttp.Request) {
		r.Response.Write("1")
		r.Middleware.Next()
		r.Response.Write("2")
	})
	s.AddMiddleware(func(r *ghttp.Request) {
		r.Response.Write("3")
		r.Middleware.Next()
		r.Response.Write("4")
	})
	s.SetPort(p)
	s.SetDumpRouteMap(false)
	s.Start()
	defer s.Shutdown()

	// 等待启动完成
	time.Sleep(200 * time.Millisecond)
	gtest.Case(t, func() {
		client := ghttp.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		gtest.Assert(client.GetContent("/"), "ac1342bd")
		gtest.Assert(client.GetContent("/test/test"), "ac13test42bd")
	})
}