package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "kratos_sim/api/sim/interface/v1"
	userV1 "kratos_sim/api/user/service/v1"
	wsV1 "kratos_sim/api/ws/service/v1"
	"time"
)

type User struct {
	Id        int64
	Username  string
	Password  string
	Nickname  string
	Avatar    string
	Mobile    string
	CreatedAt time.Time
}

type UserRepo interface {
	CreateUser(context.Context, *User) (*userV1.CreateUserReply, error)
	FindUserByUsername(context.Context, string) (*userV1.GetUserByUsernameReply, error)
	BindWs(context.Context, string) (*wsV1.BindReply, error)
	GetUser(context.Context, *User) (*userV1.GetUserReply, error)
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "usecase/interface")),
	}
}

func (uc *UserUseCase) Bind(ctx context.Context, client_id string) (*v1.BindWsReply, error) {
	_, err := uc.repo.BindWs(ctx, client_id)
	if err != nil {
		return nil, err
	}
	return &v1.BindWsReply{}, nil
}
