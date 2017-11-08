package main


import "fmt"
import s "strings"
import "time"


func threaded_func(message string) {
  for i := 0; i < 10; i++ {
    repeat := s.Repeat(".", i)
    fmt.Println(message, repeat)
  }
}


func main() {
  go threaded_func("LOADING")
  go func(message string) {
    time.Sleep(time.Nanosecond)
    fmt.Println(message)
  }("<LOADING>")
  var input string
  fmt.Scanln(&input)
}
