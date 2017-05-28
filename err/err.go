package main

import (
	"fmt"
)

type Err string

func Errorf(format string, args ...interface{}) Err {
	return Err(fmt.Sprintf(format, args))
}

func (err Err) String() string {
	return string(err)
}

func (err Err) Error() string {
	return string(err)
}
