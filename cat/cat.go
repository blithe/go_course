package main

import (
  "fmt"
  ioutil "io/ioutil"
)

func main() {
  Write("text.txt")
}

func Write(filename string) {
  contents, err := ioutil.ReadFile(filename)
  if err != nil {
    panic(err)
  }
  fmt.Printf("%s\n", contents)
}
