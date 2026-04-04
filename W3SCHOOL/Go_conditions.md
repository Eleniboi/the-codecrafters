### Go Conditions

conditional statements are Go way of making decisions

A condition can be either true or false.

- Go supports the usual comparison operators from mathematics:

    Less than <
    Less than or equal <=
    Greater than >
    Greater than or equal >=
    Equal to ==
    Not equal to !=

- Additionally, Go supports the usual logical operators:

    Logical AND &&
    Logical OR ||
    Logical NOT !

### Go if statement

if statement is use to execute code if the actual condisions are true


e.g 
```go
package main
import ("fmt")

func main() {
  if 20 > 18 {
    fmt.Println("20 is greater than 18")
  }
}
```

### Go else if statement
is like a backup plan for the if statement, it only run when the if statement is not meet

e.g
```go
package main
import ("fmt")

func main() {
  if 20 > 18 {
    fmt.Println("20 is greater than 18")
  }else if 18 < 20 {
    fmt.Println("yes 18 is less than 20")
  }
}
```


### The Nested if Statement

You can have if statements inside if statements, this is called a nested if.

e.g
```go
package main
import ("fmt")

func main() {
  num := 20
  if num >= 10 {
    fmt.Println("Num is more than 10.")
    if num > 15 {
      fmt.Println("Num is also more than 15.")
     }
  } else {
    fmt.Println("Num is less than 10.")
  }
}
```