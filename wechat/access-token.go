package wechat

import (
	"context"
	"encoding/json"
)

var accessTokenPath = "/sns/oauth2/access_token"

type AccessTokenResp struct {
	AccessTokenInfo
	Scope   string `json:"scope"`
	UnionId string `json:"unionid"`
}
type AccessTokenInfo struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int64  `json:"expires_in"` //second unit
	RefreshToken string `json:"refresh_token"`
	OpenId       string `json:"openid"`
}
type AccessTokenPayload struct {
	Code      string `json:"code"`
	GrantType string `json:"grant_type"`
}

//AccessToken get accessToken by authorization code
func (w Wechat) AccessToken(ctx context.Context, payload AccessTokenPayload) (AccessTokenResp, error) {
	var resp AccessTokenResp
	respRaw, err := w.Do(ctx, accessTokenPath, payload)
	if err != nil {
		return resp, err
	}
	err = json.Unmarshal(respRaw, &resp)
	return resp, err
}

func NewAccessTokenPayload(code string) AccessTokenPayload {
	return AccessTokenPayload{
		Code:      code,
		GrantType: "authorization_code",
	}
}
