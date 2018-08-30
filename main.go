package main

import (
  "fmt"
  "bufio"
  "os"
)

func main() {
  scanner := bufio.NewScanner(os.Stdin)
  var text string
  var compareText string = os.Args[1]

  for text != "q" {
    fmt.Print("$ ")
    scanner.Scan()
    text = scanner.Text()
    if text != "q"{
      if text == compareText {
        fmt.Println(": ", text)
        fmt.Println("\033[H\033[2J")
      }else{
        fmt.Println("\033[H\033[2J")
        fmt.Println(compareText, ".\n")
      }
    }
  }
}
