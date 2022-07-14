// Package main provides ...
package main

import (
  "fmt"
  "muffledMitosis/greetings-go"
  "log"
)

func main() {
  log.SetPrefix("greetings: ")
  // log.SetFlags(0)

  message, err := greetings.Hello("Meth")
  if err != nil {
    log.Fatal(err)
  }

  names := []string{"Meth", "Bob", "Alice"}
  messages, err := greetings.Hellos(names)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(messages)
  fmt.Println(message)
}
