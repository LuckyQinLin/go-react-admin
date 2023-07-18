package cmd

import (
	"admin-api/internal/cobra"
	"admin-api/router"
)

// serverCmd 启动子命令
var serverCmd = &cobra.Command{
	Use:               "server",
	Short:             "启动应用",
	Long:              `启动应用`,
	DisableAutoGenTag: false,
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		println(banner)
		router.Run(Param.port)
	},
}

func init() {
	serverCmd.Flags().Int64VarP(&Param.port, "port", "p", 8070, "端口")
	serverCmd.Flags().StringVarP(&Param.config, "config", "f", "", "配置文件,默认($HOME/.admin/config.toml)")
	rootCmd.AddCommand(serverCmd)
}
