package main

import (
	"fmt"
)

func main() {
	name := map[string]string{
		"name1": "Likhith",
		"name2": "Shiba",
		"name3": "Lauda",
	}

	for k, v := range name {
		fmt.Println(k, v)
	}

}
