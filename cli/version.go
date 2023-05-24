package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version string
var versionCmd *cobra.Command

// 將 versioncmd 命令添加到了 rootcmd
func init() {
	VersionCmd().Flags().StringP("version", "v", "", "File's version")
	rootcmd.AddCommand(VersionCmd())
}

func SetVersion(newVersion string) {
	Version = newVersion
}

func GetVersion() string {
	return Version
}

//Run (1)接收正在執行的指令的指標
//    (2)args []string: 命令行中輸入的參數。這個切片包含了除了命令名稱以外的所有參數。

func VersionCmd() *cobra.Command {
	if versionCmd == nil {
		versionCmd = &cobra.Command{
			Use:   "version",
			Short: "Show current version",
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Println("fops " + Version)
			},
		}
	}
	return versionCmd
}
