package svc

import (
	"github.com/tal-tech/go-zero/zrpc"
	"github.com/yguilai/timetable-micro/gateway/internal/config"
	"github.com/yguilai/timetable-micro/services/captcha/rpc/captchaclient"
	"github.com/yguilai/timetable-micro/services/jwt/rpc/jwtclient"
	"github.com/yguilai/timetable-micro/services/user/rpc/userclient"
)

type ServiceContext struct {
	Config     config.Config
	CaptchaRpc captchaclient.Captcha
	UserRpc    userclient.User
	JwtRpc     jwtclient.Jwt
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		CaptchaRpc: captchaclient.NewCaptcha(zrpc.MustNewClient(c.CaptchaRpc)),
		JwtRpc:     jwtclient.NewJwt(zrpc.MustNewClient(c.JwtRpc)),
		UserRpc:    userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
