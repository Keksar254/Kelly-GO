package main

import (
    "fmt"
    "unicode/utf8"
)

func main() {
    s := "Hëllo"
    fmt.Println("len(s):", len(s))                     // byte length
    fmt.Println("RuneCount:", utf8.RuneCountInString(s)) // actual character count

    for i, r := range s {
        fmt.Printf("index %d: %c (%d)\n", i, r, r)
    }
}