package main


import (
    "fmt"
    "tour/example"
)


func main() {
    fmt.Println("main")

    // we can use functions/closures like in js/ts
    // here i use a anonymous function 
    example.Closures(func(name string){fmt.Println(fmt.Sprintf("hello %s",name))})
    // returning a function
    newFunc := example.ReturnunFunc()
    newFunc(10)
    // we can merge all the above examples


    example.ReadBooksConcurrent()

}
