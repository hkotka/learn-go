package main

import (
	"fmt"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Warm Up
//
//  Starting with this exercise, you'll build a command-line
//  game store.
//
//  1. Declare the following structs:
//
//     + item: id (int), name (string), price (int)
//
//     + game: embed the item, genre (string)
//
//
//  2. Create a game slice using the following data:
//
//     id  name          price    genre
//
//     1   god of war    50       action adventure
//     2   x-com 2       30       strategy
//     3   minecraft     20       sandbox
//
//
//  3. Print all the games.
//
// EXPECTED OUTPUT
//  Please run the solution to see the output.
// ---------------------------------------------------------
type item struct {
	id    int
	name  string
	price int
}

type game struct {
	genre string
	item
}

func main() {
	games := []game{
		{"action adventure", item{1, "god of war", 50}},
		{"strategy", item{2, "x-com", 30}},
		{"sandbox", item{3, "minecraft", 20}},
	}

	fmt.Printf("%-4s %-12s %-10s %-5s\n", "id", "name", "price", "genre")
	fmt.Println(strings.Repeat("-", 36))
	for _, game := range games {
		fmt.Printf("%-4d %-12s %-10d %-5s\n", game.id, game.name, game.price, game.genre)
	}
}
