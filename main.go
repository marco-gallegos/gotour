package main

import (
	"fmt"
)

// func main() {
// 	fmt.Println("package main")

// 	// we can use functions/closures like in js/ts
// 	// here i use a anonymous function
// 	example.Closures(func(name string) { fmt.Println(fmt.Sprintf("hello %s", name)) })
// 	// returning a function
// 	newFunc := example.ReturnunFunc()
// 	newFunc(10)
// 	// we can merge all the above examples

// 	example.ReadBooksConcurrent()

// 	fmt.Println("challenges")
// 	challenges.Main()
// }

// channels example -> how to await
func worker(c chan struct{}) {
	// read from channel and printit
	<-c
	fmt.Println("worker done")
	close(c)
}

func main() {
	fmt.Println("Hello, playground")
	c := make(chan struct{})
	go worker(c)

	fmt.Println("main: sending signal to worker")
}

// basic channel example

// func worker(c chan int) {
// 	fmt.Println("Worker is working...")
// 	var cresult int = <-c

// 	fmt.Println("Worker is done. Result:", cresult)
// 	// close(c)
// }

// func main() {
// 	fmt.Println("Hello, World!")
// 	c := make(chan int)

// 	go worker(c)

// 	c <- 1

// 	fmt.Println("Main is waiting for the worker to finish...")
// }
