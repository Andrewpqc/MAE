package auth

import (
	"github.com/kataras/iris"
	//"github.com/muxiyun/MAE/models"
	"github.com/muxiyun/MAE/db"
	"encoding/base64"
	"fmt"
	"log"

	"strings"
	"github.com/muxiyun/MAE/models"
	"os"
	"crypto/md5"
	"io"
)

//auth中间件,检查来访请求是否持有正确的token
func AuthTokenCheckerMiddleware(ctx iris.Context){
	//获取token
	basicAuth_token:=ctx.Request().Header.Get("Authorization")[6:]//身份认证
	result,err:=base64.StdEncoding.DecodeString(basicAuth_token)
	if err!=nil{
		log.Fatal("decode basic auth token failed!")
	}
	results:=strings.Split(string(result),":")
	username:=results[0]
	password:=results[1]
	key := []byte(os.Getenv("SECRETKEY"))
	passwordhash := md5.New()
	io.WriteString(passwordhash, password)
	hash := string(passwordhash.Sum(key))
	u:=models.User{Username:username,PasswordHash:hash}
	if ok,err:=db.DBEngine.Get(&u);!ok{
		fmt.Println(err)
	}
	token:=ctx.Request().FormValue("token")
	err=u.CheckToken(token,u.Role)
	if err!=nil{
		ctx.StatusCode(403)
		ctx.WriteString("token error")
	}
	ctx.Next()
}


