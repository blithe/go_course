package main
// this doesn't work either
import (
  "sync"
	"bufio"
	"os"
	"fmt"
	http "net/http"
)

func main() {
  foo()
}

func foo() {
  var wg sync.WaitGroup
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	urls, _ := reader.ReadString('\n')
	println(urls)

  for _, url := range urls {
    wg.Add(1)
    go func() {
      defer wg.Done()
      get(url)
    }()
  }

  wg.Wait()
}

func get(url string) string {
  var time int
  client := http.Client{}

  request, err := http.NewRequest("GET", url, nil)
  if err != nil {
    panic(err)
  }

  response, err := client.Do(request)
  if err != nil {
    panic(err)
  }

  return response
}
