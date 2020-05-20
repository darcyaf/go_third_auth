package wechat

import (
	"context"
	"fmt"
	"testing"
	"time"
)

var (
	w            = NewWechat(SetAppKey(""), SetAppSecret(""))
	openId       = ""
	accessToken  = ""
	refreshToken = ""
)

func TestWechat_AccessToken(t *testing.T) {
	code := "071vCzw90WHWnA1fdKz90klww90vCzwW"
	resp, err := w.AccessToken(context.Background(), NewAccessTokenPayload(code))
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Printf("%+v\n", resp)
}
func TestWechat_RefreshToken(t *testing.T) {
	resp, err := w.RefreshToken(context.Background(), NewRefreshTokenPayload(refreshToken))
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Printf("%+v\n", resp)
}
func TestWechat_UserInfo(t *testing.T) {
	resp, err := w.UserInfo(context.Background(), NewUserInfoPayload(accessToken, openId))
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Printf("%+v", resp)
}
func TestWechat_UserInfoByAccessTokenInfo(t *testing.T) {

	var accessTokenInfo = AccessTokenInfo{
		AccessToken:  accessToken,
		ExpiresIn:    7200,
		RefreshToken: refreshToken,
		OpenId:       openId,
	}
	resp, newAccessInfo, err := w.UserInfoByAccessTokenInfo(context.Background(), accessTokenInfo, time.Now())
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Printf("%+v\n", resp)
	fmt.Printf("%+v\n", newAccessInfo)
}
