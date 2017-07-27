package main

func main() {
 distance := 123
 winner := race(distance, baby{}, developer{clumsiness: 1.0})
 println("the fastest to run", distance, "was", winner)
}

type runner interface {
 name() string
 run(distance int) (seconds int)
}

func race(distance int, runners ...runner) (winner string) {
 var fastest int
 for i, runner := range runners {
  seconds := runner.run(distance)
  if i == 0 || seconds < fastest {
   fastest, winner = seconds, runner.name()
  }
 }
 return winner
}

type runbot9000 struct{}

func (runbot9000) name() string { return "RunBot 9000" }
func (runbot9000) run(d int) int { return d }

type baby struct{}

func (baby) name() string { return "A little baby" }
func (baby) run(d int) int { return 10*d }

type developer struct{ clumsiness float64 }

func (d developer) name() string { return "An Developer" }
func (d developer) run(dist int) int { return int(3 * (1-d.clumsiness) * float64(dist)) }
