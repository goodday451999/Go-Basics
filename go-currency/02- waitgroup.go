package main

import "fmt"

func myFunc() {
    fmt.Println("Inside my goroutine")
}

func main() {
    fmt.Println("Hello World")
    go myFunc()
    fmt.Println("Finished Execution")
}

//////////////////////////////////////////////////////

package main

import (
    "fmt"
    "sync"
)

func myFunc(waitgroup *sync.WaitGroup) {
    fmt.Println("Inside my goroutine")
    waitgroup.Done()
}

func main() {
    fmt.Println("Hello World")

    var waitgroup sync.WaitGroup
    waitgroup.Add(1)
    go myFunc(&waitgroup)
    waitgroup.Wait()

    fmt.Println("Finished Execution")
}
