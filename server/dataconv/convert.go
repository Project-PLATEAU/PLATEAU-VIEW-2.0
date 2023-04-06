package dataconv

import (
	"bytes"
	_ "embed"
	"fmt"
	"image/color"
	"net/http"

	geojson "github.com/paulmach/go.geojson"
	"github.com/samber/lo"
	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/renderers"
	"github.com/vincent-petithory/dataurl"
)

const (
	wallHeight           = 100
	wallImageName        = "yellow_gradient.png"
	billboardImageDir    = "billboard_image"
	billboardPaddingH    = 16.0
	billboardPaddingV    = 8.0
	billboradFontSize    = 72.0
	billboardRadius      = 30.0
	billboardLineWidth   = 2.0
	billboardLineHeight  = 100.0
	billbordHeightMargin = 55.0
)

var (
	//go:embed assets/yellow_gradient.png
	wallImage        []byte
	wallImageDataURL = dataurl.New(wallImage, http.DetectContentType(wallImage)).String()

	billboardBgColor   = color.RGBA{R: 0, G: 189, B: 189, A: 255} // #00BDBD
	billboardTextColor = color.White

	//go:embed assets/NotoSansJP-Light.otf
	billboardFontData   []byte
	billboardFontFamily = canvas.NewFontFamily("Noto Sans Japanese")
	billboardFontStyle  = canvas.FontLight
)

func init() {
	billboardFontFamily.MustLoadFont(billboardFontData, 0, billboardFontStyle)
}

// ConvertLandmark は国土基本情報を基に作成されたランドマーク・鉄道駅GeoJSONデータをPLATEAU VIEW用のCZMLに変換します。
func ConvertLandmark(fc *geojson.FeatureCollection, id string) (any, error) {
	packets := make([]any, 0, len(fc.Features))
	for i, f := range fc.Features {
		if len(f.Geometry.Point) < 2 {
			continue
		}

		name, _ := f.PropertyString("名称")
		if name == "" {
			name, _ = f.PropertyString("駅名")
		}
		if name == "" {
			continue
		}

		height, _ := f.PropertyFloat64("高さ")
		if len(f.Geometry.Point) == 2 {
			f.Geometry.Point = append(f.Geometry.Point, height+billbordHeightMargin)
		} else if height > 0 {
			f.Geometry.Point[2] = height + billbordHeightMargin
		}

		image, err := GenerateLandmarkImage(name)
		if err != nil {
			return nil, err
		}

		packets = append(packets, map[string]any{
			"id":          fmt.Sprintf("%s_%d", id, i),
			"name":        name,
			"description": name,
			"billboard": map[string]any{
				"eyeOffset": map[string]any{
					"cartesian": []int{0, 0, 0},
				},
				"horizontalOrigin": "CENTER",
				"image":            dataurl.New(image, http.DetectContentType(image)).String(),
				"pixelOffset": map[string]any{
					"cartesian2": []int{0, 0},
				},
				"scale":          0.5,
				"show":           true,
				"verticalOrigin": "BOTTOM",
				"sizeInMeters":   true,
			},
			"position": map[string]any{
				"cartographicDegrees": lo.Map(f.Geometry.Point, func(p float64, _ int) any { return p }),
			},
			"properties": processProperties(f.Properties),
		})
	}

	return czml(id, packets...), nil
}

// GenerateLandmarkImage はランドマーク用の画像を生成します。
func GenerateLandmarkImage(name string) ([]byte, error) {
	face := billboardFontFamily.Face(billboradFontSize, billboardTextColor, billboardFontStyle, canvas.FontNormal)
	text := canvas.NewTextLine(face, name, canvas.Left)
	textBounds := text.Bounds()
	text2 := canvas.NewTextBox(face, name, textBounds.W+billboardPaddingH*2, textBounds.H+billboardPaddingV*2, canvas.Center, canvas.Middle, 0, 0)

	w := textBounds.W + billboardPaddingH*2
	h := textBounds.H + billboardPaddingV*2 + billboardLineHeight
	c := canvas.New(w, h)
	ctx := canvas.NewContext(c)
	ctx.SetCoordSystem(canvas.CartesianIV)

	ctx.SetStrokeWidth(0)
	ctx.SetFillColor(billboardBgColor)
	ctx.DrawPath(0, 0, canvas.RoundedRectangle(w, textBounds.H+billboardPaddingV*2, billboardRadius))

	ctx.SetStrokeWidth(billboardLineWidth)
	ctx.SetStrokeColor(billboardBgColor)
	ctx.DrawPath(w/2, textBounds.H+billboardPaddingV*2, canvas.Line(0, h-textBounds.H+billboardPaddingV*2))

	ctx.DrawText(0, 0, text2)

	b := bytes.NewBuffer(nil)
	if err := renderers.PNG()(b, c); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

// ConvertBorder は国土数値情報を基に作成された行政界GeoJSONデータをPLATEAU VIEW用のCZMLに変換します。
func ConvertBorder(fc *geojson.FeatureCollection, id string) (any, error) {
	packets := make([]any, 0, len(fc.Features))
	for i, f := range fc.Features {
		var possrc [][][]float64
		if len(f.Geometry.MultiLineString) > 0 {
			possrc = f.Geometry.MultiLineString
		} else if len(f.Geometry.LineString) > 0 {
			possrc = [][][]float64{f.Geometry.LineString}
		} else if len(f.Geometry.MultiPolygon) > 0 {
			possrc = lo.FilterMap(f.Geometry.MultiPolygon, func(p [][][]float64, _ int) ([][]float64, bool) {
				if len(p) == 0 {
					return nil, false
				}
				return p[0], true
			})
		} else if len(f.Geometry.Polygon) > 0 {
			possrc = f.Geometry.Polygon
		}

		if len(possrc) == 0 {
			continue
		}

		for j, pos := range possrc {
			positions := lo.FlatMap(pos, func(p []float64, _ int) []float64 {
				if len(p) < 2 {
					return nil
				}
				return []float64{p[0], p[1], wallHeight}
			})

			packets = append(packets, map[string]any{
				"id": fmt.Sprintf("%s_%d_%d", id, i+1, j+1),
				"wall": map[string]any{
					"material": map[string]any{
						"image": map[string]any{
							"image":       wallImageDataURL,
							"repeat":      true,
							"transparent": true,
						},
					},
					"positions": map[string]any{
						"cartographicDegrees": lo.Map(positions, func(p float64, _ int) any { return p }),
					},
				},
				"properties": processProperties(f.Properties),
			})
		}
	}

	return czml(id, packets...), nil
}

func czml(name string, packets ...any) any {
	return append(
		[]any{
			map[string]any{
				"id":      "document",
				"name":    name,
				"version": "1.0",
			},
		},
		packets...,
	)
}

func processProperties(p map[string]any) map[string]any {
	m := make(map[string]any, len(p))
	for k, v := range p {
		if v != nil {
			m[k] = v
		}
	}
	return m
}
