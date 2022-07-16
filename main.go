package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"whoacamel/art"
)

func main() {

	// Load all of our stuff
	fmt.Println("Loading stuff...")

	textArt := art.LoadTexts()
	bgArt := art.LoadBgs()

	// Just so go doesn't complain wheeh you ain't using it wheeeehhhhh
	art.PrintArt(bgArt["desert"])

	clearScreen()

	// Say our "welcome"
	art.PrintArt(textArt["title"])
	fmt.Println("...is what you'll be saying as you cruise through this game!")
	fmt.Println("Well, 'cruise' might be too strong of a word... oh well.")
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
