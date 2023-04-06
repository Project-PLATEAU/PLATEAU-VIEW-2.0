package sdkapi

/*
import (
	"context"
	"encoding/json"
	"testing"

	"github.com/eukarya-inc/reearth-plateauview/server/cms"
	"github.com/samber/lo"
)

func TestCMS(t *testing.T) {
	ctx := context.Background()
	cms := lo.Must(cms.New("", ""))
	c := &CMS{
		Project:              "",
		IntegrationAPIClient: cms,
	}
	res := lo.Must(c.Datasets(ctx, modelKey))
	// res := lo.Must(cms.GetItemsByKey(ctx, "", modelKey, true))
	t.Log(string(lo.Must(json.MarshalIndent(res, "", "  "))))
}

// */
