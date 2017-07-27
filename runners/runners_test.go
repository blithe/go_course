package main

import (
	"testing"
  "fmt"
)

func TestRun(t *testing.T) {
}

func TestRunners(t *testing.T) {
  runbot := runbot9000{}
  baby := baby{}
  developer := developer{}
  distance := 123

  for runner, expected := range map[runner]int {
   // table based tests 
    runbot: distance,
    baby: 1230,
    developer: 369,
  } {
    name := fmt.Sprintf("%s.run(%d)", runner.name(), distance)
    t.Run(name, func(t *testing.T) {
        if actual := runner.run(distance); actual != expected {
          t.Errorf("expected %d, got %d", expected, actual)
        }
      },
    )
  }
}

func BenchmarkRunbot(b *testing.B) {
  runbot := runbot9000{}
  distance := 123
	for i := 0; i < b.N; i++ {
		runbot.run(distance)	
	}
}
func BenchmarkBaby(b *testing.B) {
  baby := baby{}
  distance := 123
	for i := 0; i < b.N; i++ {
		baby.run(distance)	
	}
}
func BenchmarkDeveloper(b *testing.B) {
  developer := developer{}
  distance := 123
	for i := 0; i < b.N; i++ {
		developer.run(distance)	
	}
}
