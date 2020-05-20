# go-third-auth
第三方登录授权,微信,QQ,微博等
## 微信授权
- 根据code获取accessToken
```
    var w = NewWechat(SetAppKey("your-app-key"), SetAppSecret("your-app-secret"))
	code := ""
	resp, err := w.AccessToken(context.Background(), NewAccessTokenPayload(code))
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Printf("%+v\n", resp)
```
- 根据accessToken获取userInfo
```
    var openId =""
    var accessToken = ""
    var w = NewWechat(SetAppKey("your-app-key"), SetAppSecret("your-app-secret"))
	resp, err := w.UserInfo(context.Background(), NewUserInfoPayload(accessToken, openId))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", resp)
```
- 根据refreshToken刷新accessToken
- 提供了UserInfoByCode,根据code获取userinfo
- 提供了UserInfoByAccessTokenInfo根据accessInfo获取userInfo,并自动刷新token