## Go Functions
A function in go is a reusable block of code, that performs a specific task

we use the func keyword to creat a function

with a function name it is advisible to write a function name that relate with what the code is doing

## Call a Function
Functions are not executed immediately. They are "saved for later use", and will be executed when they are called

example 
```go
package main 

import (
    "fmt"
    )
    
    func mmss() {
        fmt.Println("I Love You Jesus")
    }
    func main() {
        mmss()
    }
 ```

## Go Function Returns

If you want the function to return a value, you need to define the data type of the return value (such as int, string, etc), and also use the return keyword inside the function:

```go
package main 


import (
    "fmt"
    )
    
func adds(a, b int) int{
    return a + b
}

func main() {
    a := 65
    b := 4
    fmt.Println(adds(a,b))
}
```
### Recursion Functions

recursion functions are functions that call thier self