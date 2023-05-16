package admin

import (
	"context"
	"fmt"
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	frontend "shop/api/frontend"
	"shop/internal/consts"
	"shop/internal/dao"
	"shop/internal/model"
	"shop/internal/model/entity"
	"shop/internal/service"
	"shop/utility"
	"shop/utility/response"
)

type sUser struct{}

func init() {
	// 注册路由
	service.RegisterUser(New())
}
func New() *sUser {
	return &sUser{}
}

// IsExist 查询用户是否存在
func (s *sUser) IsExist(ctx context.Context, username string) bool {
	var user *entity.UserInfo
	err := dao.UserInfo.Ctx(ctx).Where(dao.UserInfo.Columns().Name, username).Scan(&user)
	if err != nil {
		return false
	}
	if user != nil {
		return true
	}
	return false
}

// Register 用户注册
func (s *sUser) Register(ctx context.Context, in *model.UserRegisterInput) (out model.UserRegisterOutput, err error) {
	// 判断用户是否存在
	if s.IsExist(ctx, in.Name) {
		return model.UserRegisterOutput{}, fmt.Errorf(consts.ErrUserIsExistMsg)
	}
	// 加密密码
	salt := grand.S(10)
	in.Password = utility.EncryptPassword(in.Password, salt)
	in.UserSalt = salt
	id, err := dao.UserInfo.Ctx(ctx).InsertAndGetId(in)
	if err != nil {
		return model.UserRegisterOutput{}, err
	}
	return model.UserRegisterOutput{Id: id}, nil
}

// LoginBeforeFunc 前台用户登录
func (s *sUser) LoginBeforeFunc(r *ghttp.Request) (string, interface{}) {
	username := r.Get("name")
	password := r.Get("password")
	if username.String() == "" || password.String() == "" {
		response.JsonExit(r, -1, consts.ErrLoginFiledMsg)
	}
	var user = new(entity.UserInfo)
	ctx := r.Context()
	err := dao.UserInfo.Ctx(ctx).Where(dao.UserInfo.Columns().Name, username).Scan(&user)
	if err != nil {
		response.JsonExit(r, -1, consts.ErrLoginFiledMsg)
		r.ExitAll()
	}

	if !utility.CheckPassword(password.String(), user.Password, user.UserSalt) {
		response.JsonExit(r, -1, consts.ErrLoginFiledMsg)
		r.ExitAll()
	}

	return gconv.String(user.Id), user
}

// LoginAfterFunc GToken登录后操作
func (s *sUser) LoginAfterFunc(r *ghttp.Request, resp gtoken.Resp) {
	userId := resp.Get("userKey")
	var user = new(entity.UserInfo)
	err := dao.UserInfo.Ctx(r.Context()).Where(dao.User.Columns().Id, userId).Scan(&user)
	if err != nil {
		response.JsonExit(r, -1, consts.ErrLoginFiledMsg)
		return
	}
	response.JsonExit(r, consts.GTokenCode, consts.GTokenSuccessMsg, frontend.UserLoginRes{
		Type:   consts.GTokenType,
		Token:  resp.GetString(consts.GTokenTokenKey),
		Expire: consts.GTokenExpire,
		UserInfoBase: frontend.UserInfoBase{
			Name:   user.Name,
			Avatar: user.Avatar,
			Sex:    user.Sex,
			Sign:   user.Sign,
			Status: user.Status,
		},
	})
}

// AuthAfterFunc 授权后操作
func (s *sUser) AuthAfterFunc(r *ghttp.Request, resp gtoken.Resp) {
	var user entity.UserInfo
	err := gconv.Struct(resp.GetString("data"), &user)
	if err != nil {
		response.JsonExit(r, -1, consts.ErrAuthFiledMsg)
	}
	r.SetCtxVar(consts.ContextUserId, user.Id)
	r.SetCtxVar(consts.ContextUserName, user.Name)
	r.SetCtxVar(consts.ContextUserAvatar, user.Avatar)
	r.SetCtxVar(consts.ContextUserSex, user.Sex)
	r.SetCtxVar(consts.ContextUserStatus, user.Status)
	r.SetCtxVar(consts.ContextUserSign, user.Sign)
	r.Middleware.Next()
}
