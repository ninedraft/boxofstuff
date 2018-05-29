package tribool

import (
	"encoding/json"
	"fmt"
)

type Tribool interface {
	fmt.Stringer
	json.Marshaler
	json.Unmarshaler
	IsTrue() bool
	IsFalse() bool
	IsInderminate() bool

	boop()
}
