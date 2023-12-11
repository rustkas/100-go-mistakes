package main

import "fmt"

var a = func() int {
	fmt.Println("var") // 1
	return 0
}()

func init() {
	fmt.Println("init") // 2
}
func main() {
	fmt.Println("main") // 3
}
