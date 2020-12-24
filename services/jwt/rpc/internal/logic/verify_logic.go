package logic

import (
	"context"
	"fmt"
	jwtGo "github.com/dgrijalva/jwt-go"
	"github/yguilai/timetable-micro/services/jwt/rpc/internal/svc"
	"github/yguilai/timetable-micro/services/jwt/rpc/jwt"

	"github.com/tal-tech/go-zero/core/logx"
)

type VerifyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVerifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyLogic {
	return &VerifyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *VerifyLogic) Verify(in *jwt.JwtVerifyReq) (*jwt.JwtVerifyResp, error) {
	token, err := jwtGo.Parse(in.Token, func(token *jwtGo.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwtGo.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(l.svcCtx.Config.Auth.AccessSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if _, ok := token.Claims.(jwtGo.MapClaims); ok && token.Valid {
		return &jwt.JwtVerifyResp{Valid: true}, nil
	}
	return &jwt.JwtVerifyResp{Valid: false}, nil
}
