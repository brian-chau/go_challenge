package main

import (
    "fmt"
    "io"
    "log"
    "os"
    "net/http"
    "regexp"
    "strconv"
    "strings"
)

func read_html(url string, param string) (string) {
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
    return string(body)
}

func check_special_conditions(body string, param string) (string, bool) {
    // "peak.html" is the end condition
    if body == "peak.html" {
        fmt.Println("peak.html")
        os.Exit(3)
    } else if strings.Contains(body, "Divide by two") {
        // Calculate the next step, if it says to divide by two
        tmp, err := strconv.Atoi(param)
        if err != nil {
            log.Fatalf("Could not convert number to string, for some reason.")
        }
        tmp     = tmp / 2
        param   = strconv.Itoa(tmp)
        return param, true
    }
    return "", false
}

func main() {
    param := "12345"
    url   := "http://pythonchallenge.com/pc/def/linkedlist.php?nothing="
    for true {
        // Do an HTTP Get request
        body := read_html(url, param)

        // Check for special conditions
        tmp, found := check_special_conditions(body, param)
        if found {
            param = tmp
            continue
        }

        // Otherwise, the standard condition is "and the next nothing is <some number>"
        next_pattern := regexp.MustCompile("and the next nothing is (\\d+)")
        next_param   := next_pattern.FindAllStringSubmatch(body,-1)
        for _, i := range next_param {
            fmt.Printf("%s\n", i[1])
        }
        param       = next_param[0][1]
    }
}
