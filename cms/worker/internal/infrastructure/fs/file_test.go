package fs

import (
	"context"
	"io"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestNewFile(t *testing.T) {
	f, err := NewFile(mockFs(), "")
	assert.NoError(t, err)
	assert.NotNil(t, f)
}

func mockFs() afero.Fs {
	files := map[string]string{"assets/aaa.txt": "aaa", "assets/bbb.txt": "bbb"}

	fs := afero.NewMemMapFs()
	for name, content := range files {
		f, _ := fs.Create(name)
		_, _ = f.WriteString(content)
		_ = f.Close()
	}
	return fs
}

func Test_fileRepo_Read(t *testing.T) {
	f, _ := NewFile(mockFs(), "")

	r, size, err := f.Read(context.Background(), "assets/aaa.txt")
	assert.NoError(t, err)
	assert.True(t, size > 0)
	buf := make([]byte, len([]byte("aaa")))
	_, err = r.ReadAt(buf, 0)

	assert.NoError(t, err)
	assert.Equal(t, "aaa", string(buf))
	assert.NoError(t, r.Close())

}

func Test_fileRepo_Upload(t *testing.T) {
	fs := mockFs()
	f, _ := NewFile(fs, "")

	u, err := f.Upload(context.Background(), "assets/ccc.txt")
	assert.NoError(t, err)

	byte := []byte("ccc")
	_, err = u.Write(byte)
	assert.NoError(t, err)

	f2, err := fs.Open("assets/ccc.txt")
	assert.NoError(t, err)
	c, _ := io.ReadAll(f2)
	assert.Equal(t, string(byte), string(c))
}
