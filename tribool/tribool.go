package tribool

import (
	"encoding/json"
	"errors"
	"fmt"
)

type triboolImpl byte

var (
	_ Tribool = new(triboolImpl)
)

var ErrInvalidJSONRepresentation = errors.New("invalid JSON representation")

const (
	False triboolImpl = iota
	True
	Inderminate
)

func (tribool triboolImpl) boop() {}

func (tribool triboolImpl) String() string {
	switch tribool {
	case True:
		return "true"
	case False:
		return "false"
	case Inderminate:
		return "inderminate"
	default:
		panic(fmt.Errorf("undefined tribool value %d", tribool))
	}
}

func (tribool triboolImpl) Ptr() *triboolImpl {
	return &tribool
}

func (tribool triboolImpl) IsTrue() bool {
	return tribool == True
}

func (tribool triboolImpl) IsFalse() bool {
	return tribool == False
}

func (tribool triboolImpl) IsInderminate() bool {
	return tribool == Inderminate
}

type _TriboolJSON struct {
	TribolValue string `json:"tribool_value"`
}

func (tribool triboolImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(_TriboolJSON{tribool.String()})
}

func (tribool *triboolImpl) UnmarshalJSON(data []byte) error {
	var jsonTribool _TriboolJSON
	if err := json.Unmarshal(data, &jsonTribool); err != nil {
		return err
	}
	switch jsonTribool.TribolValue {
	case "true":
		*tribool = True
	case "false":
		*tribool = False
	case "inderminate":
		*tribool = Inderminate
	default:
		return ErrInvalidJSONRepresentation
	}
	return nil
}
