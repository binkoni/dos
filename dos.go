package main
import "fmt"
import "net/http"
import "os"
import "strconv"
import "sync"

func dos(wg *sync.WaitGroup, url string, id int) {
    defer wg.Done()
    request, _ := http.NewRequest("GET", url, nil)
    client := &http.Client{}
    response, _ := client.Do(request)
    fmt.Println(response.Status, id)
}

func main() {
    var url string = os.Args[1]
    var dosCount int
    dosCount, _ = strconv.Atoi(os.Args[2])

    if dosCount > 1000 {
        dosCount = 1000
    }

    var wg sync.WaitGroup

    for id := 0; id < dosCount; id++ {
        wg.Add(1)
        go dos(&wg, url, id)
    }

    wg.Wait()
    fmt.Println("Done")
}