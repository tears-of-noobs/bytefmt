# bytefmt

Golang library to convert human representation values of disk space, RAM size, etc. to bytes (float64) and back.

## How to get it
```bash
go get github.com/tears-of-noobs/bytefmt
```

## How to use it

```go
package main

import (
    "fmt"
    
    "github.com/tears-of-noobs/bytefmt"
)

func main() {
    // Parsing string. 
    // You also may pass argument in different format.
    // ParseString understand - "1,5TB", "1.5TB", "1.5TiB", "123534"
    byteResult, err := bytefmt.ParseString("5.67GiB")
    if err != nil {
        panic(err)
    }
    // Output - 6.08811614208e+09
    fmt.Printf("%v\n", byteResult)
    
    // Formatting bytes
    stringResult := bytefmt.FormatBytes(byteResult, 2, true)
    
    // Output - "5.67GiB"
    // If you pass binary to false you will get - "6.09GB"
    fmt.Printf("%s\n", stringResult)
}
```
