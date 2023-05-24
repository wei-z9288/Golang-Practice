package filesys

import (
	"fmt"
	"os"

	"github.com/gabriel-vasile/mimetype"
)

// 建立儲存檔案資訊的資料結構
// Status     是否符合格式
// Err        錯誤訊息
// Filename   檔名
type FileStatus struct {
	Status   bool
	Err      error
	Filename string
}

func File(filename string) FileStatus {
	return FileStatus{true, nil, filename}
}

// 每一個方法都會返回相同的物件，所以可以在該物件上連續調用多個方法。
// File(filename).Exists().CheckNotDir().CheckNotBinary()
// 先判斷是否有該檔案，在判斷是否是dir，在判斷是否是二進位檔案，中途有錯即返回前一個錯誤資訊
func (f FileStatus) CheckFile() FileStatus {
	return f.Exists().CheckNotDir()
}

func (f FileStatus) Exists() FileStatus {

	if f.Err != nil {
		return f
	}

	_, err := os.Stat(f.Filename)
	// 如果嘗試打開的文件不存在，os.IsNotExist(err)將返回 true
	if os.IsNotExist(err) {
		f.Status, f.Err = false, fmt.Errorf("error: No such file '%v'", f.Filename)
	}

	return f
}

func (f FileStatus) CheckNotDir() FileStatus {

	if f.Err != nil {
		return f
	}

	info, _ := os.Stat(f.Filename)

	if info.IsDir() {
		f.Status, f.Err = false, fmt.Errorf("error: Expected file got directory '%v'", f.Filename)
	}

	return f
}

func (f FileStatus) CheckNotBinary() FileStatus {

	if f.Err != nil {
		return f
	}

	mime, _ := mimetype.DetectFile(f.Filename)
	// Mach-O 是 macOS 和 iOS 系统上使用的二進文件的类型
	if mime.Is("application/x-mach-binary") || mime.Is("application/octet-stream") {
		f.Status, f.Err = false, fmt.Errorf("error: Cannot do linecount for binary file '%v'", f.Filename)
	}

	return f
}
