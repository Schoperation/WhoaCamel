package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
	"whoacamel/art"
	"whoacamel/game"
)

func main() {

	// Load all of our stuff
	fmt.Printf("Loading stuff...\n")
	beginTime := time.Now()

	textArt := art.LoadTexts()
	bgArt := art.LoadBgs()

	// Just so go doesn't complain wheeh you ain't using it wheeeehhhhh
	art.PrintArt(bgArt["desert"])

	// Set randomizer seed
	rand.Seed(time.Now().UnixMilli())

	loadTime := time.Since(beginTime).String()
	fmt.Printf("Loaded everything in %s\n", loadTime)
	game.ClearScreen()

	// Say our "welcome"
	art.PrintArt(textArt["title"])
	fmt.Printf("...is what you'll be saying as you cruise through this game!\n")
	fmt.Printf("Well, 'cruise' might be too strong of a word... oh well.\n\n")

	fmt.Printf("1. Start Your Trek!\n")
	fmt.Printf("2. Quit To OS, Before You Even Began (boooo...)!\n")

	menuChoice := game.TakeAndValidateInt(1, 2)
	if menuChoice == 1 {
		// Drum roll please...
		game.RunGame(textArt, bgArt)
	} else {
		fmt.Printf("In the dark vastness of the terminal, I find solace.\n")
		os.Exit(0)
	}
}
