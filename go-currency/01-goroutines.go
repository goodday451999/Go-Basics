package main


import (
    "fmt"
    "time"
)

// notice we've not changed anything in this function
// when compared to our previous sequential program
func compute(value int) {
    for i := 0; i < value; i++ {
        time.Sleep(time.Second)
        fmt.Println(i)
    }
}

func main() {
    fmt.Println("Goroutine")

    // sequential execution of our compute function
    compute(5)
    compute(5)

    // notice how we've added the 'go' keyword
    // in front of both our compute function calls
    go compute(5)
    go compute(5)

	
    //var input string
    fmt.Scanln()
}
