package main

import (
        "crypto/tls"
        "fmt"
        "net/http"
        "os"
        "strconv"
        "sync"
        "time"
)

var url string
var counter int
var smash int

func main() {
        var wg sync.WaitGroup
        if len(os.Args[1:]) != 2 {
                fmt.Println("URL and smash count is required to smash!")
                os.Exit(1)
        }
        url = os.Args[1]
        smash, _ = strconv.Atoi(os.Args[2])
        // TODO: make this optional - insecure by design...
        http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
        go timer()
        for i := 0; i < smash; i++ {
                wg.Add(1)
                go hit(url, &wg)
                counter += 1

        }
        wg.Wait()
        fmt.Printf("%d http.Get's were smashed! at %s", counter, url)
}

func hit(url string, wg *sync.WaitGroup) {
        defer wg.Done()
        resp, err := http.Get(url)
        if err != nil {
                fmt.Println(err)
        }
        if resp.StatusCode != 200 {
                fmt.Println("not 200")
                os.Exit(1)
        }
}

func timer() {
        for {
                fmt.Println(time.Now().Format("15:04:05"))
                time.Sleep(1000 * time.Millisecond)
        }
}
