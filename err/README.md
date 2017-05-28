## Err - string errors
Read Dave Cheney article about [constant errors](https://dave.cheney.net/2016/04/07/constant-errors).

>As constants of the Error type are not variables, they are immutable.

Usage:
```go
import(
    ."github.com/ninedraft/boxofstuff/err"
)

const (
    ErrInvalidConfigFile Err = "invalid config file"
    ErrInvalidFilePath Err = "invalid file path"
    ErrIvalidEMail Err = "invalid e-mail"
)

func loadConfig(filePath string) error {
    // ...
    if invalidConfig(configText) {
        return ErrInvalidConfigFile
    }
}

func main() {
    switch loadConfig("./config.toml") {
    case nil:
        // yay
    case ErrInvalidConnfig:
        // do stuff
    case ErrInvalidFilePath:
        // do stuff
    default:
        // do stuff
    }
}

```
