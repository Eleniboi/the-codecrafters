
## what is Go Syntax
syntax are the suppose way, structure and order we write our go program

EXAMPLE 
```go
package main

import (
    "fmt"
)

func main() {
    fmt.Println("Hello World!")
}
```

- every go program belong to a "package" and here it is the main package 

- import ("fmt") lets us import files included in the fmt package.

- func main() is the entry point of our program it is where our program start running from  "{ }" and the curly bracket is called the function body.
- fmt.Println() is a function from the fmt library which help to output text to the screen.