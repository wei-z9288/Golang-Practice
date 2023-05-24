package filesys

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

var basepath string

func TestMain(m *testing.M) {
	_, filename, _, _ := runtime.Caller(0)
	basepath = path.Join(path.Dir(filename), "..")
	exitVal := m.Run()
	os.Exit(exitVal)
}

func absPath(filepath string) string {
	return path.Join(basepath, filepath)
}

func TestCheckFile(t *testing.T) {
	cases := []struct {
		filepath string
		expects  bool
	}{
		{absPath("missing-file.txt"), false},
		{absPath("file/empty.txt"), true},
		{absPath("file/line.txt"), true},
		{absPath("file/myfile.txt"), true},
		{absPath("file/temp"), false},
	}

	f := File("")
	for _, testcase := range cases {

		f = File(testcase.filepath).CheckFile()
		assert.Equal(t, f.Status, testcase.expects, fmt.Sprintf("Failed test with %v", f.Filename))
	}
}

func TestExists(t *testing.T) {
	cases := []struct {
		filepath string
		expects  bool
	}{
		{absPath("missing-file.txt"), false},
		{absPath("file/myfile.txt"), true},
	}

	f := File("")
	for _, testcase := range cases {

		f = File(testcase.filepath).Exists()
		assert.Equal(t, f.Status, testcase.expects, fmt.Sprintf("Failed test with %v", f.Filename))
	}
}

func TestCheckNotDir(t *testing.T) {
	f := File(absPath("file")).CheckNotDir()
	assert.Equal(t, f.Status, false, fmt.Sprintf("Failed test with %v", f.Filename))
}
