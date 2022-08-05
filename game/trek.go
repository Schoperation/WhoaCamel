package game

import "fmt"

/*
	Contains game loop and controls game logic.
*/
func RunGame() {

	// Our lonely player
	player := NewBasePlayer()

	fmt.Printf("%d", player.Energy)
}
