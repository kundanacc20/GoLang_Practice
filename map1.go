package main

import "fmt"

func main() {
	mapVar := make(map[string]int)

	mapVar["apple"] = 1
	mapVar["orange"] = 2
	mapVar["banana"] = 3

	fmt.Println(mapVar)

	fmt.Println("value of banana is:", mapVar["banana"])
}
