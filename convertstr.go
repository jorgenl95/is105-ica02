package main

import "fmt"

var str = "53 61 6c 75 74 2c 20 c3 a7 61 20 76 61 20 c2 b0 2d 29 20 c3 87 61 20 63 6f 75 74 65 20 e2 80 8b e2 82 ac 35 30 e2 80 8b"

func main() {
	for i := 0; i<len(str); i++ {
		currentSymbol := str[i]
		if (currentSymbol == " ") {
			currentSymbol = "\x"
		}
	}

	fmt.Println(str)
}