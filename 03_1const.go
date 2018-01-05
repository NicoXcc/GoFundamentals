package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	inputdelimiter = '\n'
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString(inputdelimiter)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(input)
}
