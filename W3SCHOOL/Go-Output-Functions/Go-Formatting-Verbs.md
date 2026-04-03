## Go Formatting Verbs

Go Formatting Verbs: are place holder within string they determine how a value should be displayed

Example of Formatting Verbs

The following verbs can be used with all data types:

Verb 	Description

%v    ---> 	 Prints the value in the default format
%#v   --->	 Prints the value in Go-syntax format
%T 	  --->   Prints the type of the value
%% 	  --->   Prints the % sign

example

```go
package main 

import (
    "fmt"
)
func main(){
    var num := 15.5

    var hex := "255"

    //var str := "samuel"

    fmt.Printf("Amount %.2f%% \n was around up to %.f%%\n ", num, num)
    fmt.Printf("hex form %X \n then the value %v\n and the type is %T\n", hex,hex, hex)
    fmt.Printf("")

}