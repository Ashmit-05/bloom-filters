package main

import (
	"fmt"
)

func main() {
	b, err := CreateBloomFilter(16)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		keys := []string{"bat", "rat", "cats"}
		for _, key := range keys {
			b.Add(key)
		}
		for _, key := range keys {
			fmt.Println(key, b.Contains(key))
		}
		fmt.Println("cat", b.Contains("cat"))
		b.Clear()
		fmt.Println("After clearing..")
		fmt.Println("bat", b.Contains("bat"))
	}
}
