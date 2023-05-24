package cli

import (
	"bytes"
	"fmt"
	"io/ioutil"
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

func TestCount(t *testing.T) {
	cases := []struct {
		filepath string
		expects  int
	}{
		{absPath("file/empty.txt"), 0},
		{absPath("file/line.txt"), 3},
		{absPath("file/myfile.txt"), 4},
	}

	for _, testcase := range cases {
		linecount, _ := count(testcase.filepath)
		assert.Equal(t, linecount, testcase.expects, fmt.Sprintf("Failed test with %v", testcase.filepath))
	}
}

func TestLinecountRun(t *testing.T) {
	cmd := LinecountCmd()
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	cmd.SetArgs([]string{"--file", absPath("file/myfile.txt")})
	cmd.Execute()
	_, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
}
