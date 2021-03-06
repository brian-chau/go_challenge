package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
)

func find_unique_characters(input string) (map[string]int, error) {
    result := make(map[string]int)
    for _, i := range input {
        if _, ok := result[string(i)]; ok {
            result[string(i)]++
        } else {
            result[string(i)] = 1
        }
    }
    return result, nil
}

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

    // Find the characters that only happen once
    res, err := find_unique_characters(strings.Join(lines,""))
    letters := make(map[string]int)
    for k, v := range res {
        if v == 1 {
            letters[k] = v
        }
    }

    // Re-parse the characters to get the correct order
    for _, i := range strings.Join(lines,"") {
        if _, ok := letters[string(i)]; ok {
            fmt.Println(string(i))
        }
    }
}
