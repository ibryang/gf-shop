package consts

const (
	Version    = "1.0.0"      // 版本
	ContextKey = "ContextKey" // 上下文键

	ContextAdminId      = "id"
	ContextAdminName    = "username"
	ContextAdminIsAdmin = "is_admin"
	ContextAdminRoleIds = "role_ids"

	ContextUserId     = "id"
	ContextUserName   = "name"
	ContextUserSex    = "sex"
	ContextUserAvatar = "avatar"
	ContextUserStatus = "status"
	ContextUserSign   = "sign"
	
	GTokenExpire     = 10 * 24 * 60 * 60
	GTokenTokenKey   = "token"
	GTokenSuccessMsg = "登录成功"
	GTokenType       = "Bearer"
	GTokenCode       = 0

	ErrLoginFiledMsg  = "登录失败: 账号或秘密错误"
	ErrAuthFiledMsg   = "授权失败: Token校验失败"
	ErrUserIsExistMsg = "用户已存在"
)
