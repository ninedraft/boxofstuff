package option

import "encoding/json"

type String struct {
	haveValue bool
	value     string
}

func StrValue(str string) String {
	return String{
		haveValue: true,
		value:     str,
	}
}

func StrNone() String {
	return String{
		haveValue: false,
	}
}

func (str String) Empty() bool {
	return !str.haveValue
}

func (str String) String() string {
	return str.value
}

func (str String) StringDefault(defaultString string) string {
	if !str.haveValue {
		return defaultString
	}
	return str.value
}

func (str String) PutIfEmpty(value string) String {
	if str.Empty() {
		str.value = value
	}
	return str
}

func (str String) Ptr() *string {
	if str.haveValue {
		return &str.value
	}
	return nil
}

func (str String) Map(op func(string) string) String {
	if str.haveValue {
		str.value = op(str.value)
	}
	return str
}

func (str String) Slice() []string {
	if str.Empty() {
		return []string{}
	}
	return []string{str.value}
}

func (str String) Chan() <-chan string {
	var ch = make(chan string)
	go func() {
		defer close(ch)
		ch <- str.value
	}()
	return ch
}

func (str String) MarshalJSON() ([]byte, error) {
	return json.Marshal(str.value)
}

func (str *String) UnmarshalJSON(p []byte) error {
	var s string
	if err := json.Unmarshal(p, &s); err != nil {
		return err
	}
	if s != "" {
		str.value = s
	}
	return nil
}
