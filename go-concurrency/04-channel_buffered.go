package main

import (
    "fmt"
    "math/rand"
    "time"
)

func CalculateValue(c chan int) {
    value := rand.Intn(10)
    fmt.Println("Calculated Random Value: {}", value)
    time.Sleep(1000 * time.Millisecond)
    c <- value
    fmt.Println("This executes regardless as the send is now non-blocking")
}

func main() {
    fmt.Println("Channel Buffered")

    valueChannel := make(chan int, 2)
    defer close(valueChannel)

    go CalculateValue(valueChannel)
    go CalculateValue(valueChannel)

    values := <-valueChannel
    fmt.Println(values)

    time.Sleep(1000 * time.Millisecond)
}
/////
Channel Buffered
Calculated Random Value: {} 1
Calculated Random Value: {} 7
7
This executes regardless as the send is now non-blocking
This executes regardless as the send is now non-blocking
