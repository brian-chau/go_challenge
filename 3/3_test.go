package main

import (
    "fmt"
    "reflect"
    "testing"
)

func TestMap(t *testing.T) {
    input_value         := "banana"

    expected_result := make(map[string]int)
    expected_result["b"] = 1
    expected_result["n"] = 2
    expected_result["a"] = 3

    actual_result, err  := find_unique_characters(input_value)
    v := fmt.Sprintf("%T", actual_result)
    switch v {
        case "map[string]int":
            if !reflect.DeepEqual(expected_result, actual_result) || err != nil {
                t.Fatalf(`Input: "%v", Expected: "%v", Actual "%v"`, input_value, expected_result, actual_result)
            }
        default:
            t.Fatalf(`Invalid type! Expected type: "%T", Actual type: "%T"`, expected_result, actual_result)
    }
}

