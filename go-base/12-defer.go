package main
import (
  "fmt"
)
func doSum(x, y int) int{
  defer fmt.Println("Func doSum execution is done")  
  /* acts as LIFO.. Last In First Out
    To close file
  */
  defer fmt.Println("Seems the execution is ended")
  sumAns := x + y
  fmt.Println("Sum is done")
  return sumAns
}
func main(){
  x := 2
  y := 3
  sum := doSum(x, y)
  fmt.Println(sum)
  
  a := "start"
  defer fmt.Println(a) // start not end
  a = "end"
}
