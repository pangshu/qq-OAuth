#感谢：phachon
#源项目地址：https://github.com/phachon/qq-OAuth

#修改项
1、修复demo错误
2、增加返回openid

# qq-OAuth
QQ 第三方登录

# 功能
实现 QQ OAuth2.0 Open api 封装

# 示例
```
// 随机字符串，验证 state 防止 CSRF 攻击
var state = "adadasdad"
var appId = ""
var	appSecret = ""
var	callback = ""
var	scope = "" //多个以 ，号隔开，默认为"all"

// 登录逻辑
func Login(res http.ResponseWriter, req *http.Request) {
    // 初始化
	oAuth := NewOAuth(appId, appSecret, callback, scope)
	// 获取授权 URL
	loginUrl := oAuth.GetAuthorURL(state)

    // 跳转至 URL
	http.Redirect(res, req, loginUrl, 302)
}

// 回调逻辑
func Callback(res http.ResponseWriter, req *http.Request)  {
    req.ParseForm()
	reqState := req.Form.Get("state")
	if reqState != state {
	    // 防止 CSRF 攻击
		res.Write([]byte("CSRF!!"))
	}

	authCode := req.Form.Get("code")
	oAuth := NewOAuth(appId, appSecret, callback, scope)
	oAuth.Access(authCode)

	userInfo, httpCode, err := oAuth.GetUserInfo()

	// TODO
}
```
示例：example/http_example.go

## License

MIT

Thanks
---------
Create By phachon@163.com
update by ipqqqug@qq.com
