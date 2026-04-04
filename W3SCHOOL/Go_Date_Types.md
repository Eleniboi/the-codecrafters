## Go Data Type

data type specifies the size of variable value

## Go basical have three data types
1. string: they represent string value which are text.
2. Numberic: represent numbers which can be of type float, int, and complex

3. bool: which hold the value of true or false


example
```go
package main 

import (
    "fmt"
)

func main() {

    var (
        a = "brother"
        b = true
        c = 5.5
        d = 88
    )

    fmt.Printf("this is a %T ", a, a)
    fmt.Printf("this is a %T ", b, b)
    fmt.Printf("this is a %T ", c, c)
    fmt.Printf("this is an %T ", d, d)
}
```
### Go Boolean Data Type

A boolean only hold the value of true or false, and it default value is false.

example

```go 
package main
import ("fmt")
 
func main() {
  var b1 bool 

  fmt.Println(b1)
  
}
```
### Go Integer Data Types


integer holds whole numbers e.g 3  4 4 432 -80 -3
and NOTE: the default value of an integer is 0

The integer data type has two categories:

- Signed integers - can store both positive and negative values
- Unsigned integers - can only store non-negative values

example
```go 
package main
import ("fmt")

func main() {
  var x int = 500
  var y int = -4500

  var a unit = 890
  var b unit = 54 
  fmt.Printf("Type: %T, value: %v", x, x)
  fmt.Printf("Type: %T, value: %v", y, y)
  fmt.Printf("Type: %T, value: %v", a, a)
  fmt.Printf("Type: %T, value: %v", b, b)
}
```

### Go Float Data Types
Float holds decimal number like 2.3, 1.5..
they can also store negetive decimal numbers.

The default type for float is float64. If you do not specify a type, the type will be float64.


example
```go
package main

import (
    "fmt"
)

func main() {

    sam := 7.7
    omafu := 8.9

    fmt.Println(sam, " ", omafu)
}

```

### Go String Data Type 

Go string stores text with double quote and anything in that quote becomes a string
NOTE: string are sequence of bytes

Example
```go
package main
import ("fmt")

func main() {
  var txt1 string = "Hello!"
  var txt2 string
  txt3 := "World 1"

  fmt.Printf("Type: %T, value: %v\n", txt1, txt1)
  fmt.Printf("Type: %T, value: %v\n", txt2, txt2)
  fmt.Printf("Type: %T, value: %v\n", txt3, txt3)
}
```
