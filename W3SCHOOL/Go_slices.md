## Go Slices
slices are flexible data structure, which hold a values of specific type

a slice can grow and shrink


## In Go, there are several ways to create a slice:

example

```go 


package main

import "fmt"


func main() {
// Using the []datatype{values} format

sam := []string{"samuel","you","are","great"}

// Create a slice from an array

arr1 :=[5]string{"God","is","good","to","me"}
slice :=arr1[0:3]

//    Using the make() function
sammy := make([]int, 5)

fmt.Println(sam)
fmt.Println(arr1)
fmt.Println(slice)
fmt.Println(sammy)
}

```
- In Go, there are two functions that can be used to return the length and capacity of a slice:

len() function - returns the length of the slice (the number of elements in the slice)
cap() function - returns the capacity of the slice (the number of elements the slice can grow or shrink to)

example 

```go
package main
import ("fmt")

func main() {
  myslice1 := []int{}
  fmt.Println(len(myslice1))
  fmt.Println(cap(myslice1))
  fmt.Println(myslice1)

  myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
  fmt.Println(len(myslice2))
  fmt.Println(cap(myslice2))
  fmt.Println(myslice2)
}
```


### Go Access, Change, Append and Copy Slices

You can access a specific slice element by referring to the index number.

## Append Elements To a Slice

append element to the end of a slice 

example 
```go
package main


import ("fmt")

func main() {
  myslice1 := []int{1, 2, 3, 4, 5, 6}
  fmt.Printf("myslice1 = %v\n", myslice1)
  fmt.Printf("length = %d\n", len(myslice1))
  fmt.Printf("capacity = %d\n", cap(myslice1))

  myslice2 := append(myslice1, 20, 21)
  fmt.Printf("myslice1 = %v\n", myslice1)
  fmt.Printf("length = %d\n", len(myslice1))
  fmt.Printf("capacity = %d\n", cap(myslice1))


  myslice3 := append(myslice1, myslice2...)

  fmt.Printf("length myslice3 = %d\n", len(myslice3))
  fmt.Printf("capacity myslice3 = %d\n", cap(myslice3))
  fmt.Printf("myslice3 = %v\n", myslice3)

}
```

### Memory Efficiency

When using slices, Go loads all the underlying elements into the memory.

If the array is large and you need only a few elements, it is better to copy those elements using the copy() function.

The copy() function creates a new underlying array with only the required elements for the slice. This will reduce the memory used for the program. 

```go
package main 


import (
    "fmt"
)

func main() {
    
    number :=[14]int{1,2,3,4,5,6,7,8,9,10,11,12,13,14}
    
    fmt.Println(number)
    fmt.Println(len(number))
    
    slicenumber := number[1:7]
    
    fmt.Println(slicenumber)
    fmt.Println(len(slicenumber))
     fmt.Println(cap(slicenumber))
    
    copynumber := make([]int, len(slicenumber))
    
    copy(copynumber, slicenumber)
    
    fmt.Println(copynumber)
    fmt.Println(len(copynumber))
    fmt.Println(cap(copynumber))
    
}
```

