package str

import "testing"

func TestVector_Inverse(t *testing.T) {
	var x = Vector{"a", "b", "c"}
	var y = Vector{"c", "b", "a"}
	var invertedX = x.Inverse()
	if !invertedX.Eq(y) {
		t.Fatalf("%v and %v are not equal", invertedX, y)
	}
}
