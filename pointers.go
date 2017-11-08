package main

import "fmt"


func main() {
  i := 1
  fmt.Println(i)
  var ptr *int = &i
  *ptr = 0
  fmt.Println(i)
}
