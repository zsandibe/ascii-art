package main

import (
	"fmt"
	"os"
	"zsandibe/banner"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	arg := os.Args[1]
	for _, ch := range arg {
		if ch > 127 {
			fmt.Println("Dude,write ASCII characters!!!")
			return
		}
	}
	inputFile := "standard.txt"
	standBanner, err := banner.ReadBanner(inputFile)
	if err != nil {
		fmt.Println("Error: %v", err)
		return
	}
	banner.PrintBanner(arg, standBanner)
}
