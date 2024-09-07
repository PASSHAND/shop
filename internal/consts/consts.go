package consts

const (
	ProjectName              = "GoFrame Shop"
	ProjectUsage             = "study author:zgm"
	ProjectBrief             = "start http server"
	Version                  = "v0.2.0"             // 当前服务版本(用于模板展示)
	CaptchaDefaultName       = "CaptchaDefaultName" // 验证码默认存储空间名称
	ContextKey               = "ContextKey"         // 上下文变量存储键名，前后端系统共享
	FileMaxUploadCountMinute = 10                   // 同一用户1分钟之内最大上传数量
	GTokenAdminPrefix        = "Admin:"             //gtoken管理后台前缀
	GTokenFrontendPrefix     = "User:"              //gtoken管理前台前缀
	//for admin
	CtxAdminId      = "CtxAdminId"
	CtxAdminName    = "CtxAdminName"
	CtxAdminIsAdmin = "CtxAdminIsAdmin"
	CtxAdminRoleIds = "CtxAdminRoleIds"
	//for user
	CtxUserId     = "CtxUserId"
	CtxUserName   = "CtxUserName"
	CtxUserAvatar = "CtxUserAvatar"
	CtxUserSex    = "CtxUserSex"
	CtxUserSign   = "CtxUserSign"
	CtxUserStatus = "CtxUserStatus"
	//for login
	CacheModelRedis    = 2
	TokenType          = "Bearer"
	BackendServerName  = "shop"
	MultiLogin         = true
	FrontendMultiLogin = false
	GTokenExpireIn     = 10 * 24 * 60 * 60
	//同一管理错误信息
	CodeMissingParameterMsg = "请检查是否缺少参数"
	ErrLoginFaulMsg         = "登录失败，账号或密码错误"
)
