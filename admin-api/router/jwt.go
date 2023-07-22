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
	"time"
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

// JwtMiddle 认证配置
func JwtMiddle() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			parseJwt = func(tokenClaims string) (claims *vo.UserClaims, isExpired bool, err error) {
				var (
					t  *jwt.Token
					ok bool
				)
				if t, err = jwt.ParseWithClaims(tokenClaims, &vo.UserClaims{}, func(token *jwt.Token) (any, error) {
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
			newToken  string
			err       error
		)
		// 后续打包后放过 /admin开头的前端路由
		if strings.HasPrefix(path, "/admin") || authPath(whiteList, path) {
			core.Log.Info("当前请求路径[%s]不做拦截", path)
			c.Next()
		} else {
			// 获取请求头数据
			if auth = c.GetHeader(vo.AuthHeader); auth == "" {
				c.Abort()
				core.Log.Error("当前请求路径[%s], 认证信息不存在", path)
				c.JSON(http.StatusUnauthorized, response.Fail(response.AuthNotExist))
				return
			}
			// 解析Token
			if claims, isExpired, err = parseJwt(auth); err != nil {
				c.Abort()
				core.Log.Info("当前请求路径[%s], 认证信息错误[%s]", path, err.Error())
				c.JSON(http.StatusUnauthorized, response.Fail(response.AuthFail))
				return
			}
			// 判断是否过期
			if isExpired {
				c.Abort()
				core.Log.Info("当前请求路径[%s], 认证信息过期[%s]", path, err.Error())
				c.JSON(http.StatusUnauthorized, response.Fail(response.TokenTimeOut))
				return
			}
			// 续期
			if (claims.ExpiresAt.Sub(time.Now()).Minutes() * 3) <= float64(core.Config.Jwt.ExpiresTime) {
				core.Log.Info("Token剩余时间小于1/3，系统进行续期操作", err.Error())
				claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Duration(core.Config.Jwt.ExpiresTime) * time.Minute))
				t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
				if newToken, err = t.SignedString([]byte(core.Config.Jwt.SecretKey)); err != nil {
					core.Log.Info("Token续期失败[%s]", err.Error())
				}
				// 将新的Token写入Header中
				c.Header("NewToken", newToken)
			}
			c.Set(vo.ClaimsInfo, claims)
			c.Next()
			c.Set(vo.ClaimsInfo, "")
		}
	}
}
