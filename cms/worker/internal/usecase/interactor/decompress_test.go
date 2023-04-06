package interactor

import (
	"context"
	"io"
	"os"
	// "testing"

	// wfs "github.com/reearth/reearth-cms/worker/internal/infrastructure/fs"
	// "github.com/reearth/reearth-cms/worker/internal/usecase/gateway"
	"github.com/reearth/reearth-cms/worker/pkg/asset"

	"github.com/samber/lo"
	"github.com/spf13/afero"
	// "github.com/stretchr/testify/assert"
	// "github.com/stretchr/testify/require"
)

// func TestUsecase_Decompress(t *testing.T) {
// 	fs := mockFs()
// 	mCMS := NewCMS()
// 	fileGateway, err := wfs.NewFile(fs, "")
// 	require.NoError(t, err)

// 	uc := NewUsecase(gateway.NewGateway(fileGateway, mCMS))

// 	assert.NoError(t, uc.Decompress(context.Background(), "aaa", "test.zip"))

// 	f := lo.Must(fs.Open("test/test1.txt"))
// 	content := lo.Must(io.ReadAll(f))
// 	_ = f.Close()
// 	assert.Equal(t, "hello1", string(content))

// 	f = lo.Must(fs.Open("test/test2.txt"))
// 	content = lo.Must(io.ReadAll(f))
// 	_ = f.Close()
// 	assert.Equal(t, "hello2", string(content))

// 	// unsupported extenstion doesn't return error
// 	assert.NoError(t, uc.Decompress(context.Background(), "aaa", "test.tar.gz"))
// }

func mockFs() afero.Fs {
	fs := afero.NewMemMapFs()
	zf := lo.Must(os.Open("testdata/test.zip"))

	zf2 := lo.Must(fs.Create("test.zip"))
	_ = lo.Must(io.Copy(zf2, zf))
	_ = zf2.Close()

	zf3 := lo.Must(fs.Create("test.tar.gz"))
	_ = lo.Must(io.Copy(zf3, zf))
	_ = zf3.Close()

	_ = zf.Close()

	return fs
}

type mockCMS struct {
}

func NewCMS() *mockCMS {
	return &mockCMS{}
}

func (c *mockCMS) NotifyAssetDecompressed(ctx context.Context, assetId string, status *asset.ArchiveExtractionStatus) error {
	return nil
}
