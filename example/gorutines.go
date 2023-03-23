package example
/*
# 2023-03-21

old behavior works just use go and a function.

new concepts
- waitgroup
- mutex


## wait group

### why?

this solves the problem about waiting gorutines execution (eje without this if you run a function n times this probably will miss some executions if is a simple main call)


sync.WaitGroup
- Add
- Done
- Wait



## mutex

### why?

this solves the problem accesing the same memory direction (read or write) at the same time.

sync.Mutex | sync.RWMutex

- Lock()
- Unlock()
- RLock()
- RUnlock()

### notes

wait groups is like sync await in js and mutex work like accessing same variables in micropython on every core (rp2040 on rpi pico)


## TODO

- [ ] read more about this

*/

import (
    "fmt"
    "sync"
)

type Book struct {
    Id int
    Name string
    Readed bool
}

// in memory db 
var BooksDb []Book = []Book{
    {1, "vaquero", false},
    {2, "condorito v1", false},
    {3, "condorito v2", false},
    {4, "el principito", false},
    {5, "sunny sunshine canticos", false},
    {6, "calculo diferencial", false},
    {7, "algebra lineal", false},
    {8, "la biblia de mysql", false},
}


// simply find
func findBook(id int, wg *sync.WaitGroup, mtx *sync.Mutex) (int, Book){
    var index int = -1
    var book Book

    mtx.Lock()
    for i,b := range BooksDb {
        if b.Id == id {
            index = i
            book = b
        }
    }
    mtx.Unlock()

    return index, book
}

func readBook(id int, wg *sync.WaitGroup, mtx *sync.Mutex) bool {
    index, book := findBook(id, wg, mtx)
    if index < 0 {
        return false
    }
    
    mtx.Lock()
    book.Readed = true
    BooksDb[index] = book
    mtx.Unlock()

    fmt.Printf("readed book %s \n", book.Name)

    wg.Done()

    return true
}


// example function
// the idea here is use a in memory var (like a db) and try to acces and modify this 
// using gorutines in a way to afront concurrency troubles 
func ReadBooksConcurrent(){
    fmt.Println("reading books")
    // memory access limitation
    var mutex = &sync.Mutex{}
    
    // concurrency handler (wait)
    var waitgroup = &sync.WaitGroup{}
    
    var booksToRead []int = []int{
        4,
        8,
        7,
        2,
    }

    for _, number := range booksToRead {
        // need add a wait group element on every go call
        waitgroup.Add(1)

        // same old concurrency call
        go readBook(number, waitgroup, mutex)
    }

    // wait for all Add() to be Done()
    waitgroup.Wait()
}


