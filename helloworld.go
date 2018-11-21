package main

import "fmt"

func main() {
	fmt.Println("Hello world!!!")
	fmt.Println("I am Coming here")

	i := 0
	for {
		i++
		if i == 10 {
			break
		}

		println("Print line : ", i)
	}

	println("++++++++++++++++++++++++")
	for j := 0; j < 10; j++ {
		if j%2 == 0 {
			continue
		}
		println("Print line : ", j)
	}
}
