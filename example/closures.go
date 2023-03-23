package example

import (
    "fmt"
)

// example using a function as a arg
func Closures(localFn func(string)){
    fmt.Println("closures example")
    localFn("marco gallegos")
}


func ReturnunFunc() func(int) {
    return func(repeats int){
        for x := range make([]int, repeats){
            fmt.Println("marco", x)
        }
    }
}
