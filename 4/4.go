package main

import (
    "fmt"
    "io"
    "log"
    "net/http"
    "regexp"
)

func main() {
    go_next := true
    //param := "12345"
    param := "8022"
    url := "http://pythonchallenge.com/pc/def/linkedlist.php?nothing="
    for go_next == true {
        resp, err := http.Get(url + param)
        if err != nil {
            log.Fatalf("Could not open URL")
        }
        defer resp.Body.Close()
        body, err := io.ReadAll(resp.Body)

        if err != nil {
            log.Fatalf("Could not read body")
        }
        res := string(body)

        re         := regexp.MustCompile("and the next nothing is (\\d+)")
        next_param := re.FindAllStringSubmatch(res,-1)
        for _, i := range next_param {
            fmt.Printf("%s\n", i[0])
            fmt.Printf("%s\n", i[1])
        }
        param       = next_param[0][1]
    }
}
