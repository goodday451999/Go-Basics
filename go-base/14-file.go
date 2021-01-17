package main

import (
  "fmt"
  "os"
)

func main (){
  f := createFile("/Home/code.txt")
  defer closeFile(f)
  wrteFile(f)
}

func createFile(path string) *os.File {
  defer fmt.Println("Created")
  f, err := os.Create(path)
  if err != nil {
    panic(err)
  }
  return f
}

func writeFile(f *os.File) {
  fmt.Println("Writing")
  fmt.Fprintln(f, "data 1")
}

func closeFile(f *os.File) {
  defer fmt.Println("Closed")
  err := f.Close()
  
  if err != nil {
    fmt.Fprintf(os.Stderr, "error : %v\n", err) ////prints initial value of the variable  
    os.Exit(1)
  }
}
  

/////////////////////////////////////////////////////////////////////////////////////////////

func main (){
  f := createFile("") // no path is defined
  defer closeFile(f)
  wrteFile(f)
}

func createFile(path string) *os.File {
  fmt.Println("Creating")         // creating - will be printing - same as -
  defer fmt.Println("Created")    // created
  f, err := os.Create(path)
  if err != nil {
    panic(err)                    // then the panic err will be printed
  }
  return f
}


// once panic is encountered the whole prog. will stop
