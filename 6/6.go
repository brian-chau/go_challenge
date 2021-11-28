package main

import (
    "bufio"
    "fmt"
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
    pattern  := regexp.MustCompile(`Next nothing is (\d+)`)
    filename := "channel/90052.txt"
    for true {
        fmt.Println(filename)

        lines, _ := read_lines(filename)
        match    := pattern.FindAllStringSubmatch(strings.Join(lines,""),-1)
        if len(match) == 0 {
            fmt.Println(lines)
            break
        }

        var next_file string
        for _, i := range match {
            next_file = i[1]
        }
        filename = "channel/" + next_file + ".txt"
    }
}
