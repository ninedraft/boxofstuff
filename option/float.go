package option

import "encoding/json"

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

func (str Float) Empty() bool {
	return !str.haveValue
}

func (str Float) Float() float64 {
	return str.value
}

func (str Float) FloatDefault(defaultFloat float64) float64 {
	if !str.haveValue {
		return defaultFloat
	}
	return str.value
}

func (str Float) PutIfEmpty(value float64) Float {
	if str.Empty() {
		str.value = value
	}
	return str
}

func (str Float) Ptr() *float64 {
	if str.haveValue {
		return &str.value
	}
	return nil
}

func (str Float) Map(op func(float64) float64) Float {
	if str.haveValue {
		str.value = op(str.value)
	}
	return str
}

func (str Float) Slice() []float64 {
	if str.Empty() {
		return []float64{}
	}
	return []float64{str.value}
}

func (str Float) Chan() <-chan float64 {
	var ch = make(chan float64)
	go func() {
		defer close(ch)
		ch <- str.value
	}()
	return ch
}

func (float Float) MarshalJSON() ([]byte, error) {
	return json.Marshal(float.value)
}

func (float *Float) UnmarshalJSON(p []byte) error {
	var s float64
	if err := json.Unmarshal(p, &s); err != nil {
		return err
	}
	float.value = s
	return nil
}
