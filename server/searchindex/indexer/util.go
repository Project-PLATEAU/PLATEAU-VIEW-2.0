package indexer

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"time"

	"github.com/qmuntal/gltf"
	b3dms "github.com/reearth/go3dtiles/b3dm"
	"github.com/reearth/reearthx/log"
	"gonum.org/v1/gonum/mat"
)

const (
	AxisX = 0
	AxisY = 1
	AxisZ = 2
)

type CesiumRTC struct {
	Center [3]float64
}

func isRtcCenterEmpty(arr [3]float64) bool {
	for _, val := range arr {
		if val != 0 {
			return false
		}
	}
	return true
}

func getRtcTransform(ft *b3dms.B3dmFeatureTable, gltf *gltf.Document) (*mat.Dense, error) {
	rtcCenter := ft.RtcCenter
	if isRtcCenterEmpty(rtcCenter) {
		var temp CesiumRTC
		if err := json.Unmarshal(gltf.Extensions["CESIUM_RTC"].(json.RawMessage), &temp); err != nil {
			return nil, fmt.Errorf("unmarshal failed for cesium_rtc: %v", err)
		}
		rtcCenter = temp.Center
	}
	rtcTransform := eyeMat(4)
	if len(rtcCenter) > 0 {
		rtcTransform = mat4FromCartesian(cartesianFromSlice(rtcCenter[:]))
	}
	return rtcTransform, nil
}

// Creates a rotation matrix around the x-axis.
func getYUpToZUp() *mat.Dense {
	sinAngle, cosAngle := math.Sincos(math.Pi / 2)
	d := []float64{
		1.0,
		0.0,
		0.0,
		0.0,
		0.0,
		cosAngle,
		sinAngle,
		0.0,
		0.0,
		-sinAngle,
		cosAngle,
		0.0,
		0.0,
		0.0,
		0.0,
		1.0,
	}

	return mat.NewDense(4, 4, d)
}

func getZUpTransform() *mat.Dense {
	// discuss if we need gltfAxisUpAxis
	upAxis := AxisY
	transform := eyeMat(4)
	if upAxis == AxisY {
		transform = getYUpToZUp()
	}
	return transform
}

var floatType = reflect.TypeOf(float64(0))

func getFloat(value float32) (float64, error) {
	v := reflect.ValueOf(value)
	v = reflect.Indirect(v)
	if !v.Type().ConvertibleTo(floatType) {
		return 0, fmt.Errorf("cannot convert %v to float64", v.Type())
	}
	fv := v.Convert(floatType)
	return fv.Float(), nil
}

var intType = reflect.TypeOf(int(0))

func getInt(value interface{}) (int64, error) {
	v := reflect.ValueOf(value)
	v = reflect.Indirect(v)
	if !v.Type().ConvertibleTo(intType) {
		return 0, fmt.Errorf("cannot convert %v to int", v.Type())
	}
	iv := v.Convert(intType)
	return iv.Int(), nil
}

func Map(elem []interface{}, f func(interface{}) (float64, error)) ([]float64, error) {
	result := make([]float64, len(elem))
	var err error
	for i, v := range elem {
		result[i], err = f(v)
		if err != nil {
			return nil, fmt.Errorf("failed to apply function the function: %w", err)
		}
	}
	return result, err
}

func minMaxOfSlice(array []float64) (float64, float64) {
	var max float64 = array[0]
	var min float64 = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

type RetriableFunc func() error

func Retry(RetriableFunc RetriableFunc) error {
	var n uint
	config := newDefaultRetryConfig()
	if err := config.context.Err(); err != nil {
		return err
	}

	shouldRetry := true
	for shouldRetry {
		err := RetriableFunc()
		if err != nil {
			log.Errorf("retry error (%d): %v", n, err)

			// if this is last attempt - don't wait
			if n == config.attempts-1 {
				return fmt.Errorf("retry failed after %d attempts: %v", n, err)
			}

			select {
			case <-time.After(delay(config, n, err)):
			case <-config.context.Done():
				return config.context.Err()
			}

		} else {
			return nil
		}

		n++
		shouldRetry = shouldRetry && n < config.attempts
	}

	return nil
}

type RetryConfig struct {
	attempts    uint
	delay       time.Duration
	maxDelay    time.Duration
	maxJitter   time.Duration
	delayType   DelayTypeFunc
	context     context.Context
	maxBackOffN uint
}

type DelayTypeFunc func(n uint, err error, config *RetryConfig) time.Duration

func newDefaultRetryConfig() *RetryConfig {
	return &RetryConfig{
		attempts:  uint(5),
		delay:     100 * time.Millisecond,
		maxJitter: 100 * time.Millisecond,
		delayType: combineDelay(backOffDelay, randomDelay),
		context:   context.Background(),
	}
}

func combineDelay(delays ...DelayTypeFunc) DelayTypeFunc {
	const maxInt64 = uint64(math.MaxInt64)
	return func(n uint, err error, config *RetryConfig) time.Duration {
		var total uint64
		for _, delay := range delays {
			total += uint64(delay(n, err, config))
			if total > maxInt64 {
				total = maxInt64
			}
		}
		return time.Duration(total)
	}
}

func delay(config *RetryConfig, n uint, err error) time.Duration {
	delayTime := config.delayType(n, err, config)
	if config.maxDelay > 0 && delayTime > config.maxDelay {
		delayTime = config.maxDelay
	}
	return delayTime
}

func backOffDelay(n uint, _ error, config *RetryConfig) time.Duration {
	// 1 << 63 would overflow signed int64 (time.Duration), thus 62.
	const max uint = 62
	if config.delay <= 0 {
		config.delay = 1
	}
	config.maxBackOffN = max - uint(math.Floor(math.Log2(float64(config.delay))))

	return config.delay << n
}

func randomDelay(_ uint, _ error, config *RetryConfig) time.Duration {
	return time.Duration(rand.Int63n(int64(config.maxJitter)))
}
