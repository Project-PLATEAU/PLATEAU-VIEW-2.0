package indexer

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCartogrphicFromCartesian(t *testing.T) {
	tests := []struct {
		name          string
		input         *Cartesian3
		want          *Cartographic
		expectedError error
		wantErr       bool
	}{
		{"openError", nil, nil, errors.New("failed to do scaleToGeoticSurface transformation: cartesian is required"), true},
		{"RealInput", &Cartesian3{
			X: -3962987.9740489936,
			Y: 3361559.5611439403,
			Z: 3685745.6979325633,
		}, &Cartographic{
			Longitude: 2.4381220341286984,
			Latitude:  0.6200812555240751,
			Height:    78.32792784513339,
		}, nil, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := cartographicFromCartesian3(tt.input)
			if tt.wantErr {
				if assert.Error(t, err) {
					assert.Equal(t, err, tt.expectedError, "Expected an error")
				}
			} else {
				if assert.NoError(t, err) {
					assert.Equal(t, got, tt.want, "scaleToGeoDeticSurface() = false")
				}
			}
		})
	}
}
