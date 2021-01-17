
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
    fmt.Println("Only Executes after another goroutine performs a receive on the channel")
}

func main() {
    fmt.Println("Channel Unbuffered")

    valueChannel := make(chan int)
    defer close(valueChannel)

    go CalculateValue(valueChannel)
    go CalculateValue(valueChannel)

    values := <-valueChannel
    fmt.Println(values)
}

//////
Calculated Random Value: {} 1
Calculated Random Value: {} 7
1
Only Executes after another goroutine performs a receive on the channel
