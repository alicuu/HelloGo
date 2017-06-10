package main

import "fmt"

func main() {
	ele := make(map[string]string)
	ele["Danny"] = "Peking"
	fmt.Println(ele)

	ele2 := ele
	ele2["Danny"] = "Jinan"
	fmt.Println(ele)
	fmt.Println(ele2)
}
