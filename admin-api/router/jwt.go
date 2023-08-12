package router

import (
	"admin-api/app/models/response"
	"admin-api/app/models/vo"
	"admin-api/core"
	"admin-api/internal/gin"
	"encoding/json"
	"errors"
	"fmt"
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
			parseJwt = func(token string) (*vo.UserClaims, bool, float64, error) {
				var (
					key    string
					value  string
					expire float64
					user   vo.UserClaims
					err    error
				)
				key = fmt.Sprintf("%s:%s", vo.RedisToken, token)
				// 判断是否过期
				if expire = core.Cache.IsExpire(key); expire == -2 {
					return nil, true, 0, nil
				}
				// 获取数据
				if value, err = core.Cache.GetKey(key); err != nil {
					return nil, false, 0, errors.New("获取用户数据失败")
				}
				// 序列化
				if err = json.Unmarshal([]byte(value), &user); err != nil {
					return nil, false, 0, errors.New("解析用户数据失败")
				}
				return &user, false, expire, nil
			} // 解析Token

			whiteList = core.Config.Web.Whites()
			path      = c.Request.URL.Path
			user      *vo.UserClaims
			isExpired bool
			expired   float64
			auth      string
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
			if user, isExpired, expired, err = parseJwt(auth); err != nil {
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
			if expired*3 <= float64(core.Config.Jwt.ExpiresTime) {
				core.Log.Info("Token剩余时间小于1/3，系统进行续期操作")
				key := fmt.Sprintf("%s:%s", vo.RedisToken, auth)
				if _, err = core.Cache.KeyExpired(
					key,
					time.Duration(core.Config.Jwt.ExpiresTime)*time.Minute,
				); err != nil {
					core.Log.Info("认证Token[%s]续期失败", key, err.Error())
				}
			}
			c.Set(vo.ClaimsInfo, user)
			c.Next()
			c.Set(vo.ClaimsInfo, "")
		}
	}
}
