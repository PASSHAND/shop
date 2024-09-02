package cmd

import (
	"context"
	"goframe-shop/internal/controller"
	"goframe-shop/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"goframe-shop/internal/controller/hello"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				//group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Middleware( //自定义中间件
					service.Middleware().Ctx,
					service.Middleware().ResponseHandler,
				)
				group.Bind(
					hello.NewV1(),
					controller.Rotation, //轮播图
					controller.Position, //手工位
					controller.Admin,    //管理员
					controller.Login,    //登陆
				)
			})
			s.Run()
			return nil
		},
	}
)
