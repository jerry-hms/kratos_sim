package server

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	jwt2 "github.com/golang-jwt/jwt/v4"
	"go.opentelemetry.io/otel/sdk/trace"
	"kratos_sim/api/sim/interface/v1"
	"kratos_sim/app/sim/interface/internal/conf"
	"kratos_sim/app/sim/interface/internal/pkg/response"
	"kratos_sim/app/sim/interface/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

func NewWhiteListMatcher() selector.MatchFunc {
	whiteList := make(map[string]struct{})
	whiteList[v1.Sim_Login_FullMethodName] = struct{}{}
	whiteList[v1.Sim_Register_FullMethodName] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}

func ConstMyJwtError() {
	reason := "UNAUTHORIZED"
	jwt.ErrMissingJwtToken = errors.Unauthorized(reason, "请登录后再试")
	jwt.ErrTokenExpired = errors.Unauthorized(reason, "token已过期")
	jwt.ErrTokenInvalid = errors.Unauthorized(reason, "无效的token")
	jwt.ErrTokenParseFail = errors.Unauthorized(reason, "token解析失败")
}

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, sim *service.SimService, ac *conf.Auth, tp *trace.TracerProvider, logger log.Logger) *http.Server {
	// 修改jwt默认错误信息
	ConstMyJwtError()
	var opts = []http.ServerOption{
		http.ErrorEncoder(response.EncoderError()),
		http.ResponseEncoder(response.EncoderResponse()),
		http.Middleware(
			recovery.Recovery(),
			tracing.Server(
				tracing.WithTracerProvider(tp)),
			selector.Server(
				jwt.Server(func(token *jwt2.Token) (interface{}, error) {
					return []byte(ac.ApiKey), nil
				}, jwt.WithSigningMethod(jwt2.SigningMethodHS256), jwt.WithClaims(func() jwt2.Claims {
					return &jwt2.MapClaims{}
				})),
			).
				Match(NewWhiteListMatcher()).
				Build(),
		),
	}

	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)

	v1.RegisterSimHTTPServer(srv, sim)
	return srv
}
