// Code generated by goctl. DO NOT EDIT!
// Source: jwt.proto

//go:generate mockgen -destination ./jwt_mock.go -package jwtclient -source $GOFILE

package jwtclient

import (
	"context"

	"github.com/yguilai/timetable-micro/services/jwt/rpc/jwt"

	"github.com/tal-tech/go-zero/zrpc"
)

type (
	Token          = jwt.Token
	JwtCreateReq   = jwt.JwtCreateReq
	JwtVerifyReq   = jwt.JwtVerifyReq
	JwtRefreshReq  = jwt.JwtRefreshReq
	Request        = jwt.Request
	Response       = jwt.Response
	Claim          = jwt.Claim
	JwtCreateResp  = jwt.JwtCreateResp
	JwtVerifyResp  = jwt.JwtVerifyResp
	JwtRefreshResp = jwt.JwtRefreshResp

	Jwt interface {
		Ping(ctx context.Context, in *Request) (*Response, error)
		Create(ctx context.Context, in *JwtCreateReq) (*JwtCreateResp, error)
		Verify(ctx context.Context, in *JwtVerifyReq) (*JwtVerifyResp, error)
		Refresh(ctx context.Context, in *JwtRefreshReq) (*JwtRefreshResp, error)
	}

	defaultJwt struct {
		cli zrpc.Client
	}
)

func NewJwt(cli zrpc.Client) Jwt {
	return &defaultJwt{
		cli: cli,
	}
}

func (m *defaultJwt) Ping(ctx context.Context, in *Request) (*Response, error) {
	client := jwt.NewJwtClient(m.cli.Conn())
	return client.Ping(ctx, in)
}

func (m *defaultJwt) Create(ctx context.Context, in *JwtCreateReq) (*JwtCreateResp, error) {
	client := jwt.NewJwtClient(m.cli.Conn())
	return client.Create(ctx, in)
}

func (m *defaultJwt) Verify(ctx context.Context, in *JwtVerifyReq) (*JwtVerifyResp, error) {
	client := jwt.NewJwtClient(m.cli.Conn())
	return client.Verify(ctx, in)
}

func (m *defaultJwt) Refresh(ctx context.Context, in *JwtRefreshReq) (*JwtRefreshResp, error) {
	client := jwt.NewJwtClient(m.cli.Conn())
	return client.Refresh(ctx, in)
}