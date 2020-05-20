package wechat

import (
	"context"
	"time"
)

var (
	baseUri = "https://api.weixin.qq.com"
)

//调用频率限制
//接口名 	频率限制
//通过 code 换取 access_token 	5 万/分钟
//获取用户基本信息 	5 万/分钟
//刷新 access_token 	10 万/分钟
type Option func(o *options)

type Wechat struct {
	opts *options
}
type options struct {
	AppId     string `json:"appid"`
	AppSecret string `json:"secret"`
}

func NewWechat(opts ...Option) *Wechat {
	o := options{}
	for _, opt := range opts {
		opt(&o)
	}
	w := &Wechat{
		opts: &o,
	}
	return w
}

func SetAppKey(appKey string) Option {
	return func(o *options) {
		o.AppId = appKey
	}
}
func SetAppSecret(appSecret string) Option {
	return func(o *options) {
		o.AppSecret = appSecret
	}
}
func (w Wechat) UserInfoByCode(ctx context.Context, code string) (UserInfoResp, AccessTokenResp, error) {
	var accessTokenResp AccessTokenResp
	var userInfoResp UserInfoResp
	var err error

	accessTokenResp, err = w.AccessToken(ctx, NewAccessTokenPayload(code))
	if err != nil {
		return userInfoResp, accessTokenResp, err
	}
	userInfoResp, err = w.UserInfo(ctx, NewUserInfoPayload(accessTokenResp.AccessToken, accessTokenResp.OpenId))
	return userInfoResp, accessTokenResp, err
}
func (w Wechat) UserInfoByAccessTokenInfo(ctx context.Context, accessInfo AccessTokenInfo, lastRequestTime time.Time) (UserInfoResp, AccessTokenInfo, error) {
	var userInfoResp UserInfoResp
	var refreshTokenResp RefreshTokenResp
	var err error
	var refreshed bool
	if lastRequestTime.Add(time.Duration(accessInfo.ExpiresIn) * time.Second).Before(time.Now()) {
		refreshed = true
		refreshTokenResp, err = w.RefreshToken(ctx, NewRefreshTokenPayload(accessInfo.RefreshToken))
		if err != nil {
			return userInfoResp, accessInfo, err
		}
		accessInfo = refreshTokenResp.AccessTokenInfo
	}
	userInfoResp, err = w.UserInfo(ctx, NewUserInfoPayload(accessInfo.AccessToken, accessInfo.OpenId))

	// retry if accessToken is Expired
	if !refreshed && (IsAccessTokenExpired(err) || IsInvalidCredential(err)) {
		refreshTokenResp, err = w.RefreshToken(ctx, NewRefreshTokenPayload(accessInfo.RefreshToken))
		if err != nil {
			return userInfoResp, accessInfo, err
		}
		accessInfo = refreshTokenResp.AccessTokenInfo
		userInfoResp, err = w.UserInfo(ctx, NewUserInfoPayload(accessInfo.AccessToken, accessInfo.OpenId))
	}
	return userInfoResp, accessInfo, err
}
