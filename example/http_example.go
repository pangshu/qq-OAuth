package main

import (
	"fmt"
	"net/http"
	"github.com/pangshu/qq-OAuth"
	"log"
)

var state = qq_OAuth.NewUtils().RandString(8)
var appId = ""
var appSecret = ""
var callback = ""
var scope = ""
// login action
func Login(res http.ResponseWriter, req *http.Request) {
	oAuth := qq_OAuth.NewOAuth(appId, appSecret, callback, scope)
	loginUrl := oAuth.GetAuthorURL(state)

	http.Redirect(res, req, loginUrl, 302)
}

// callback action
func Callback(res http.ResponseWriter, req *http.Request)  {
	req.ParseForm()
	reqState := req.Form.Get("state")
	if reqState != state {
		res.Write([]byte("error"))
	}

	authCode := req.Form.Get("code")

	oAuth := qq_OAuth.NewOAuth(appId, appSecret, callback, scope)
	openid,_ := oAuth.Access(authCode)
	fmt.Print(openid)
	userInfo,_,_ := oAuth.GetUserInfo()
	fmt.Print(userInfo)
}

func main()  {

	http.HandleFunc("/qq/login", Login)
	http.HandleFunc("/qq/callback", Callback)

	err := http.ListenAndServe(":8088", nil)
	if err != nil {
		log.Println(err.Error())
	}
}