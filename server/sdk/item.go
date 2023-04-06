package sdk

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
	// asset: citygml
	CityGML string `json:"citygml,omitempty" cms:"citygml,asset"`
	// asset: max_lod
	MaxLOD string `json:"max_lod,omitempty" cms:"max_lod,asset"`
	// select: max_lod_status: 未実行, 実行中, 完了, エラー
	MaxLODStatus Status `json:"max_lod_status,omitempty" cms:"max_lod_status,select"`
	// select: 公開する・公開しない
	SDKPublication string `json:"sdk_publication,omitempty" cms:"sdk_publication,select"`
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

func (i Item) IsPublicOnSDK() bool {
	return i.SDKPublication == "公開する"
}
