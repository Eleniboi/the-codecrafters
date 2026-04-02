## Go Comments //, /**/

go comments are use to discribe a function, a particular line of code or used to igore a particular feature in go code, they make code more readable. 

### Types Of Go Comments 

1. single line comment: Single-line comments start with two forward slashes (//).

EXAMPLE: 


```GO
// This is a comment
package main
import ("fmt")

func main() {
  // This is a comment
  fmt.Println("Hello World!")
}
```

2. Multiple line comment: the can contain more than one line, so long as the content is between this /* and */ 


EXAMPLE: 

```go
package main
import ("fmt")

func main() {
  /* The code below will print Hello World
  to the screen, and it is amazing */
  fmt.Println("Hello World!")
}
```