/*
@Time : 2020/4/24 2:16 下午
@Author : mmy83
@File : jwttoken.go
@Software: GoLand
*/

package jwtauth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"leonote/app/model"
	"leonote/config"
	"log"
	"time"
)


var JwtToken *jwtToken

func init() {
	JwtToken = &jwtToken{
		SigningAlgorithm:config.CfgJwt.GetString("SigningAlgorithm"),
		Key: []byte(config.CfgJwt.GetString("key")),
		Timeout: config.CfgJwt.GetInt64("timeout"),
		MaxRefresh: config.CfgJwt.GetInt64("timeout"),
	}
}

type jwtToken struct {
	SigningAlgorithm string
	Key []byte
	Timeout int64
	MaxRefresh int64
}

type JWTClaims struct {  // token里面添加用户信息，验证token后可能会用到用户信息
	jwt.StandardClaims
	ID int64 `json:"id"`
	UserName string `json:"username"`
	NickName string `json:"nickname"`
	IsAdmin int64 `json:"isadmin"`
}

func (jt *jwtToken) CreateTokenString(user model.User) (string,time.Time,error) {

	claims := JWTClaims{
		ID: user.ID,
		UserName: user.UserName,
		NickName: user.NickName,
		IsAdmin: user.IsAdmin,
	}
	claims.IssuedAt = time.Now().Unix()

	expires := time.Now().Add(time.Second * time.Duration(jt.Timeout))
	claims.ExpiresAt = expires.Unix()

	token := jwt.NewWithClaims(jwt.GetSigningMethod(jt.SigningAlgorithm), claims)
	tokenString, err := token.SignedString(jt.Key)
	if err != nil {
		return "",expires,err
	}
	return tokenString,expires,nil
}

func (jt *jwtToken) TokenVerify(tokenString string) bool {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jt.Key, nil
	})
	if err != nil {
		log.Printf("parse token err: %s\n",err)
		return false
	}
	if err := token.Claims.Valid(); err != nil {
		log.Printf("Valid token err: %s\n",err)
		return false
	}
	return true
}

func (jt *jwtToken) ParseToken(tokenString string) (*JWTClaims,error){
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jt.Key, nil
	})
	if err != nil {
		return &JWTClaims{},err
	}
	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		return &JWTClaims{},errors.New("token.Claims err")
	}
	if err := token.Claims.Valid(); err != nil {
		return &JWTClaims{},err
	}
	return claims,nil
}

