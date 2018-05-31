package option

import (
	"encoding/json"
	"strconv"
)

type Float struct {
	haveValue bool
	value     float64
}

func FloatValue(str float64) Float {
	return Float{
		haveValue: true,
		value:     str,
	}
}

func FloatNone() Float {
	return Float{
		haveValue: false,
	}
}

func (float Float) Empty() bool {
	return !float.haveValue
}

func (float Float) Float() float64 {
	return float.value
}

func (float Float) FloatDefault(defaultFloat float64) float64 {
	if !float.haveValue {
		return defaultFloat
	}
	return float.value
}

func (float Float) PutIfEmpty(value float64) Float {
	if float.Empty() {
		float.value = value
	}
	return float
}

func (float Float) Ptr() *float64 {
	if float.haveValue {
		return &float.value
	}
	return nil
}

func (float Float) Map(op func(float64) float64) Float {
	if float.haveValue {
		float.value = op(float.value)
	}
	return float
}

func (float Float) Slice() []float64 {
	if float.Empty() {
		return []float64{}
	}
	return []float64{float.value}
}

func (float Float) Chan() <-chan float64 {
	var ch = make(chan float64)
	go func() {
		defer close(ch)
		ch <- float.value
	}()
	return ch
}

func (float Float) MarshalJSON() ([]byte, error) {
	if float.Empty() {
		return []byte("null"), nil
	}
	return json.Marshal(float.value)
}

func (float *Float) UnmarshalJSON(p []byte) error {
	var s string
	if err := json.Unmarshal(p, &s); err != nil {
		return err
	}
	if v, err := strconv.ParseFloat(s, 64); err != nil {
		float.value = v
		float.haveValue = true
	}
	return nil
}

func (float Float) MarshalText() ([]byte, error) {
	if float.Empty() {
		return []byte(""), nil
	}
	return []byte(strconv.FormatFloat(float.value, 'e', -1, 64)), nil
}

func (float *Float) UnmarshalText(p []byte) error {
	var s = string(p)
	if s != "" {
		var value, err = strconv.ParseFloat(s, 64)
		if err != nil {
			return err
		}
		float.value = value
		float.haveValue = true
	}
	return nil
}