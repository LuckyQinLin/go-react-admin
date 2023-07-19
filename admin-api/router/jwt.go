package router

import (
	"admin-api/app/models/response"
	"admin-api/app/models/vo"
	"admin-api/core"
	"admin-api/internal/gin"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
)

// authPath 判断当前路径是否在白名单中，在的话，直接放行
func authPath(list []string, path string) bool {
	for _, item := range list {
		if item == path {
			return true
		}
	}
	return false
}

// next 执行下一个handler
func next(c *gin.Context, claims *vo.UserClaims) {
	if c.Request.URL.Path == "/api/ws" {
		c.Set(vo.WSHeaderKey, true)
	}
	c.Set(vo.HeaderUserKey, claims.UserId)
	c.Next()
	c.Set(vo.HeaderUserKey, "")
}

// JwtMiddle 认证配置
func JwtMiddle() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			parseJwt = func(tokenClaims string) (claims *vo.UserClaims, isExpired bool, err error) {
				var (
					t  *jwt.Token
					ok bool
				)
				if t, err = jwt.ParseWithClaims(tokenClaims, &vo.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
					return []byte(core.Config.Jwt.SecretKey), nil
				}); err != nil {
					return nil, false, err
				}
				if !t.Valid {
					return nil, true, nil
				}
				if claims, ok = t.Claims.(*vo.UserClaims); ok {
					return claims, false, nil
				}
				return nil, false, errors.New("Token解析或非法Token")
			} // 解析Token
			whiteList = core.Config.Web.Whites()
			path      = c.Request.URL.Path
			claims    *vo.UserClaims
			isExpired bool
			auth      string
			err       error
		)
		if strings.HasPrefix(path, "/admin") || authPath(whiteList, path) {
			core.Log.Info("当前请求路径 => [%s]不做拦截", path)
			c.Next()
		} else {
			if path == "/api/ws" {
				auth = c.DefaultQuery("token", "")
			} else {
				auth = c.GetHeader(vo.AuthHeader)
			}
			if auth == "" {
				c.Abort()
				core.Log.Error("当前请求路径 => [%s], 认证信息不存在", path)
				c.JSON(http.StatusUnauthorized, response.Result(response.AuthNotExist, "认证信息不存在"))
				return
			}
			// 解析Token
			if claims, isExpired, err = parseJwt(auth); err != nil {
				c.Abort()
				core.Log.Info("当前请求路径 => [%s], 认证信息错误 => [%s]", path, err.Error())
				c.JSON(http.StatusUnauthorized, response.Result(response.AuthFail, "认证信息错误"))
				return
			}
			if isExpired {
				c.Abort()
				core.Log.Info("当前请求路径 => [%s], 认证信息过期 => [%s]", path, err.Error())
				c.JSON(http.StatusUnauthorized, response.Result(response.TokenTimeOut, "认证信息过期"))
				return
			}
			next(c, claims)
		}
	}
}
