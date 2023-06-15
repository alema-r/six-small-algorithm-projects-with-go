package main

import "fmt"

func main() {
	bytes, err := fmt.Println("Will this be printed without errors?")
	fmt.Printf("%d bytes written to stdout, err:%v", bytes, err)
}
