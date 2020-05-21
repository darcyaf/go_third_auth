package wechat

import (
	"context"
	"encoding/json"
)

var (
	Sex_Male   = 1
	Sex_Female = 2
)

var userInfoPath = "/sns/userinfo"

type UserInfoPayload struct {
	AccessToken string `json:"access_token"`
	OpenId      string `json:"openid"`
}
type UserInfoResp struct {
	OpenId       string   `json:"openid"`
	Nickname     string   `json:"nickname"`
	Sex          int      `json:"sex"` // 1为男 2为女
	Province     string   `json:"province"`
	City         string   `json:"city"`
	Country      string   `json:"country"`
	HeadImageUrl string   `json:"headimgurl"`
	Privilege    []string `json:"privilege"`
	UnionId      string   `json:"unionid"`
}

func (w Wechat) UserInfo(ctx context.Context, payload UserInfoPayload) (UserInfoResp, error) {
	var resp UserInfoResp
	respRaw, err := w.Do(ctx, userInfoPath, payload)
	if err != nil {
		return resp, err
	}
	err = json.Unmarshal(respRaw, &resp)
	return resp, err
}
func NewUserInfoPayload(accessToken string, openid string) UserInfoPayload {
	return UserInfoPayload{
		AccessToken: accessToken,
		OpenId:      openid,
	}
}