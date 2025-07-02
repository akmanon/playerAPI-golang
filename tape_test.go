package poker

import (
	"io"
	"testing"

	"github.com/alecthomas/assert"
)

func TestTape_Write(t *testing.T) {
	file, clean := createTempFile(t, "12344")
	defer clean()
	tape := &tape{file}

	tape.Write([]byte("abc"))
	file.Seek(0, io.SeekStart)
	newFileContent, _ := io.ReadAll(file)

	got := string(newFileContent)
	want := "abc"

	assert.Equal(t, got, want, "got %q, want %q", got, want)
}
