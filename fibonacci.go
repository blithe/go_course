package main

func main() {
  c := make(chan int)
  go fibonacci(10, c)
  for i := range c {
    println(i)
  }
}

func fibonacci(n int, c chan int) {
  for i, a, b := 0, 1, 1; i <= n; i++ {
    c <- a
    a, b = b, a+b
  }
  close(c)
}
