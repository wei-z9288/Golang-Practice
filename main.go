package main

import (
	"bytes"
	"fmt"
	algo "fops/algo"
	cli "fops/cli"
	filesys "fops/filesys"
	"io/ioutil"
	"os"
	"path"
)

var basepath string

func absPath(filepath string) string {
	return path.Join(basepath, filepath)
}
func main() {
	//algo test
	var a algo.Algorithm = algo.GetMD5()
	data := []byte("hello world")
	fmt.Printf("Type: %s\n", a.GetType())
	fmt.Printf("Checksum: %s\n", a.GetChecksum(data))

	var b algo.Algorithm = algo.GetSHA1()
	fmt.Printf("Type: %s\n", b.GetType())
	fmt.Printf("Checksum: %s\n", b.GetChecksum(data))
	// file verification test
	cases := []string{
		"missing-file.txt",
		"file/empty.txt",
		"file/line.txt",
		"file/myfile.txt",
		"file/temp",
		"file/",
		"file",
	}
	for key, testcase := range cases {
		fs := filesys.File(testcase).CheckFile().CheckNotBinary()
		fmt.Printf("Test case %d\n", key)
		fmt.Printf("File Name: %s\n", fs.Filename)
		fmt.Printf("File Status: %t\n", fs.Status)
		fmt.Printf("Error Message: %v\n", fs.Err)

	}

	//cmd version test

	//cmd line count test
	cmd := cli.LinecountCmd()
	buffer := bytes.NewBufferString("")
	cmd.SetOut(buffer)
	cmd.SetArgs([]string{"--file", absPath("file/myfile.txt")})
	cmd.Execute()
	message, err := ioutil.ReadAll(buffer)
	fmt.Println(message)
	fmt.Println(err)
	//cmd checksum
	contents, _ := os.ReadFile("file/myfile.txt")
	fmt.Printf("myfile.txt md5 Checksum: %s\n", cli.GetChecksum(contents, "md5"))
	fmt.Printf("myfile.txt sha1 Checksum: %s\n", cli.GetChecksum(contents, "sha1"))
	fmt.Printf("myfile.txt sha256 Checksum: %s\n", cli.GetChecksum(contents, "sha256"))
}
