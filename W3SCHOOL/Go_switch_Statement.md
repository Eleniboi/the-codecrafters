### Go switch Statement

we use the switch statement to select one of many code blocks to be executed.

e.g
```go
package main

import (
    "fmt"
)

func main() {
    
    man := true 
    
    
    switch man{
        case true:
        fmt.Println("i love you lord")
        case false:
        fmt.Println("i still love you lord")
    }
}
```
