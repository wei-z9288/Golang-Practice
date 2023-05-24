package main

import (
	"fmt"
	algo "fops/algo"
	filesys "fops/filesys"
)

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
	// cli.SetVersion("1.0.0")
	// cmd := cli.VersionCmd()
	// buffer := bytes.NewBufferString("")
	// cmd.SetOut(buffer)
	// cmd.SetArgs([]string{"-v"})
	// cmd.Execute()
	// output := buffer.String()
	// fmt.Println(output)
	// cli.SetVersion("1.0.0") // 设置 Version 变量
	// v := cli.GetVersion()
	// fmt.Println(v)

	// cmd := cli.VersionCmd()
	// cmd.SetArgs([]string{"version"})    // 获取版本命令
	// buffer := bytes.NewBufferString("") // 创建一个新的 Buffer
	// cmd.SetOut(buffer)                  // 将 Buffer 设置为命令的输出
	// cmd.Execute()                       // 执行命令
	// output := buffer.String()           // 从 Buffer 中获取命令的输出
	// fmt.Println(output)                 // 打印输出

}
