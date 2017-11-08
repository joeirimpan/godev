package main


import s "strings"
import "fmt"


func main() {
  fmt.Println(s.Join([]string{"joe", "paul"}, " "))
  fmt.Println(len("joe"))
  var split_words= s.Split("joe-paul", "-")
  for _, word := range split_words {
    fmt.Println(word)
  }
  fmt.Println(split_words[0])
}
