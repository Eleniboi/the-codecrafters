### Go Arrays

Arrays are fixed data structure  used to store multiple values of the same type in a single variable, instead of declaring separate variables for each value.

### how to declare an array in golang


```go
// Example 

package main

import (
    "fmt"
)

func main() {

//USING THE var KEYWORD
var arr1 = [4]int{5,2 ,4 ,4,}// here the length is define
var arr2 = []int{8,3,3}// here the length is inferred


//USING THE := 
array_name := [length]datatype{values} // here length is defined

array_name := [...]datatype{values} // here length is inferred 

}
```

### Access Elements of an Array

You can access a specific array element by referring to the index number. in golang index start at 0
example

```go

package main
import ("fmt")

func main() {
  prices := [3]int{10,20,30}

  fmt.Println(prices[0])
  fmt.Println(prices[2])
}
```
### Change Elements of an Array
You can also change the value of a specific array element by referring to the index number.

```go

package main

import (
    "fmt"
)

func main() {

    arr1 := [3]int{8,3,1}
    arr1[2] = 7

    fmt.Println(arr1)
}
```

### Array Initialization

If an array or one of its elements has not been initialized in the code, it is assigned the default value of its type.

### Find the Length of an Array
we use the len() function to find the length of an array 
length of an array is basically determine by the values in the array

Example
```go

package main
import ("fmt")

func main() {
  arr1 := [4]string{"Volvo", "BMW", "Ford", "Mazda"}
  arr2 := [...]int{1,2,3,4,5,6}

  fmt.Println(len(arr1))
  fmt.Println(len(arr2))
}
```
