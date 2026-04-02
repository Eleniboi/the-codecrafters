## Go Multiple Variable Declaration

In Go, it is possible to declare multiple variables on the same line.

example 

```go

package main

import (
    "fmt"
)

func main() {

    var brother, sister, mom string = "samuel", "enenu", "mummy" // here the type is declared 

    var a, b, c, d int = 1, 3, 5, 7 // here the type is declared 


    var name, age = "samue", 16

    var (
        small string 
        b int = 1
        c string = "Joseph"
    )
}
```

