package cmd

import (
	"admin-api/internal/cobra"
	"fmt"
	"github.com/mitchellh/go-homedir"
	"os"
	"path/filepath"
)

const toml = `
# 服务配置
[web]
    # 端口
    port = 8080
    # 根路由
    context-path = "api"
    read-timeout = 10
    write-timeout = 10
    max-header-bytes = 50
    # 运行模型
    run-model = "debug"
    # 白名单
    white-list = ["/user/login", "/user/register", "/user/captcha"]
# 数据库配置
[db]
	host = "192.168.110.107"
	port = 5432
	username = "lucky"
	password = "zjnslovef2354"
	db-name = "manager_db"
	schema = "public"
# jwt配置信息
[jwt]
    issuer = "Lucky.麒麟"
    expires-time = 24
    secret-key = "lucky admin"
# 日志配置
[log]
	prefix = "admin"
`

// initCmd 初始化子命令
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "初始化应用环境",
	Long:  `初始化应用工作目录和相关配置文件`,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			homePath string
			err      error
		)
		if homePath, err = homedir.Expand(Param.home); err != nil {
			panic("初始化失败，路径解析失败")
		}
		if _, err = os.Stat(homePath); err != nil {
			if os.IsNotExist(err) {
				if err = os.Mkdir(homePath, os.ModePerm); err == nil {
					if err = os.WriteFile(filepath.Join(homePath, "config.toml"), []byte(toml), 0666); err != nil {
						panic("写入配置文件失败" + err.Error())
					}
				}
			}
		} else {
			if _, err = os.Stat(filepath.Join(homePath, "config.toml")); err != nil {
				if os.IsNotExist(err) {
					if err = os.WriteFile(filepath.Join(homePath, "config.toml"), []byte(toml), 0666); err != nil {
						panic("写入配置文件失败" + err.Error())
					}
				}
			}
		}
		fmt.Println("初始化应用环境成功")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
