package main

import (
  "fmt"
)

func accountDetails(userId *int, ifsc *int) {
  fmt.Println("%d  %d\n", *userId, *ifsc)
  fmt.Println("from func")
}

func main() {
  userId := 1112
  ifsc := 24
  
  accountDetails(&userId, &ifsc)
  fmt.Prinlnln("from main")
  
}

////////////////////////////

1112  24
from func
from main


///////////////////////////

func accountDetails(userId *int, ifsc *int) {
  defer fmt.Println("defer from func")
  if userId == nil{
    panic("userId can't be nil")
  }
  if ifsc == nil{
    panic("ifsc can't be nil")
  }
  fmt.Println("%d  %d\n", *userId, *ifsc)
  fmt.Println("from func")
}

func main() {
  defer fmt.Println("defer from main")
  userId := 1112
  //ifsc := 24
  
  accountDetails(&userId, nil)
  fmt.Prinlnln("from main")
  
}

/////////////////////////////

defer from func
defer from main
ifsc can't be nil

////////////////////////////

once panic occured execution stop but check for defer to print then panic will be printed at end
