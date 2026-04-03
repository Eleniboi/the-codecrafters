### The Println() Function

The Println() function is similar to Print() with the difference that a whitespace is added between the arguments, and a newline is added at the end:

Example
```go 
package main
import ("fmt")

func main() {
  var i,j string = "Hello","World"

  fmt.Println(i,j)
}
```

