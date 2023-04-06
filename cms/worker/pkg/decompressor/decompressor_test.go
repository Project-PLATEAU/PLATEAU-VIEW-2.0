package decompressor

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

// Buffer implements io.ReadAtCloser and io.WriteCloser
type Buffer struct {
	bytes.Buffer
}

func (*Buffer) Close() error {
	return nil
}

func TestNewUnzipper(t *testing.T) {
	zf := lo.Must(os.Open("testdata/test.zip"))
	fInfo := lo.Must(zf.Stat())

	wFn := func(name string) (io.WriteCloser, error) {
		return &Buffer{bytes.Buffer{}}, nil
	}
	_, err := New(zf, fInfo.Size(), "zip", wFn)
	assert.NoError(t, err)

	_, err = New(zf, 0, "zip", wFn)
	assert.Error(t, err)

	zf = lo.Must(os.Open("testdata/test.7z"))
	_, err = New(zf, fInfo.Size(), "7z", wFn)
	assert.NoError(t, err)

	_, err = New(zf, 0, "7z", wFn)
	assert.Error(t, err)

	f := lo.Must(os.Open("testdata/test1.txt"))
	fInfo2 := lo.Must(f.Stat())
	_, err2 := New(f, fInfo2.Size(), "txt", wFn)
	// txt is not unsupported
	assert.Same(t, ErrUnsupportedExtention, err2)
}

// func TestDecompressor_Decompress(t *testing.T) {
// 	zf := lo.Must(os.Open("testdata/test.zip"))
// 	szf := lo.Must(os.Open("testdata/test.7z"))

// 	fInfo, err := zf.Stat()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	expectedFiles := map[string][]byte{
// 		"test1.txt": []byte("hello1"),
// 		"test2.txt": []byte("hello2"),
// 	}

// 	// map of buffers which will keep unzipped data
// 	files := map[string]*Buffer{
// 		"test1.txt": {bytes.Buffer{}},
// 		"test2.txt": {bytes.Buffer{}},
// 	}

// 	// normal scenario
// 	uz, err := New(zf, fInfo.Size(), "zip", func(name string) (io.WriteCloser, error) {
// 		return files[name], nil
// 	})
// 	require.NoError(t, err)

// 	assert.NoError(t, uz.Decompress("testdata"))
// 	for k, v := range files {
// 		assert.Equal(t, expectedFiles[k], v.Bytes())
// 	}

// 	// Redefine files because buffer overwriting will occur.
// 	files = map[string]*Buffer{
// 		"test1.txt": {bytes.Buffer{}},
// 		"test2.txt": {bytes.Buffer{}},
// 	}
// 	uz2, err := New(szf, fInfo.Size(), "7z", func(name string) (io.WriteCloser, error) {
// 		return files[name], nil
// 	})
// 	require.NoError(t, err)
// 	assert.NoError(t, uz2.Decompress("testdata"))
// 	for k, v := range files {
// 		assert.Equal(t, expectedFiles[k], v.Bytes())
// 	}

// 	// exception: test if  wFn's error is same as what Unzip returns
// 	uz, err = New(zf, fInfo.Size(), "zip", func(name string) (io.WriteCloser, error) {
// 		return nil, errors.New("test")
// 	})
// 	require.NoError(t, err)
// 	assert.Equal(t, errors.New("test"), uz.Decompress("testdata"))

// 	uz, err = New(szf, fInfo.Size(), "7z", func(name string) (io.WriteCloser, error) {
// 		return nil, errors.New("test")
// 	})
// 	require.NoError(t, err)
// 	assert.Equal(t, errors.New("test"), uz.Decompress("testdata"))
// }
