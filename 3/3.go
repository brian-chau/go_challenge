package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "regexp"
    "strings"
)

func read_lines(path string) ([]string, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines, scanner.Err()
}

func main() {
    if len(os.Args) != 2 {
        fmt.Println("Expecting filename!")
        os.Exit(3)
    }

    lines, err := read_lines(os.Args[1])
    if err != nil {
        log.Fatalf("Could not handle filename!")
    }

    re := regexp.MustCompile(`[^A-Z][A-Z]{3}([a-z]{1})[A-Z]{3}[^A-Z]`)
    res := re.FindAllStringSubmatch(strings.Join(lines,""),-1)
    for _, i := range res {
        fmt.Printf("%s", i[1])
    }
    fmt.Println()
}
