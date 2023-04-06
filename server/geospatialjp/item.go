package geospatialjp

import (
	"strconv"
	"strings"

	"github.com/eukarya-inc/reearth-plateauview/server/cms"
)

type Status string

const (
	StatusReady      Status = "未実行"
	StatusProcessing Status = "実行中"
	StatusOK         Status = "完了"
	StatusError      Status = "エラー"
)

type Item struct {
	ID                  string `json:"id,omitempty" cms:"id"`
	Specification       string `json:"specification,omitempty" cms:"specification,select"`
	Prefecture          string `json:"prefecture,omitempty" cms:"prefecture,select"`
	CityName            string `json:"city_name,omitempty" cms:"city_name,text"`
	CityGML             string `json:"citygml,omitempty" cms:"citygml,asset"`
	CityGMLGeoSpatialJP string `json:"citygml_geospatialjp,omitempty" cms:"citygml_geospatialjp,asset"`
	Catalog             string `json:"catalog,omitempty" cms:"catalog,asset"`
	All                 string `json:"all,omitempty" cms:"all,asset"`
	ConversionStatus    Status `json:"conversion_status,omitempty" cms:"conversion_status,select"`
	CatalogStatus       Status `json:"catalog_status,omitempty" cms:"catalog_status,select"`
	// 公開する・公開しない
	SDKPublication string `json:"sdk_publication,omitempty" cms:"sdk_publication,select"`
}

func ItemFrom(item cms.Item) (i Item) {
	item.Unmarshal(&i)
	return
}

func (i Item) Fields() (fields []cms.Field) {
	item := &cms.Item{}
	cms.Marshal(i, item)
	return item.Fields
}

func (i Item) SpecVersion() float64 {
	d := strings.TrimSuffix(strings.TrimPrefix(i.Specification, "第"), "版")
	v, err := strconv.ParseFloat(d, 64)
	if err != nil {
		return 0
	}
	return v
}

func (i Item) IsPublicOnSDK() bool {
	return i.SDKPublication == "公開する"
}
