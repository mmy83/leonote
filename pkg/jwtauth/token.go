/*
@Time : 2020/4/22 3:29 下午
@Author : mmy83
@File : token.go
@Software: GoLand
*/

package jwtauth

//import (
//	"github.com/dgrijalva/jwt-go"
//	"leonote/config"
//	"log"
//	"time"
//)
//
//var JwtToken *jwtToken
//
//func init() {
//	JwtToken = &jwtToken{
//		SigningAlgorithm:"HS256",
//		Key: []byte(config.CfgJwt.GetString("key")),
//		Timeout: time.Hour,
//		MaxRefresh: time.Hour,
//	}
//}
//
//type jwtToken struct {
//	SigningAlgorithm string
//	Key []byte
//	Timeout time.Duration
//	MaxRefresh time.Duration
//
//}
//
//func (ja *jwtToken) signedString(token *jwt.Token) (string, error) {
//	var tokenString string
//	var err error
//	tokenString, err = token.SignedString(ja.Key)
//	return tokenString, err
//}
//
//
//func (ja *jwtToken)CreateToken(data jwt.MapClaims) (string,time.Time,error){
//
//	token := jwt.New(jwt.GetSigningMethod(ja.SigningAlgorithm))
//	claims := token.Claims.(jwt.MapClaims)
//
//	for k,v := range data{
//		claims[k] = v
//	}
//
//	expire := time.Now().UTC().Add(ja.Timeout)
//	claims["exp"] = expire.Unix()
//	claims["iat"] = time.Now().Unix()
//
//	tokenString, err := ja.signedString(token)
//	if err != nil {
//		return "", time.Time{}, err
//	}
//
//	return tokenString, expire, nil
//}
//
//func (ja *jwtToken) parseTokenString(token string) (*jwt.Token, error) {
//	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
//		if jwt.GetSigningMethod(ja.SigningAlgorithm) != t.Method {
//			return nil, ErrInvalidSigningAlgorithm
//		}
//		return ja.Key, nil
//	})
//}
//
//func (ja *jwtToken)IsExpire(token *jwt.Token) bool  {
//
//	if int64(token.Claims.(jwt.MapClaims)["exp"].(float64)) < time.Now().Unix() {
//		return true
//	}
//
//	return false
//}
//
//func (ja *jwtToken) GetClaimsFromToken(token *jwt.Token) (jwt.MapClaims, error) {
//
//	log.Printf("token.Claims: %s\n",token.Claims)
//	claims := jwt.MapClaims{}
//	for key, value := range token.Claims.(jwt.MapClaims) {
//		claims[key] = value
//	}
//	return claims, nil
//}
//
//
