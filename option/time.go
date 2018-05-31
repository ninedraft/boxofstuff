package option

import (
	"encoding/json"
	"time"
)

type Time struct {
	haveValue bool
	value     time.Time
}

func TimeValue(t time.Time) Time {
	return Time{
		haveValue: true,
		value:     t,
	}
}

func TimeNone() Time {
	return Time{
		haveValue: false,
	}
}

func (t Time) Empty() bool {
	return !t.haveValue
}

func (t Time) Time() time.Time {
	return t.value
}

func (t Time) TimeDefault(defaultTime time.Time) time.Time {
	if !t.haveValue {
		return defaultTime
	}
	return t.value
}

func (t Time) PutIfEmpty(value time.Time) Time {
	if t.Empty() {
		t.value = value
	}
	return t
}

func (t Time) Ptr() *time.Time {
	if t.haveValue {
		return &t.value
	}
	return nil
}

func (t Time) Map(op func(time.Time) time.Time) Time {
	if t.haveValue {
		t.value = op(t.value)
	}
	return t
}

func (t Time) Slice() []time.Time {
	if t.Empty() {
		return []time.Time{}
	}
	return []time.Time{t.value}
}

func (t Time) Chan() <-chan time.Time {
	var ch = make(chan time.Time)
	go func() {
		defer close(ch)
		ch <- t.value
	}()
	return ch
}

func (t Time) MarshalJSON() ([]byte, error) {
	if t.Empty() {
		return []byte("null"), nil
	}
	return json.Marshal(t.value)
}

func (t *Time) UnmarshalJSON(p []byte) error {
	var value time.Time
	if string(p) == "null" {
		return nil
	}
	if err := json.Unmarshal(p, &value); err != nil {
		return err
	}
	t.value = value
	t.haveValue = true
	return nil
}

func (t Time) MarshalText() ([]byte, error) {
	if t.Empty() {
		return []byte(""), nil
	}
	return t.value.MarshalText()
}

func (t *Time) UnmarshalText(p []byte) error {
	var s = string(p)
	if s != "" {
		var value time.Time
		if err := value.UnmarshalText(p); err != nil {
			return err
		}
		t.value = value
		t.haveValue = true
	}
	return nil
}

func (t Time) String() string {
	if t.Empty() {
		return ""
	}
	return t.value.String()
}
