package main

import (
    "testing"
)

func TestMap(t *testing.T) {
    input_value        := "map"
    expected_result    := "ocr"
    actual_result, err := convert_string(input_value)
    if expected_result != actual_result || err != nil {
        t.Fatalf(`Input: "%v", Expected: "%v", Actual "%v"`, input_value, expected_result, actual_result)
    }
}

func TestNonLetters(t *testing.T) {
    input_value        := "12345!@#$%"
    expected_result    := "12345!@#$%"
    actual_result, err := convert_string(input_value)
    if expected_result != actual_result || err != nil {
        t.Fatalf(`Input: "%v", Expected: "%v", Actual "%v"`, input_value, expected_result, actual_result)
    }
}

func TestUppercaseLetters(t *testing.T) {
    input_value        := "THISISATEST"
    expected_result    := "VJKUKUCVGUV"
    actual_result, err := convert_string(input_value)
    if expected_result != actual_result || err != nil {
        t.Fatalf(`Input: "%v", Expected: "%v", Actual "%v"`, input_value, expected_result, actual_result)
    }
}

func TestMixedCase(t *testing.T) {
    input_value        := "ThisIsATest"
    expected_result    := "VjkuKuCVguv"
    actual_result, err := convert_string(input_value)
    if expected_result != actual_result || err != nil {
        t.Fatalf(`Input: "%v", Expected: "%v", Actual "%v"`, input_value, expected_result, actual_result)
    }
}

func TestSpaces(t *testing.T) {
    input_value        := "this is a test"
    expected_result    := "vjku ku c vguv"
    actual_result, err := convert_string(input_value)
    if expected_result != actual_result || err != nil {
        t.Fatalf(`Input: "%v", Expected: "%v", Actual "%v"`, input_value, expected_result, actual_result)
    }
}

func TestMixedLettersNumbers(t *testing.T) {
    input_value        := "this is a 123 test"
    expected_result    := "vjku ku c 123 vguv"
    actual_result, err := convert_string(input_value)
    if expected_result != actual_result || err != nil {
        t.Fatalf(`Input: "%v", Expected: "%v", Actual "%v"`, input_value, expected_result, actual_result)
    }
}

func TestLowerWrapAroundAlphabet(t *testing.T) {
    input_value        := "uvwxyza"
    expected_result    := "wxyzabc"
    actual_result, err := convert_string(input_value)
    if expected_result != actual_result || err != nil {
        t.Fatalf(`Input: "%v", Expected: "%v", Actual "%v"`, input_value, expected_result, actual_result)
    }
}

func TestUpperWrapAroundAlphabet(t *testing.T) {
    input_value        := "UVWXYZA"
    expected_result    := "WXYZABC"
    actual_result, err := convert_string(input_value)
    if expected_result != actual_result || err != nil {
        t.Fatalf(`Input: "%v", Expected: "%v", Actual "%v"`, input_value, expected_result, actual_result)
    }
}
