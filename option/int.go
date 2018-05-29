package option

import "encoding/json"

type Int struct {
	haveValue bool
	value     int
}

func IntValue(i int) Int {
	return Int{
		haveValue: true,
		value:     i,
	}
}

func IntNone() Int {
	return Int{
		haveValue: false,
	}
}

func (i Int) Empty() bool {
	return !i.haveValue
}

func (i Int) Int() int {
	return i.value
}

func (i Int) IntDefault(defaultInt int) int {
	if !i.haveValue {
		return defaultInt
	}
	return i.value
}

func (i Int) PutIfEmpty(value int) Int {
	if i.Empty() {
		i.value = value
	}
	return i
}

func (i Int) Ptr() *int {
	if i.haveValue {
		return &i.value
	}
	return nil
}

func (i Int) Map(op func(int) int) Int {
	if i.haveValue {
		i.value = op(i.value)
	}
	return i
}

func (i Int) Slice() []int {
	if i.Empty() {
		return []int{}
	}
	return []int{i.value}
}

func (i Int) Chan() <-chan int {
	var ch = make(chan int)
	go func() {
		defer close(ch)
		ch <- i.value
	}()
	return ch
}

func (i Int) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.value)
}

func (i *Int) UnmarshalJSON(p []byte) error {
	var s int
	if err := json.Unmarshal(p, &s); err != nil {
		return err
	}
	i.value = s
	return nil
}
