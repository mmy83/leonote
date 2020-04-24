/*
@Time : 2020/4/22 11:49 上午
@Author : mmy83
@File : extractor.go
@Software: GoLand
*/

package jwtauth

import (
	"github.com/gin-gonic/gin"
	"strings"
)

var JwtExtractor *Extractor

func init() {
	JwtExtractor = &Extractor{
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:   "Bearer",
	}
}

type Extractor struct {
	TokenHeadName string
	TokenLookup string
}

func (e *Extractor)ExtractToken(c *gin.Context) (string ,error) {
	var token string
	var err error

	methods := strings.Split(e.TokenLookup, ",")
	for _, method := range methods {
		if len(token) > 0 {
			break
		}
		parts := strings.Split(strings.TrimSpace(method), ":")
		k := strings.TrimSpace(parts[0])
		v := strings.TrimSpace(parts[1])
		switch k {

		case "header":
			token, err = e.jwtFromHeader(c, v)
		case "query":
			token, err = e.jwtFromQuery(c, v)
		case "cookie":
			token, err = e.jwtFromCookie(c, v)
		case "param":
			token, err = e.jwtFromParam(c, v)
		default:
			token,err = "",ErrEmptyToken
		}
	}

	if err != nil {
		return "", err
	}

	return token,nil
}

func (e *Extractor) jwtFromHeader(c *gin.Context, key string) (string, error) {
	authHeader := c.Request.Header.Get(key)

	if authHeader == "" {
		return "", ErrEmptyAuthHeader
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == e.TokenHeadName) {
		return "", ErrInvalidAuthHeader
	}

	return parts[1], nil
}

func (e *Extractor) jwtFromQuery(c *gin.Context, key string) (string, error) {
	token := c.Query(key)

	if token == "" {
		return "", ErrEmptyQueryToken
	}

	return token, nil
}

func (e *Extractor) jwtFromCookie(c *gin.Context, key string) (string, error) {
	cookie, _ := c.Cookie(key)

	if cookie == "" {
		return "", ErrEmptyCookieToken
	}

	return cookie, nil
}

func (e *Extractor) jwtFromParam(c *gin.Context, key string) (string, error) {
	token := c.Param(key)

	if token == "" {
		return "", ErrEmptyParamToken
	}

	return token, nil
}




