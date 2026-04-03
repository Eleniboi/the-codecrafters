## Go Constants

- A constants in go is use to declare a fix value 
we use the "const" keyword to declare a constant which means that it is unchangeable and read-only.

syntax 
```go
- const samuel string = "great guy"

- const num = 50 // it mean the value of num can not be change anywhere in the program


package main
import ("fmt")

const PI = 3.14

func main() {
  fmt.Println(PI)
}

```

## Constant Rules

    Constant names follow the same naming rules as variables
    Constant names are usually written in uppercase letters (for easy identification and differentiation from variables)
    Constants can be declared both inside and outside of a function

## Constant Types

There are two types of constants:

- Typed constants := their type are explicitly defined
- Untyped constants := their type are inferred from the value 

## multiple const 

multiple constants either of same or different size can be group into one block of readability
