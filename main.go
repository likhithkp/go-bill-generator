package main

import (
	"fmt"
	"strings"
)

func getFirstLettersInCaps(n string) (string, string) {
	arr := strings.Split(n, " ")

	var names []string

	for _, value := range arr {
		names = append(names, value[:1])
	}

	if len(names) > 1 {
		return strings.ToUpper(names[0]), strings.ToUpper(names[1])
	}

	return names[0], "_"
}

func main() {
	a, b := getFirstLettersInCaps("likhith kp")
	c, d := getFirstLettersInCaps("shiba")

	fmt.Println(a, b)
	fmt.Println(c, d)
}
