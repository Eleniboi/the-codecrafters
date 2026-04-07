### Go Struct
A struct is a collection of related data to form a single unit or entity, a struct can store data of different type e.g bool, int, float...

A struct can be use to keep records.

### Declare a struct 
You can declare a struct using the type and struct keywords

example 
```go 

type person struct{

    name string 
    age int
    married false
}
```

### Accessing a struct members
To  access a particular struct member we use the dot(.) operator between the struct intance and the structure member

example 
```go

package main 

import (
    "fmt"
)

type person struct{
    Name string
    Address string
    Age int
    job string
}

func main(){
    
    //we use the above struct to keep the following records by accessing each member using the dot(.) operator
    pers1 := person{Name: "Samuel Oloche Omafu", Address: "Califonia", Age: 22, job: "software engineer"}

    fmt.Printf("Name: %q\n",pers1.Name)
    fmt.Printf("Address: %s\n",pers1.Address)
    fmt.Printf("Age: %d\n",pers1.Age)
    fmt.Printf("Job: %s\n\n",pers1.job)
    fmt.Println("----------------------")
    
    
    //This is for another different person using thesame struct but defferent information
    pers2 :=person{Name: "samzuk", Address: "brocklyn", Age: 25, job: "A Tech founder"}
    fmt.Printf("Name: %q\n",pers2.Name)
    fmt.Printf("Address: %s\n",pers2.Address)
    fmt.Printf("Age: %d\n",pers2.Age)
    fmt.Printf("Job: %s\n",pers2.job)
}
```

### Pass Struct as Function Arguments
A struct can also be pass as a function argument, which will be of type the struct name. lets handle some example for clarity

example
```go

package main 

import (
    "fmt"
)

type person struct{
    Name string
    Age int
    Job string
    Married bool
}

func main() {

    pers1 := person{Name: "Mr God'swill", Age: 27, Job: "Code mentor and men builder", Married: false }

    pers2 := person{Name: "Somafu", Age: 22, Job: "Great developer", Married: true}
    
    (IDperson(pers1))
    (IDperson(pers2))
    
}

func IDperson(pers person){
    fmt.Println("Name: ", pers.Name)
    fmt.Println("Age: ", pers.Age)
    fmt.Println("Job: ", pers.Job)
    fmt.Println("Married: ", pers.Married)
}
```