package main

import "fmt"

func main() {
	a := []string{}
	a = append(a, "Likhith", "Shiba", "Shibashanth")
	rangeOne := a[1:2]
	fmt.Println(a)
	fmt.Println(rangeOne)
}
