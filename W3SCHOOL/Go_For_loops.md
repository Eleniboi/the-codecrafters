## Go For Loops

The for loop loops through a block of code a specified number of times.

The for loop is the only loop available in Go.

A loop keep iterating until a particular block of code is met

iteration is the process of repeating a block of code.

it is triger by the for keyword

example 
```go
package main
import ("fmt")

func main() {
  for i:=0; i < 5; i++ {
    fmt.Println(i)
  }
}
```
## The continue Statement

The continue statement is used to skip one or more iterations in the loop. It then continues with the next iteration in the loop.

## The break Statement

The break statement is used to break/terminate the loop execution.
## Nested Loops

is a loop place in another loop

example
```go
package main
import ("fmt")

func main() {
  adj := [2]string{"big", "tasty"}
  fruits := [3]string{"apple", "orange", "banana"}
  for i:=0; i < len(adj); i++ {
    for j:=0; j < len(fruits); j++ {
      fmt.Println(adj[i],fruits[j])
    }
  }
}
```