package main

import "fmt"

func main() {
	myBill := newBill("Likhith")
	fmt.Println(myBill.format())

	myBill.format()
}
