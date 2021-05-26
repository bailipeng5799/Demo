package common

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"jxbdproject/model"
	"time"
)

type Myclaims struct{
	Userid int `json:"userid"`
	jwt.StandardClaims
}
var MySecret = []byte("blpblp")
//创建token
func  Createmv(user *model.User)(string,error){
	expireTime := time.Now().Add(7*24*time.Hour)
	c :=&Myclaims{
		Userid:         user.Id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt:	expireTime.Unix(),//失效时间
			Issuer: "blp",//签发人
			IssuedAt: time.Now().Unix(),//签发时间
			Subject:  "user token",	//主题
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,c)
	tokenstring,err:=token.SignedString(MySecret)
	if err!=nil{
		fmt.Println("加密失败",err)
		return "",err
	}
	return tokenstring,err
}

func ParseToken(tokenstring string)(*jwt.Token,*Myclaims,error){//解析加密后的令牌返回未加密的令牌
	//解析token
	myclaims:=&Myclaims{}
	token,err := jwt.ParseWithClaims(tokenstring,myclaims,func(token *jwt.Token)(i interface{},err error){
		return MySecret,nil
	})
	return token,myclaims,err
}