package login

import (
	"context"
	"fmt"
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"shop/api/backend"
	"shop/internal/consts"
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
	var user = new(entity.AdminInfo)
	err = dao.AdminInfo.Ctx(ctx).Where(dao.User.Columns().Username, in.Username).Scan(&user)
	if err != nil {
		return err
	}
	encryptPassword := utility.EncryptPassword(in.Password, user.UserSalt)
	if user.Password != encryptPassword {
		return fmt.Errorf("用户名或密码错误")
	}
	err = service.Session().SetUser(ctx, user)
	if err != nil {
		return err
	}
	service.BizCtx().SetUser(ctx, &model.ContextUser{
		Id:       user.Id,
		Username: user.Name,
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
	var user = new(entity.AdminInfo)
	ctx := r.Context()
	err := dao.AdminInfo.Ctx(ctx).Where(dao.AdminInfo.Columns().Name, username).Scan(&user)
	if err != nil {
		response.JsonExit(r, -1, "账号或密码错误.")
		r.ExitAll()
	}
	encryptPassword := utility.EncryptPassword(password.String(), user.UserSalt)
	if user.Password != encryptPassword {
		response.JsonExit(r, -1, "账号或密码错误.")
		r.ExitAll()
	}

	return gconv.String(user.Id), model.AdminItemOutput{
		Id:      user.Id,
		Name:    user.Name,
		IsAdmin: user.IsAdmin,
		RoleIds: gconv.Ints(user.RoleIds),
	}
}

// LoginAfterFunc GToken登录后操作
func (s *sLogin) LoginAfterFunc(r *ghttp.Request, resp gtoken.Resp) {
	userId := resp.Get("userKey")
	var user = new(entity.AdminInfo)
	err := dao.AdminInfo.Ctx(r.Context()).Where(dao.User.Columns().Id, userId).Scan(&user)
	fmt.Println(user)
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
	var user model.AdminItemOutput
	err := gconv.Struct(resp.GetString("data"), &user)
	if err != nil {
		response.JsonExit(r, -1, "请求未授权.")
	}
	r.SetCtxVar(consts.ContextUserId, user.Id)
	r.SetCtxVar(consts.ContextUsername, user.Name)
	r.SetCtxVar(consts.ContextUserIsAdmin, user.IsAdmin)
	r.SetCtxVar(consts.ContextUserRoleIds, user.RoleIds)
	r.Middleware.Next()
}
