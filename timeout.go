package main
// this is not even close to working
func main() {
  select {
  case <- c:
    println("done")
  case <-time.After(time.Second * 1.5):
    println("timeout")
  }
}

func foo() int {
  seconds := rand.Intn(3)
  time.After(seconds)
}
