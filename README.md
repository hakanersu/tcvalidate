So its validates Turkish Identification Number.

There is numbers of packages that do same thing and its my way to do it.

```go
package main

import (
    "fmt"
    "github.com/hakanersu/tcvalidate"
)

func main() {

    sonuc := validatetc.Validate("17223038680")

    fmt.Println(sonuc) // true

}

```

