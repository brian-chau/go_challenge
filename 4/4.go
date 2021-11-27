package main

import (
    "fmt"
    "io"
    "log"
    "os"
    "net/http"
    "regexp"
    "strconv"
)

func main() {
    param := "12345"
    url := "http://pythonchallenge.com/pc/def/linkedlist.php?nothing="
    for true {
        // Do an HTTP Get request
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

        // Check for special conditions
        if res == "peak.html" {
            fmt.Println("peak.html")
            os.Exit(3)
        } else if res == "Yes. Divide by two and keep going." {
            tmp, err := strconv.Atoi(param)
            if err != nil {
                log.Fatalf("Could not convert number to string, for some reason.")
            }
            tmp     = tmp / 2
            param   = strconv.Itoa(tmp)
            continue
        }

        // Otherwise, the standard condition is "and the next nothing is <some number>"
        next_pattern := regexp.MustCompile("and the next nothing is (\\d+)")
        next_param   := next_pattern.FindAllStringSubmatch(res,-1)
        for _, i := range next_param {
            fmt.Printf("%s\n", i[1])
        }
        param       = next_param[0][1]
    }
}
