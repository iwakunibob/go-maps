package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
		"blue":  "#0000FF",
	}
	grayscale := make(map[string]string)
	grayscale["white"] = "#FFFFFF"
	grayscale["black"] = "#000000"
	grayscale["gray"] = "#808080"
	fmt.Println(colors)
	fmt.Println(grayscale)
	delete(grayscale, "black")
	fmt.Println(grayscale)
	printMap(colors)
	printMap(grayscale)
}
func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
