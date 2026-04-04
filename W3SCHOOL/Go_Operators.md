### Go Operators

The "+" operator adds together two values

e.g 
```go
package main
import ("fmt")

func main() {
  var a = 15 + 25
  fmt.Println(a)
}
```

### Go Arithmetic Operators...

|Operator |	Name | 	Description | 	Example|
|---------|------|--------------|----------|
|+ | 	Addition |	Adds together two values |	x + y |
| - |Subtraction |	Subtracts one value from another |	x - y |	
|* |	Multiplication |	Multiplies two values |	x * y |	
|/ |Division |	Divides one value by another |	x / y |	
|% |	Modulus |	Returns the division remainder|x % y |	
|++|Increment |	Increases the value of a variable by 1 |	x++ |	
|-- |	Decrement |	Decreases the value of a variable by 1|x--|


## Assignment Operators

Assignment operators are used to assign values to variables.

e.g 
=
+= 
-=
|=
/= 


## Go Comparison operators
they are use to compare between two or more values

Example
```go
package main
import ("fmt")

func main() {
  var x = 5
  var y = 3
  fmt.Println(x>y) // returns 1 (true) because 5 is greater than 3
}
```

|Operator |	Name 	|Example|
|---------|---------|-------|
|== |Equal to| 	x == y| 	
|!= 	|Not equal |	x != y |	
|> |	Greater than |	x > y| 	
|< |	Less than |	x < y |
|>= |Greater than or equal to |	x >= y| 	
|<= |	Less than or equal to |	x <= y|

### Go Logical Operators

Logical operators are used to determine the logic between variables or values:

e.g 
```go

package main

import (
    "fmt"
)

func main() {
    v := 5
    x := 8
   // p :=9

   var r bool= v < x || p > x


   fmt.Println(r)
}
```
### Bitwise Operators

Bitwise operators are used on (binary) numbers:

