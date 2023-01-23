package main

import "fmt"

func main() {
	// Different syntax to declare maps
	// #1. colors := map[string]string{
	// "red": "#ff0000",
	// "green": "#47nf8j",
	// }
	// #2. var colors map[string]string
	// #3. colors := make(map[string]string)
	// colors["white"] = "#ffffff"
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#47nf8j",
		"white": "#ffffff",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
