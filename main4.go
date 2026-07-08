package main

import "fmt"

type weekday int

const (
	sunday weekday= iota
	monday
	tuesday
	wednesday
	thursday
	friday
	saturday
)

func main() {
fmt.Println(wednesday)

}