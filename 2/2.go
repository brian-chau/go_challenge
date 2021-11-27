package main

import (
    "fmt"
    "log"
)

func convert_string(input string) (string, error) {
    result := ""

    for _, i := range input {
        if i < 'a' || i > 'z' {
            if i < 'A' || i > 'Z' {
                result += string(i)
                continue
            } else {
                result += fmt.Sprintf("%s", string(rune(((int(i) - 'A' + 2) % 26) + 'A')))
            }
        } else {
            result     += fmt.Sprintf("%s", string(rune(((int(i) - 'a' + 2) % 26) + 'a')))
        }
    }

    return result, nil
}

func main() {
    res, err := convert_string("g fmnc wms bgblr rpylqjyrc gr zw fylb. rfyrq ufyr amknsrcpq ypc dmp. bmgle gr gl zw fylb gq glcddgagclr ylb rfyr'q ufw rfgq rcvr gq qm jmle. sqgle qrpgle.kyicrpylq() gq pcamkkclbcb. lmu ynnjw ml rfc spj.")
    if err != nil {
        log.Fatalf("This shouldn't happen\n")
    }
    fmt.Println(res)

    res, err = convert_string("map")
    if err != nil {
        log.Fatalf("This shouldn't happen\n")
    }
    fmt.Println(res)
}
