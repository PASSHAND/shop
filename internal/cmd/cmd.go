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
			//管理后台路由组
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
					controller.Admin.Create,
					controller.Login, //jwt登陆,已弃用
				)
				//需要登录的路由组绑定
				group.Group("/", func(group *ghttp.RouterGroup) {
					//for jwt
					//group.Middleware(service.Middleware().Auth)
					err := gfAdminToken.Middleware(ctx, group)
					if err != nil {
						panic(err)
					}
					//group.ALLMap(g.Map{
					//	"/admin/info": controller.Admin.Info,
					//})
					group.Bind(
						controller.Rotation,   //轮播图
						controller.Position,   //手工位
						controller.Data,       //数据大屏
						controller.Role,       //角色
						controller.Permission, //权限
						controller.Admin.List,
						controller.Admin.Delete,
						controller.Admin.Update, //管理员
						controller.Admin.Info,   //查询当前管理员信息
						controller.File,         //文件入库
						controller.Upload,       //可跨项目使用的文件上云工具类
						controller.Category,     //商品分类管理
						controller.Coupon,       //商品优惠券
						controller.UserCoupon,   //用户优惠券
						controller.Goods,        //商品管理
						controller.GoodsOptions, //商品规格
						controller.Article,      //文章管理
					)
				})
			})
			//---------------------------------------------------------------------
			//启动前台路由
			frontendToken, err := StartFrontendGToken()
			if err != nil {
				return err
			}
			//管理前台路由组
			s.Group("/frontend", func(group *ghttp.RouterGroup) {
				group.Middleware( //自定义中间件
					service.Middleware().CORS,
					service.Middleware().Ctx,
					service.Middleware().ResponseHandler,
				)
				group.Bind(
					controller.User.Register,
					controller.Goods,
					//controller.Login, //登陆
				)
				//需要登录鉴权路由组
				group.Group("/", func(group *ghttp.RouterGroup) {
					err := frontendToken.Middleware(ctx, group)
					if err != nil {
						return
					}
					group.Bind(
						controller.User.Info,
						controller.User.UpdatePassword, //当前用户修改密码
						controller.Collection,          //收藏
						controller.Praise,              //点赞
						controller.Comment,             //评论
					)
				})
			})
			s.Run()
			return nil
		},
	}
)
