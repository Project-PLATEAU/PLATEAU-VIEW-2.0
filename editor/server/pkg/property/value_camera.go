package property

import (
	"github.com/mitchellh/mapstructure"
)

var ValueTypeCamera = ValueType("camera")

type Camera struct {
	Lat      float64 `json:"lat" mapstructure:"lat"`
	Lng      float64 `json:"lng" mapstructure:"lng"`
	Altitude float64 `json:"altitude" mapstructure:"altitude"`
	Heading  float64 `json:"heading" mapstructure:"heading"`
	Pitch    float64 `json:"pitch" mapstructure:"pitch"`
	Roll     float64 `json:"roll" mapstructure:"roll"`
	FOV      float64 `json:"fov" mapstructure:"fov"`
}

func (c *Camera) Clone() *Camera {
	if c == nil {
		return nil
	}
	return &Camera{
		Lat:      c.Lat,
		Lng:      c.Lng,
		Altitude: c.Altitude,
		Heading:  c.Heading,
		Pitch:    c.Pitch,
		Roll:     c.Roll,
		FOV:      c.FOV,
	}
}

type typePropertyCamera struct{}

func (*typePropertyCamera) I2V(i interface{}) (interface{}, bool) {
	if v, ok := i.(Camera); ok {
		return v, true
	}

	if v, ok := i.(*Camera); ok {
		if v != nil {
			return *v, true
		}
		return nil, false
	}

	v := Camera{}
	if err := mapstructure.Decode(i, &v); err == nil {
		return v, true
	}
	return nil, false
}

func (*typePropertyCamera) V2I(v interface{}) (interface{}, bool) {
	return v, true
}

func (*typePropertyCamera) Validate(i interface{}) bool {
	_, ok := i.(Camera)
	return ok
}

func (v *Value) ValueCamera() (vv Camera, ok bool) {
	if v == nil {
		return
	}
	vv, ok = v.Value().(Camera)
	return
}
