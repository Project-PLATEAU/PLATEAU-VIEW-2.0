package fme

import (
	"net/url"
)

type ConversionRequest struct {
	ID     string
	Target string
	// JGD2011平面直角座標第1～19系のEPSGコード（6669〜6687）
	PRCS string
	// 政令指定都市を分割するかしないか
	DevideODC bool
	// 品質検査パラメータファイル
	QualityCheckParams string
	// 品質検査を行うか
	QualityCheck bool
}

func (r ConversionRequest) Query() url.Values {
	q := url.Values{}
	q.Set("id", r.ID)
	q.Set("target", r.Target)
	if r.PRCS != "" {
		q.Set("prcs", r.PRCS)
	}
	if !r.DevideODC {
		q.Set("divide_odc", "false")
	}
	if r.QualityCheckParams != "" {
		q.Set("config", r.QualityCheckParams)
	}
	return q
}

func (r ConversionRequest) Name() string {
	if r.QualityCheck {
		return "plateau2022-cms/quality-check-and-convert-all"
	}
	return "plateau2022-cms/convert-all"
	// only quality check: plateau2022-cms/quality-check
}
