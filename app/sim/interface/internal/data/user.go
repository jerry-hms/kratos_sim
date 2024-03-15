package data

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwt2 "github.com/golang-jwt/jwt/v4"
	userV1 "kratos_sim/api/user/service/v1"
	wsV1 "kratos_sim/api/ws/service/v1"
	"kratos_sim/app/sim/interface/internal/biz"
)

type UserRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &UserRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "repo/server")),
	}
}

func (u *UserRepo) FindUserByUsername(ctx context.Context, username string) (*userV1.GetUserByUsernameReply, error) {
	return u.data.uc.GetUserByUsername(ctx, &userV1.GetUserByUsernameReq{Username: username})
}

func (u *UserRepo) CreateUser(ctx context.Context, user *biz.User) (*userV1.CreateUserReply, error) {
	return u.data.uc.CreateUser(ctx, &userV1.CreateUserReq{
		Username: user.Username,
		Password: user.Password,
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
		Mobile:   user.Mobile,
	})
}

func (u *UserRepo) GetUser(ctx context.Context, b *biz.User) (*userV1.GetUserReply, error) {
	return u.data.uc.GetUser(ctx, &userV1.GetUserReq{
		Id: b.Id,
	})
}

func (u *UserRepo) BindWs(ctx context.Context, client_id string) (*wsV1.BindReply, error) {
	claims, ok := jwt.FromContext(ctx)
	if !ok {
		return nil, errors.New("获取用户信息失败")
	}
	user, ok := claims.(*jwt2.MapClaims)
	if !ok {
		return nil, errors.New("获取用户信息失败")
	}
	user_id := (*user)["user_id"].(float64)
	_, err := u.data.wc.Bind(ctx, &wsV1.BindReq{
		UserId:   int64(user_id),
		ClientId: client_id,
	})
	if err != nil {
		return nil, err
	}
	return &wsV1.BindReply{}, nil
}
