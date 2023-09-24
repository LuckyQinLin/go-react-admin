package cmd

import (
	_ "embed"
	"gitee.com/molonglove/goboot/cobra"
)

//go:embed banner.txt
var banner string

var Param CommandParam

// CommandParam 命令参数
type CommandParam struct {
	port     int64
	config   string
	username string
	version  string
	home     string
}

var rootCmd = &cobra.Command{
	Use:   "helper",
	Short: "帮助你快速配置服务器",
	Long:  banner,
}

func init() {
	rootCmd.PersistentFlags().Int64VarP(&Param.port, "port", "p", 8077, "端口")
	rootCmd.PersistentFlags().StringVarP(&Param.home, "home", "", "~/.admin", "工作目录")
	rootCmd.PersistentFlags().StringVarP(&Param.config, "config", "f", "~/.admin/config.toml", "配置文件")
	rootCmd.PersistentFlags().StringVarP(&Param.username, "username", "u", "admin", "系统管理员账号")
}

// Execute 执行 rootCmd 命令并检测错误
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
