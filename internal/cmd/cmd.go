package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"goframe-shop/internal/consts"
	"goframe-shop/internal/controller"
	"goframe-shop/internal/service"

	"goframe-shop/internal/controller/hello"
)

// var gfToken *gtoken.GfToken
var (
	Main = gcmd.Command{
		Name:  consts.ProjectName,
		Usage: consts.ProjectUsage,
		Brief: consts.ProjectBrief,
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			// 启动管理后台gtoken
			gfAdminToken, err := StartBackendGToken()
			s.Group("/backend", func(group *ghttp.RouterGroup) {
				//group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Middleware( //自定义中间件
					service.Middleware().CORS,
					service.Middleware().Ctx,
					service.Middleware().ResponseHandler,
				)
				//gtoken中间件
				//err := gfToken.Middleware(ctx, group)
				//if err != nil {
				//	panic(err)
				//}
				//不需要登陆的路由组绑定
				group.Bind(
					hello.NewV1(),
					controller.Rotation, //轮播图
					controller.Position, //手工位
					controller.Admin.Create,
					controller.Admin.List,
					controller.Admin.Delete,
					controller.Admin.Update, //管理员
					controller.Login,        //登陆
					controller.Data,         //数据大屏
					controller.Role,         //角色
					controller.Permission,   //权限
				)
				//需要登录的路由组绑定
				group.Group("/", func(group *ghttp.RouterGroup) {
					//for jwt
					//group.Middleware(service.Middleware().Auth)
					err := gfAdminToken.Middleware(ctx, group)
					if err != nil {
						panic(err)
					}
					group.ALLMap(g.Map{
						"/admin/info": controller.Admin.Info,
					})
					group.Bind(
						controller.File,   //文件入库
						controller.Upload, //可跨项目使用的文件上云工具类
					)
				})
			})
			s.Run()
			return nil
		},
	}
)
