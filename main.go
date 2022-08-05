package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"time"
	"whoacamel/art"
	"whoacamel/game"
)

func main() {

	// Load all of our stuff
	fmt.Println("Loading stuff...")

	textArt := art.LoadTexts()
	bgArt := art.LoadBgs()

	// Just so go doesn't complain wheeh you ain't using it wheeeehhhhh
	art.PrintArt(bgArt["desert"])

	// Set randomizer seed
	rand.Seed(time.Now().UnixMilli())

	clearScreen()

	// Say our "welcome"
	art.PrintArt(textArt["title"])
	fmt.Println("...is what you'll be saying as you cruise through this game!")
	fmt.Println("Well, 'cruise' might be too strong of a word... oh well.")

	// TODO allow for miles input and difficulty (implemented later down the line, but may as well have it in the system)

	// Drum roll please...
	game.RunGame()
}

func clearScreen() {

	cmd := exec.Command("clear")
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	}

	// Ensure that the output of the cmd goes to Go's output
	cmd.Stdout = os.Stdout
	cmd.Run()
}
