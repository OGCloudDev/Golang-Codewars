package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "ADDICTED"
	if s == strings.ToUpper(s) {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}
