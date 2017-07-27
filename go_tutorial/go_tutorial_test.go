package main

import (
	"testing"
  "fmt"
)

func TestSquare(t *testing.T) {
  for input, want := range map[int]int{
   // table based tests 
    2: 4,
    3: 9,
    4: 16,
  } {
  name := fmt.Sprintf("square(%d)", input)
  t.Run(name, func(t *testing.T) {
      if have := square(input); have!= want {
        t.Errorf("want %d, have %d", want, have)
      }
    })
  }
}
