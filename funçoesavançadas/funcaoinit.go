package main

import "fmt"

var n int

func init() {
	fmt.Println("init sendo executada")
	n = 19
}

func main() {
	fmt.Println("main em execuçao")
	fmt.Println(n)
}
