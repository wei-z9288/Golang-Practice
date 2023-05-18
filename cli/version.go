package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version string

// 將 versioncmd 命令添加到了 rootcmd
func init() {
	rootcmd.AddCommand(versionCmd)
}

func SetVersion(newVersion string) {
	Version = newVersion
}

func GetVersion() string {
	return Version
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show current version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("fops " + Version)
	},
}

//Run (1)接收正在執行的指令的指標
//    (2)args []string: 命令行中輸入的參數。這個切片包含了除了命令名稱以外的所有參數。
