package err

import (
	"fmt"
)

// Err - string err
// Usage:
/*
```
import(
	. "github.com/ninedraft/boxofstuff"
)

const(
	ErrInvalidConfigFile Err = "invalid config file"
)
```
*/
type Err string

// Errorf - creates Err from formatted string
func Errorf(format string, args ...interface{}) Err {
	return Err(fmt.Sprintf(format, args...))
}

func (err Err) String() string {
	return string(err)
}

func (err Err) Error() string {
	return string(err)
}
