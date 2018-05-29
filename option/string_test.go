package option

import (
	"strings"
	"testing"
	"time"
	"encoding/json"
)

func TestStrNone(test *testing.T) {
	if StrNone() != (String{}) {
		test.Errorf("StrNone is not equal to unitialized String option: %#v", StrNone())
	}
}

func TestStrValue(test *testing.T) {
	if StrValue("boop") != (String{haveValue: true, value: "boop"}) {
		test.Errorf("StrValue(\"boop\") is not equal to valued String option: %#v", StrValue("boop"))
	}
}

func TestString_Empty(t *testing.T) {
	if !StrNone().Empty() {
		t.Errorf("method Empty of unitialized String option: expected true, got false")
	}
	if StrValue("boop").Empty() {
		t.Errorf("method Empty of initialized String option: expected false, got true")
	}
}

func TestString_Ptr(t *testing.T) {
	var empty = StrNone()
	var value = StrValue("boop")
	if empty.Ptr() != nil {
		t.Errorf("empty.Ptr(): expected nil, got %#v", empty.Ptr())
	}
	if value.Ptr() == nil {
		t.Errorf("value.Ptr(): expected non-nil, got nil")
	}
	var strPtr = value.Ptr()
	*strPtr = "fooba"
	if value.String() == "fooba" {
		t.Errorf("value.Ptr() allows to mutate inner value")
	}
}

func TestString_Map(t *testing.T) {
	var value = StrValue(" boop_ ").
		Map(strings.TrimSpace).
		Map(func(s string) string { return strings.TrimSuffix(s, "_") }).
		String()
	if value != "boop" {
		t.Fatalf("expected boop, got %q", value)
	}
}

func TestString_Chan(t *testing.T) {
	select {
	case v := <-StrValue("boop").Chan():
		if v != "boop" {
			t.Errorf(`StrValue("boop").Chan(): expected "boop", got %v`, v)
		}
	case <-time.Tick(100 * time.Millisecond):
		t.Errorf(`StrValue("boop").Chan(): timeout'`)
	}
}

func TestString_Slice(t *testing.T) {
	var slice = StrValue("boop").Slice()
	if len(slice) != 1 {
		t.Errorf(`StrValue("boop").Chan(): expected len=1, got %d`, len(slice))
	}
	for _, v := range slice {
		if v != "boop" {
			t.Errorf(`StrValue("boop").Slice(): expected "boop", got %v`, v)
		}
	}
}

func TestString_MarshalJSON(t *testing.T) {
	type user struct {
		Username String
	}
	var in = user{
		Username: StrValue("merlin"),
	}
	data, err := json.Marshal(in)
	t.Log(string(data), err)

	var out user
	json.Unmarshal(data, &out)
	t.Log(out.Username.Empty())
}
