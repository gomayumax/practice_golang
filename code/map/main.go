package main

import "fmt"

func main() {
	//colors := make(map[string]string)
	//colors["white"] = "#fff"
	//delete(colors, "white")

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, hex)
	}
}
