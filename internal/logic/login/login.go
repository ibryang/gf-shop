package login

import (
	"context"
	"fmt"
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"shop/api/backend"
	"shop/internal/dao"
	"shop/internal/model"
	"shop/internal/model/entity"
	"shop/internal/service"
	"shop/utility"
	"shop/utility/response"
)

type sLogin struct{}

func init() {
	// 注册路由
	service.RegisterLogin(New())
}
func New() *sLogin {
	return &sLogin{}
}

func (s *sLogin) Login(ctx context.Context, in *model.LoginInput) (err error) {
	var user = new(entity.User)
	err = dao.User.Ctx(ctx).Where(dao.User.Columns().Username, in.Username).Scan(&user)
	if err != nil {
		return err
	}
	encryptPassword := utility.EncryptPassword(in.Password, user.Salt)
	if user.Password != encryptPassword {
		return fmt.Errorf("用户名或密码错误")
	}
	err = service.Session().SetUser(ctx, user)
	if err != nil {
		return err
	}
	service.BizCtx().SetUser(ctx, &model.ContextUser{
		Id:       user.Id,
		Username: user.Username,
		IsAdmin:  user.IsAdmin,
	})

	return
}

// LoginBeforeFunc GToken登录
func (s *sLogin) LoginBeforeFunc(r *ghttp.Request) (string, interface{}) {
	username := r.Get("username")
	password := r.Get("password")
	if username.String() == "" || password.String() == "" {
		response.JsonExit(r, -1, "账号或密码错误.")
	}
	var user = new(entity.User)
	ctx := r.Context()
	err := dao.User.Ctx(ctx).Where(dao.User.Columns().Username, username).Scan(&user)
	if err != nil {
		response.JsonExit(r, -1, "账号或密码错误.")
		r.ExitAll()
	}
	encryptPassword := utility.EncryptPassword(password.String(), user.Salt)
	if user.Password != encryptPassword {
		response.JsonExit(r, -1, "账号或密码错误.")
		r.ExitAll()
	}

	return gconv.String(user.Id), model.UserItemOutput{
		Id:       user.Id,
		Username: user.Username,
		IsAdmin:  user.IsAdmin,
		RoleIds:  gconv.Ints(user.RoleIds),
	}
}

// LoginAfterFunc GToken登录后操作
func (s *sLogin) LoginAfterFunc(r *ghttp.Request, resp gtoken.Resp) {
	userId := resp.Get("userKey")
	var user = new(entity.User)
	err := dao.User.Ctx(r.Context()).Where(dao.User.Columns().Id, userId).Scan(&user)
	if err != nil {
		response.JsonExit(r, -1, "账号或密码错误.")
		return
	}
	response.JsonExit(r, 0, "登录成功.", backend.LoginRes{
		Type:    "Bearer",
		Token:   resp.GetString("token"),
		Expire:  10 * 24 * 60 * 60,
		IsAdmin: user.IsAdmin,
		RoleIds: gconv.Ints(user.RoleIds),
	})
}

// AuthAfterFunc 授权后操作
func (s *sLogin) AuthAfterFunc(r *ghttp.Request, resp gtoken.Resp) {
	var user model.UserItemOutput
	err := gconv.Struct(resp.GetString("data"), &user)
	if err != nil {
		response.JsonExit(r, -1, "账号或密码错误.")
	}
	r.SetCtxVar("id", user.Id)
	r.SetCtxVar("username", user.Username)
	r.SetCtxVar("isAdmin", user.IsAdmin)
	r.SetCtxVar("roleIds", user.RoleIds)
	r.Middleware.Next()
}
