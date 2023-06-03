package cli

import (
	"fmt"
	algo "fops/algo"
	filesys "fops/filesys"
	"os"

	"github.com/spf13/cobra"
)

var checksum *cobra.Command

// 定义全局变量用于存储flag的值
var md5Flag, sha1Flag, sha256Flag bool

func init() {

	//使用VarP函数将flag值绑定到变量
	ChecksumCmd().Flags().StringVarP(&file, "file", "f", "", "File to operate on")
	ChecksumCmd().Flags().BoolVar(&md5Flag, "md5", false, "checksum using md5")
	ChecksumCmd().Flags().BoolVar(&sha1Flag, "sha1", false, "checksum using sha1")
	ChecksumCmd().Flags().BoolVar(&sha256Flag, "sha256", false, "checksum using sha256")

	ChecksumCmd().MarkFlagRequired("file")

	rootcmd.AddCommand(ChecksumCmd())
}

func GetChecksum(value []byte, hashFn string) string {
	return algo.GetChecksum(value, hashFn)
}

func ChecksumRun(cmd *cobra.Command, args []string) error {
	//Verify the file
	fs := filesys.File(file).CheckFile()

	if fs.Err != nil {
		return fs.Err
	} else {
		contents, _ := os.ReadFile(file)
		algorithmFlag := ""
		if md5Flag {
			algorithmFlag = "md5"
		} else if sha1Flag {
			algorithmFlag = "sha1"
		} else if sha256Flag {
			algorithmFlag = "sha256"
		}

		if algorithmFlag == "" {
			return fmt.Errorf("checksum flag is missing. one of md5, sha1, sha256 is required")
		}

		fmt.Println(GetChecksum(contents, algorithmFlag))
	}

	return nil
}

func ChecksumCmd() *cobra.Command {
	if checksum == nil {
		checksum = &cobra.Command{
			Use:   "checksum",
			Short: "Generate checksum of a file",
			Long: `Checksum retrieves the cryptographic hash of a file.
		A valid text file and an algorithm flag (MD5, SHA1, SHA256) are needed.`,
			RunE: ChecksumRun,
		}
	}
	return checksum
}
