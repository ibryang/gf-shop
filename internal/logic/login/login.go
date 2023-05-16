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
	if err != nil || user == nil {
		err = fmt.Errorf(consts.ErrLoginFiledMsg)
		return err
	}

	// Verify the user's password.
	if !utility.CheckPassword(in.Password, user.Password, user.UserSalt) {
		return fmt.Errorf(consts.ErrLoginFiledMsg)
	}
	// Set the user in the session and context.
	if err := service.Session().SetUser(ctx, user); err != nil {
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
		response.JsonExit(r, -1, consts.ErrLoginFiledMsg)
	}
	var user = new(entity.AdminInfo)
	ctx := r.Context()
	err := dao.AdminInfo.Ctx(ctx).Where(dao.AdminInfo.Columns().Name, username).Scan(&user)
	if err != nil {
		response.JsonExit(r, -1, consts.ErrLoginFiledMsg)
		r.ExitAll()
	}
	encryptPassword := utility.EncryptPassword(password.String(), user.UserSalt)
	if user.Password != encryptPassword {
		response.JsonExit(r, -1, consts.ErrLoginFiledMsg)
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
	if err != nil {
		response.JsonExit(r, -1, consts.ErrLoginFiledMsg)
		return
	}
	response.JsonExit(r, consts.GTokenCode, consts.GTokenSuccessMsg, backend.LoginRes{
		Type:    consts.GTokenType,
		Token:   resp.GetString(consts.GTokenTokenKey),
		Expire:  consts.GTokenExpire,
		IsAdmin: user.IsAdmin,
		RoleIds: gconv.Ints(user.RoleIds),
	})
}

// AuthAfterFunc 授权后操作
func (s *sLogin) AuthAfterFunc(r *ghttp.Request, resp gtoken.Resp) {
	var user model.AdminItemOutput
	err := gconv.Struct(resp.GetString("data"), &user)
	if err != nil {
		response.JsonExit(r, -1, consts.ErrAuthFiledMsg)
	}
	r.SetCtxVar(consts.ContextAdminId, user.Id)
	r.SetCtxVar(consts.ContextAdminName, user.Name)
	r.SetCtxVar(consts.ContextAdminIsAdmin, user.IsAdmin)
	r.SetCtxVar(consts.ContextAdminRoleIds, user.RoleIds)
	r.Middleware.Next()
}
