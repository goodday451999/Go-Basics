package main

import (
  "fmt"
)

func recoverFunc(){
  if r := recover(); r != nil {
    fmt.Println("recovered from \n", r)
    fmt.Println("recover executed")
  }
}

func accountDetails(id *int, ifsc *int) {
  //defer fmt.Println("defer from func")
  defer recoverFunc()
  
  if id == nil {
    panic("id can't be nil")
  }
  if ifsc == nil {
    panic("ifsc can't be nil")
  }
  fmt.Printf("%d, %d\n", *id, *ifsc)
  fmt.Println("from func") 
}
  

func main() {
  defer fmt.Println("defer from main")
  id := 12
  ifsc := 20
  
  accountDetails(&id, nil) //&ifsc
  fmt.Println("from main")
}

/////////////////////////////////

recovered from ifsc can't be nil
recover executed
from main
defer from main

///////////////////////////////////

recover is to keep on the execution of main function though the panic is occured 
