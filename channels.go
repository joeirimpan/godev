package main

import "fmt"
import s "strings"

func main() {

    messages := make(chan string)

    var message string
    fmt.Scanln(&message)

    go func() { messages <- message }()
    go func() {
      msg := <- messages
      if s.ToLower(msg) == "ping" {
        fmt.Println("PONG")
      }
    }()

    var exit string
    fmt.Scanln(&exit)
}
