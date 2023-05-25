package cli

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetChecksum(t *testing.T) {
	cases := []struct {
		filepath  string
		algorithm string
		expects   string
	}{
		{absPath("file/myfile.txt"), "md5", "810c3af7c4a46e239bdad70dd9dbb702"},
		{absPath("file/myfile.txt"), "sha1", "1a91adfcea17a987a81fd793b2c3758d0c43f51f"},
		{absPath("file/myfile.txt"), "sha256", "4a778138d86158f24577664315dfbff928b9396fb9ee9d8d6b7b3c4607644dac"},
	}

	for _, testcase := range cases {
		contents, _ := os.ReadFile(testcase.filepath)
		sum := GetChecksum(contents, testcase.algorithm)
		assert.Equal(t, sum, testcase.expects, fmt.Sprintf("Failed test with %v", testcase.filepath))
	}
}

func TestChecksumRun(t *testing.T) {
	cmd := ChecksumCmd()
	b := bytes.NewBufferString("")
	assert.Error(t, cmd.RunE(cmd, []string{}))
	cmd.SetOut(b)
	cmd.SetArgs([]string{"--file", absPath("assets/myfile.txt"), "--md5"})
	cmd.Execute()
	_, err := io.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
}
