package searchindex

import "github.com/eukarya-inc/reearth-plateauview/server/cms"

type Status string

const (
	StatusReady      Status = "未実行"
	StatusProcessing Status = "実行中"
	StatusOK         Status = "完了"
	StatusError      Status = "エラー"
)

type Item struct {
	ID string `json:"id,omitempty" cms:"id"`
	// asset: bldg
	Bldg []string `json:"bldg,omitempty" cms:"bldg,asset"`
	// asset: search_index
	SearchIndex []string `json:"search_index,omitempty" cms:"search_index,asset"`
	// select: search_index_status: 未実行, 実行中, 完了, エラー
	SearchIndexStatus Status `json:"search_index_status,omitempty" cms:"search_index_status,select"`
}

func (i Item) Fields() (fields []cms.Field) {
	item := &cms.Item{}
	cms.Marshal(i, item)
	return item.Fields
}

func ItemFrom(item cms.Item) (i Item) {
	item.Unmarshal(&i)
	return
}
