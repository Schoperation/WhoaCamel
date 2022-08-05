package game

import (
	"fmt"
	"math/rand"
	"time"
)

/*
	Where details about the player are stored.
	Includes inventory, status (thirst), and camel status.
*/
type Player interface {
	ShowWarnings()
	ShowInventory()
	HasItemInInventory(name string) (bool, []Item)
	DrinkFromCanteen()
	Travel(speed int)
	Scout()
}

// A basic item that can be in the player's inventory.
type Item struct {
	Name  string
	Desc  string
	Data  int
	Data2 int
}

// A camel, the player's trusty steed.
type Camel struct {
	Name         string
	TopSpeed     int
	TopEndurance int
	Thirst       int // They can last 7 - 10 days without water, on average. Max is 300.
	Energy       int // Refer to BasePlayer.Energy, but doubled
}

// Main player struct
type BasePlayer struct {
	Thirst        int // 0 not thirsty, 100 so parched everything looks liquid
	Energy        int // 100 not tired, 0 so tired everything looks like tires... what?
	Inventory     []Item
	Camel         Camel
	MilesTraveled int
}

// Constructor
func NewBasePlayer() *BasePlayer {

	// Start with a canteen
	var canteen = Item{
		"Canteen",
		"Your trusty companion that stays loyal even in the harshest of desert storms. Bear in mind though, that it's no miracle worker.",
		5, // Number of sips
		5, // Max number of sips
	}

	// And a camel
	var camel = Camel{
		".",
		rand.Intn(25-10) + 10, // (max - min) + min
		rand.Intn(20-10) + 10,
		0,
		200,
	}

	return &BasePlayer{
		0,
		100,
		[]Item{canteen},
		camel,
		0,
	}
}

// Runs everyday to display warning messages about thirst, energy, etc...
func (p *BasePlayer) ShowWarnings() {
	fmt.Printf("Warning, need to complete a function!")
}

// Prints out their inventory.
func (p *BasePlayer) ShowInventory() {
	fmt.Printf("You look into your pouch and, other than stray grains of sand, you find...\n")
	for i, item := range p.Inventory {
		fmt.Printf("#%d: %s -- %s\n", i, item.Name, item.Desc)
	}
}

// Returns whether or not the player has a particular item. Can use the bool for a simple check, or the array for the actual items to change them. Or both!
func (p *BasePlayer) HasItemInInventory(name string) (bool, []Item) {
	var hasIt bool = false
	var list []Item
	for _, item := range p.Inventory {
		if item.Name == name {
			hasIt = true
			list = append(list, item)
		}
	}

	return hasIt, list
}

// Drinks from a canteen.
func (p *BasePlayer) DrinkFromCanteen() {
	// For the suspense! Whoa!
	fmt.Printf("You unscrew the rusty cap, lift the mini fountain of youth up and...\n")
	time.Sleep(4 * time.Second)

	// Get all canteens
	hasCanteen, canteens := p.HasItemInInventory("Canteen")

	if !hasCanteen {
		fmt.Printf("You wake up from your daydream. What was that thing in your hand, anyway?\n")
		return
	}

	// Go through each canteen, in case the first one(s) have no sips.
	for _, canteen := range canteens {
		if canteen.Data <= 0 {
			fmt.Printf("Hm, not a measly drop in that one...\n")
			time.Sleep(1 * time.Second)
		} else {
			canteen.Data -= 1
			p.Thirst -= 50
			fmt.Printf("Ahhhhhh, a cool escape for the parched throat.\n")

			if canteen.Data == 1 {
				fmt.Printf("Just a single measly sip left. Oh dear...\n")
			} else {
				fmt.Printf("You peer into your canteen and see about %d sips floating around.\n", canteen.Data)
			}

			break
		}
	}
}

// Most useful function ever, who would've thought? Makes the player and their camel travel.
func (p *BasePlayer) Travel(speed int) {
	/*
		3 Speeds: Slow 1, Moderate 2, Full 3
		No need for break in Go, apparently... hm...
	*/
	var baseMiles int

	switch speed {
	case 1:
		baseMiles = rand.Intn(10-5) + 5
		fmt.Printf("Winning a race so soon?\n")
		p.Energy -= rand.Intn(15-5) + 5
		p.Thirst += rand.Intn(15-5) + 5
	case 2:
		baseMiles = rand.Intn(min(p.Camel.TopSpeed, 15)-8) + 8
		fmt.Printf("Tried and true, you suppose.\n")
		p.Energy -= rand.Intn(25-15) + 15
		p.Thirst -= rand.Intn(25-15) + 15
	default:
		baseMiles = rand.Intn(p.Camel.TopSpeed-12) + 12
		fmt.Printf("FULL SPEED AHEAD!!! YA CAMEL!!!\n")
		p.Energy -= rand.Intn(40-25) + 25
		p.Thirst -= rand.Intn(40-25) + 25
	}

	time.Sleep(3 * time.Second)
	fmt.Printf("You've traveled %d miles today.\n", baseMiles)

	/*
		Calculate how tired camel is

		Compare miles traveled to endurance.
		If miles > endurance, multiply difference by 10 and subtract that from the camel's energy.
		If miles <= endurance, subtract a base energy of ~10.
	*/
	if baseMiles <= p.Camel.TopEndurance {
		p.Camel.Energy -= rand.Intn(15-10) + 10
	} else {
		diff := baseMiles - p.Camel.TopEndurance
		p.Camel.Energy -= rand.Intn(diff*10-diff*7) + diff*7
	}
}

// Scout out the area for an oasis, town, death trap, anything really. Coming soon
func (p *BasePlayer) Scout() {
	fmt.Printf("Eh, today's not the day.\n")
}

// Go's internal function only takes in floats... I don't wanna cast or nothing, okay?
func min(first int, second int) int {
	if first < second {
		return first
	} else {
		return second
	}
}
