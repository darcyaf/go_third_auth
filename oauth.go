package go_third_auth

import "github.com/darcyaf/go_third_auth/wechat"

func NewWechat(opts ...wechat.Option) *wechat.Wechat {
	return wechat.NewWechat(opts...)
}
