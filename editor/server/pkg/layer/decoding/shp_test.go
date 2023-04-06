package decoding

import (
	"os"
	"strings"
	"testing"

	"github.com/reearth/reearth/server/pkg/shp"
	"github.com/stretchr/testify/assert"
)

var _ Decoder = &ShapeDecoder{}
var _ ShapeReader = &shp.ZipReader{}
var _ ShapeReader = &shp.Reader{}

type identityTestFunc func(*testing.T, [][]float64, []shp.Shape)
type shapeGetterFunc func(string, *testing.T) []shp.Shape
type testCaseData struct {
	points [][]float64
	tester identityTestFunc
}

var dataForReadTests = map[string]testCaseData{
	"shapetest/shapes.zip": {
		points: [][]float64{
			{10, 10},
			{5, 5},
			{0, 10},
		},
		tester: testPoint,
	},
	"shapetest/point.shp": {
		points: [][]float64{
			{10, 10},
			{5, 5},
			{0, 10},
		},
		tester: testPoint,
	},
	"shapetest/polyline.shp": {
		points: [][]float64{
			{0, 0},
			{5, 5},
			{10, 10},
			{15, 15},
			{20, 20},
			{25, 25},
		},
		tester: testPolyLine,
	},
	"shapetest/polygon.shp": {
		points: [][]float64{
			{0, 0},
			{0, 5},
			{5, 5},
			{5, 0},
			{0, 0},
		},
		tester: testPolygon,
	},
}

func testPoint(t *testing.T, points [][]float64, shapes []shp.Shape) {
	for n, s := range shapes {
		p, ok := s.(*shp.Point)
		if !ok {
			t.Fatal("Failed to type assert.")
		}
		assert.Equal(t, []float64{p.X, p.Y}, points[n])
	}
}

func testPolyLine(t *testing.T, points [][]float64, shapes []shp.Shape) {
	for n, s := range shapes {
		p, ok := s.(*shp.PolyLine)
		if !ok {
			t.Fatal("Failed to type assert.")
		}
		for k, point := range p.Points {
			assert.Equal(t, points[n*3+k], []float64{point.X, point.Y})
		}
	}
}

func testPolygon(t *testing.T, points [][]float64, shapes []shp.Shape) {
	for n, s := range shapes {
		p, ok := s.(*shp.Polygon)
		if !ok {
			t.Fatal("Failed to type assert.")
		}
		for k, point := range p.Points {
			assert.Equal(t, points[n*3+k], []float64{point.X, point.Y})
		}
	}
}

func TestSHPReadZip(t *testing.T) {
	testshapeIdentity(t, "shapetest/shapes.zip", getShapesFromFile)
}

func TestSHPReadPoint(t *testing.T) {
	testshapeIdentity(t, "shapetest/point.shp", getShapesFromFile)
}

func TestSHPReadPolyLine(t *testing.T) {
	testshapeIdentity(t, "shapetest/polyline.shp", getShapesFromFile)
}

func TestSHPReadPolygon(t *testing.T) {
	testshapeIdentity(t, "shapetest/polygon.shp", getShapesFromFile)
}

func testshapeIdentity(t *testing.T, prefix string, getter shapeGetterFunc) {
	shapes := getter(prefix, t)
	d := dataForReadTests[prefix]
	d.tester(t, d.points, shapes)
}

func getShapesFromFile(filename string, t *testing.T) (shapes []shp.Shape) {
	var reader ShapeReader
	var err error
	osr, err := os.Open(filename)
	assert.NoError(t, err)
	if strings.HasSuffix(filename, ".shp") {
		reader, err = shp.ReadFrom(osr)
	} else {
		reader, err = shp.ReadZipFrom(osr)
	}
	if err != nil {
		t.Fatal("Failed to open shapefile: " + filename + " (" + err.Error() + ")")
	}

	for reader.Next() {
		_, shape := reader.Shape()
		shapes = append(shapes, shape)
	}
	if reader.Err() != nil {
		t.Errorf("error while getting shapes for %s: %v", filename, reader.Err())
	}

	return shapes
}
