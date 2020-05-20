package wechat

import (
	"strconv"
)

type WechatError struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func (a WechatError) Error() string {
	return strconv.Itoa(a.ErrCode) + ":" + a.ErrMsg
}
func IsInvalidCredential(err error) bool {
	return parseErrorCode(40001, err)
}
func IsAccessTokenExpired(err error) bool {
	return parseErrorCode(42001, err)
}
func IsRefreshTokenExpired(err error) bool {
	return parseErrorCode(42002, err)
}
func IsInvalidOpenId(err error) bool {
	return parseErrorCode(40003, err)
}
func IsInvalidCode(err error) bool {
	return parseErrorCode(40029, err)
}
func IsInvalidRefreshToken(err error) bool {
	return parseErrorCode(40030, err)
}

func parseErrorCode(code int, err error) bool {
	if e, ok := err.(WechatError); ok {
		if e.ErrCode == code {
			return true
		}
	}
	return false
}

//返回码 	错误码描述 	说明
//40001 	invalid credential 	不合法的调用凭证
//40002 	invalid grant_type 	不合法的 grant_type
//40003 	invalid openid 	不合法的 OpenID
//40004 	invalid media type 	不合法的媒体文件类型
//40007 	invalid media_id 	不合法的 media_id
//40008 	invalid message type 	不合法的 message_type
//40009 	invalid image size 	不合法的图片大小
//40010 	invalid voice size 	不合法的语音大小
//40011 	invalid video size 	不合法的视频大小
//40012 	invalid thumb size 	不合法的缩略图大小
//40013 	invalid appid 	不合法的 AppID
//40014 	invalid access_token 	不合法的 access_token
//40015 	invalid menu type 	不合法的菜单类型
//40016 	invalid button size 	不合法的菜单按钮个数
//40017 	invalid button type 	不合法的按钮类型
//40018 	invalid button name size 	不合法的按钮名称长度
//40019 	invalid button key size 	不合法的按钮 KEY 长度
//40020 	invalid button url size 	不合法的 url 长度
//40023 	invalid sub button size 	不合法的子菜单按钮个数
//40024 	invalid sub button type 	不合法的子菜单类型
//40025 	invalid sub button name size 	不合法的子菜单按钮名称长度
//40026 	invalid sub button key size 	不合法的子菜单按钮 KEY 长度
//40027 	invalid sub button url size 	不合法的子菜单按钮 url 长度
//40029 	invalid code 	不合法或已过期的 code
//40030 	invalid refresh_token 	不合法的 refresh_token
//40036 	invalid template_id size 	不合法的 template_id 长度
//40037 	invalid template_id 	不合法的 template_id
//40039 	invalid url size 	不合法的 url 长度
//40048 	invalid url domain 	不合法的 url 域名
//40054 	invalid sub button url domain 	不合法的子菜单按钮 url 域名
//40055 	invalid button url domain 	不合法的菜单按钮 url 域名
//40066 	invalid url 	不合法的 url
//41001 	access_token missing 	缺失 access_token 参数
//41002 	appid missing 	缺失 appid 参数
//41003 	refresh_token missing 	缺失 refresh_token 参数
//41004 	appsecret missing 	缺失 secret 参数
//41005 	media data missing 	缺失二进制媒体文件
//41006 	media_id missing 	缺失 media_id 参数
//41007 	sub_menu data missing 	缺失子菜单数据
//41008 	missing code 	缺失 code 参数
//41009 	missing openid 	缺失 openid 参数
//41010 	missing url 	缺失 url 参数
//42001 	access_token expired 	access_token 超时
//42002 	refresh_token expired 	refresh_token 超时
//42003 	code expired 	code 超时
//43001 	require GET method 	需要使用 GET 方法请求
//43002 	require POST method 	需要使用 POST 方法请求
//43003 	require https 	需要使用 HTTPS
//43004 	require subscribe 	需要订阅关系
//44001 	empty media data 	空白的二进制数据
//44002 	empty post data 	空白的 POST 数据
//44003 	empty news data 	空白的 news 数据
//44004 	empty content 	空白的内容
//44005 	empty list size 	空白的列表
//45001 	media size out of limit 	二进制文件超过限制
//45002 	content size out of limit 	content 参数超过限制
//45003 	title size out of limit 	title 参数超过限制
//45004 	description size out of limit 	description 参数超过限制
//45005 	url size out of limit 	url 参数长度超过限制
//45006 	picurl size out of limit 	picurl 参数超过限制
//45007 	playtime out of limit 	播放时间超过限制（语音为 60s 最大）
//45008 	article size out of limit 	article 参数超过限制
//45009 	api freq out of limit 	接口调动频率超过限制
//45010 	create menu limit 	建立菜单被限制
//45011 	api limit 	频率限制
//45012 	template size out of limit 	模板大小超过限制
//45016 	can't modify sys group 	不能修改默认组
//45017 	can't set group name too long sys group 	修改组名过长
//45018 	too many group now, no need to add new 	组数量过多
//50001 	api unauthorized 	接口未授权
