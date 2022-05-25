package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func WhatNumberToConvert(a string) string {
	Number := bufio.NewReader(os.Stdin)
	fmt.Println("Can you Please enter a number")
	num, err := Number.ReadString('\n')

	if err != nil {
		fmt.Println(err)
	}
	newNum := strconv.Atoi(num)
}

func main() {
	ans := WhatNumberToConvert(num)
	fmt.Println("The Negative value of ", ans)
}
