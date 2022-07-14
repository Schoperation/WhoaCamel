package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func main() {

	// Load all of our stuff
	fmt.Println("Loading stuff...")

	clearScreen()

	// Say our "welcome"
	fmt.Println("Whoa Camel!")
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
