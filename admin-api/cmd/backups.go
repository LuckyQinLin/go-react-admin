package cmd

import (
	"fmt"
	"gitee.com/molonglove/goboot/cobra"
)

// backupsCmd 版本子命令
var backupsCmd = &cobra.Command{
	Use:   "backup",
	Short: "数据备份备份",
	Long:  `将当前应用关键数据做备份`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("数据备份开始......")
	},
}

func init() {
	rootCmd.AddCommand(backupsCmd)
}
