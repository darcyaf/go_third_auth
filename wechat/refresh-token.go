package wechat

import (
	"context"
	"encoding/json"
)

var refreshTokenPath = "/sns/oauth2/refresh_token"

type RefreshTokenPayload struct {
	RefreshToken string `json:"refresh_token"`
	GrantType    string `json:"grant_type"`
}
type RefreshTokenResp struct {
	AccessTokenInfo
	Scope string `json:"scope"`
}

func (w Wechat) RefreshToken(ctx context.Context, payload RefreshTokenPayload) (RefreshTokenResp, error) {
	var resp RefreshTokenResp
	respRaw, err := w.Do(ctx, refreshTokenPath, payload)
	if err != nil {
		return resp, err
	}
	err = json.Unmarshal(respRaw, &resp)
	return resp, err
}
func NewRefreshTokenPayload(refreshToken string) RefreshTokenPayload {
	return RefreshTokenPayload{
		RefreshToken: refreshToken,
		GrantType:    "refresh_token",
	}
}
