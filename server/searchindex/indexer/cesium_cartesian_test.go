package indexer

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScaleGeodeticSurface(t *testing.T) {
	tests := []struct {
		name          string
		input         *Cartesian3
		want          *Cartesian3
		expectedError error
		wantErr       bool
	}{
		{"openError", nil, nil, errors.New("cartesian is required"), true},
		{"RealInput", &Cartesian3{
			-3.9624720294689154e+06,
			3.3609612889271434e+06,
			3.686769915095649e+06,
		}, &Cartesian3{
			-3.9624481582785267e+06,
			3.3609410414285567e+06,
			3.6867475551372315e+06,
		}, nil, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oneOverRadii := Wgs84OneOverRadii
			oneOverRadiiSquared := Wgs84OneOverRadiiSquared
			centerToleranceSquared := Wgs84CenterToleranceSquared
			got, err := scaleToGeodeticSurface(tt.input, &oneOverRadii, &oneOverRadiiSquared, centerToleranceSquared)
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
