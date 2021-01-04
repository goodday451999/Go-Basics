package main
import (
  "fmt"
)
func doSum(x, y int) int{
  defer fmt.Println("Func doSum execution is done")
  sumAns := x + y
  fmt.Println("Sum is done")
  return sumAns
}
func main(){
  x := 2
  y := 3
  sum := doSum(x, y)
  fmt.Println(sum)
}
