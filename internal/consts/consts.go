package consts

const (
	Version    = "1.0.0"      // 版本
	ContextKey = "ContextKey" // 上下文键

	ContextUserId      = "id"
	ContextUsername    = "username"
	ContextUserIsAdmin = "is_admin"
	ContextUserRoleIds = "role_ids"

	GTokenExpire     = 10 * 24 * 60 * 60
	GTokenTokenKey   = "token"
	GTokenSuccessMsg = "登录成功"
	GTokenType       = "Bearer"
	GTokenCode       = 0

	ErrLoginFiledMsg = "登录失败: 账号或秘密错误"
	ErrAuthFiledMsg  = "授权失败: Token校验失败"
)
