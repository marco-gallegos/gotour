package example


/*
# 2023-03-20

better support on closures (and anonymous functions), looks inspired by js.

is important define return and input types
return functions
declare inline
use current functions
use as parameter

## todo

- [ ] apply a function passed as arg to another input
- [ ] how flexible is now with generics

*/

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
