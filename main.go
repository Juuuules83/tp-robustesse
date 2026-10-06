package main

import "fmt"

func main() {
	ligne := []int{0, 2, 4, 6, 8}
	v := 4
	searchV1, err := SearchV1(ligne, v)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(searchV1)
}
