package filesys

import (
	"fmt"
	"os"

	"github.com/gabriel-vasile/mimetype"
)

// Create a data structure for storing file information
// Status     Whether it fits the format
// Err        Error message
// Filename   File Name
type FileStatus struct {
	Status   bool
	Err      error
	Filename string
}

func File(filename string) FileStatus {
	return FileStatus{true, nil, filename}
}

// Each method returns the same object, so you can call multiple methods on that object in series.
// File(filename).Exists().CheckNotDir().CheckNotBinary()
// First determine whether there is the file, then determine whether it is a dir, then determine whether it is a binary file,
// if there is an error in the process that returns the previous error message
func (f FileStatus) CheckFile() FileStatus {
	return f.Exists().CheckNotDir()
}

func (f FileStatus) Exists() FileStatus {

	if f.Err != nil {
		return f
	}

	_, err := os.Stat(f.Filename)
	// IsNotExist(err) will return true if the file you are trying to open does not exist
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
	// Mach-O is the type of binary file used on macOS and iOS systems
	if mime.Is("application/x-mach-binary") || mime.Is("application/octet-stream") {
		f.Status, f.Err = false, fmt.Errorf("error: Cannot do linecount for binary file '%v'", f.Filename)
	}

	return f
}
