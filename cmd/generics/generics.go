package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

func Print[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}
