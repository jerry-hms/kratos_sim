package biz

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	v1 "kratos_sim/api/sim/interface/v1"
	"kratos_sim/app/sim/interface/internal/conf"
	"kratos_sim/app/sim/interface/internal/pkg/encrypt"
	"kratos_sim/app/sim/interface/internal/pkg/jwt"
)

type AuthUseCase struct {
	userRepo UserRepo
	conf     *conf.Auth
	log      *log.Helper
}

func NewAuthUseCase(repo UserRepo, conf *conf.Auth, logger log.Logger) *AuthUseCase {
	return &AuthUseCase{
		userRepo: repo,
		conf:     conf,
		log:      log.NewHelper(log.With(logger, "module", "usecase/auth")),
	}
}

func (a *AuthUseCase) Register(ctx context.Context, req *v1.RegisterReq) (*v1.RegisterReply, error) {
	username, _ := a.userRepo.FindUserByUsername(ctx, req.Username)
	if username != nil {
		return nil, errors.New("该用户已存在！")
	}
	pass, _ := encrypt.HashString(req.Password)
	newUser, err := a.userRepo.CreateUser(ctx, &User{
		Username: req.Username,
		Password: pass,
		Nickname: req.Nickname,
		Avatar:   req.Avatar,
		Mobile:   req.Mobile,
	})
	if err != nil {
		return nil, err
	}
	return &v1.RegisterReply{
		UserInfo: &v1.User{
			Id:        newUser.Id,
			Nickname:  newUser.Nickname,
			Avatar:    newUser.Avatar,
			Mobile:    newUser.Mobile,
			CreatedAt: newUser.CreatedAt,
		},
	}, nil
}

func (a *AuthUseCase) Login(ctx context.Context, req *v1.LoginReq) (*v1.LoginReply, error) {
	user, _ := a.userRepo.FindUserByUsername(ctx, req.Username)
	if user == nil {
		return nil, errors.New("用户名或密码不正确")
	}

	err := encrypt.CheckPassword(req.Password, user.Password)
	if err != nil {
		return nil, errors.New("用户名或密码不正确")
	}

	token, err := jwt.CreateUserTokenFactory(a.conf.ApiKey).GenerateToken(jwt.Claims{
		UserId:    user.Id,
		CreatedAt: user.CreatedAt,
	}, a.conf.Ttl)
	if err != nil {
		return nil, err
	}
	return &v1.LoginReply{
		UserInfo: &v1.User{
			Id:        user.Id,
			Nickname:  user.Nickname,
			Avatar:    user.Avatar,
			Mobile:    user.Mobile,
			CreatedAt: user.CreatedAt,
		},
		Token: token,
	}, nil
}
