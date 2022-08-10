package game

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"whoacamel/art"
)

/*
	Contains game loop and controls game logic.
*/
func RunGame(textArt map[string][40][1]string, bgArt map[string][40][1]string) {

	// Our lonely player
	player := NewBasePlayer()

	// How long should our trek be?
	ClearScreen()
	art.PrintArt(textArt["treklength"])
	fmt.Printf("How long shall your arduous trek be?\n")
	fmt.Printf("A normal game is around 500 miles. Takes about a half hour.\n")
	fmt.Printf("Enter in 0 for a surprise! Anything from a walk in the park to kingdom come.\n")
	trekLength := TakeAndValidateInt(0, math.MaxInt)

	if trekLength == 0 {
		trekLength = rand.Intn(10000-200) + 200
	}

	// Difficulty? (Not implemented yet but may as well)
	ClearScreen()
	art.PrintArt(textArt["difficulty"])
	fmt.Printf("How arduous shall your long trek be?\n\n")
	fmt.Printf("1. Easy, easy!\n")
	fmt.Printf("2. Can't I have anything normal?\n")
	fmt.Printf("3. Eh, I'm itching for a challenge...\n")
	fmt.Printf("4. BRING IT ON!!!\n")
	difficulty := TakeAndValidateInt(1, 4)

	// Temporarily stuck here forever
	fmt.Printf("%d", difficulty)

	// And last, but not least... our camel's name.
	ClearScreen()
	art.PrintArt(textArt["camel"])
	fmt.Printf("And now, for the most important part...\n")
	fmt.Printf("What is your newly befriended camel's name?\n")
	in := bufio.NewReader(os.Stdin)
	player.Camel.Name, _ = in.ReadString('\n')
	player.Camel.Name = strings.TrimRight(player.Camel.Name, "\n")

	// Now for the story...
	fmt.Printf("\n\nAfter more than a decade of strife, civil war, and an alleged snake bite, the new Caesar Augustus is upon us!\n")
	fmt.Printf("No more shall Rome quake to endless bickering, violence, and peasantry. Finally, the Empire shall rise up to rule its place in history from the jaws of fate!\n")
	fmt.Printf("At least, that's what you'd hear if you were there.\n\n")

	fmt.Printf("You may be a Roman, hell even a respectable veteran legionary, but you don't know it yet. That the Republic you bled for is no more.\n")
	fmt.Printf("You're currently stuck somewhere in the Parthian Empire, on a business trip staking out silky smooth deals with the local merchants.\n")
	fmt.Printf("But alas, as you conclude your haggling and load up your horse with exotic goodies... wait where did he go?\n\n")

	fmt.Printf("Not good, not good! If only you still had your dosh... if only the Chinese bought silk...\n")
	fmt.Printf("Luckily, you spot on the fringes of town a lone camel. Perfect! A chance to get back home!\n")
	fmt.Printf("As you wander around him, you ponder a name that will help you bond with this rather fond drom.\n")
	fmt.Printf("%s seems fine, you suppose.\n", player.Camel.Name)
	fmt.Printf("Just as you finished packing him up, you hear something. No, wait, someone. And they're not happy.\n")
	fmt.Printf("somethingsomethingmycamelsomethinggetarmytogetyoursomethingasssomething\n")
	fmt.Printf("Whatever it is, it's not music to Apollo's ears. In a quick blink, with no thought, you scream YAAAA!!! %s bolts into action, onto the vast sea of sand before you.\n\n", player.Camel.Name)

	fmt.Printf("And now, the legend is yours to continue... after you hit enter...")
	_, _ = in.ReadString('\n')

	// Finally! Now we can play!
	for {
		ClearScreen()

		// Print status
		fmt.Printf("Traveled %d out of the %d mile journey.\n", player.MilesTraveled, trekLength)
		// debug
		player.ShowInventory()
		break
	}
}

// TODO move these functions over to a utils.go??? Maybe one day idk

/*
	Simple(tm) function to validate our constant integer inputs.
	lowerLimit and upperLimit are both inclusive.
*/
func TakeAndValidateInt(lowerLimit int, upperLimit int) int {
	var input int
	for {
		fmt.Printf("\n>")
		_, err := fmt.Scanf("%d", &input)
		if input > upperLimit || input < lowerLimit || err != nil {
			fmt.Printf("Invalid input. Must be a number between %d and %d.", lowerLimit, upperLimit)
		} else {
			break
		}
	}

	return input
}

/*
	Clears the screen according to the OS. What else?
	Windows still needs testing because I don't have that at the moment...
*/
func ClearScreen() {

	cmd := exec.Command("clear")
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	}

	// Ensure that the output of the cmd goes to Go's output
	cmd.Stdout = os.Stdout
	cmd.Run()
}
