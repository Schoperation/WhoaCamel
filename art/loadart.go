package art

import (
	"bufio"
	"fmt"
	"os"
)

/*
	Loads in all of our art from the files. I put them in text files so they're easier to edit.
*/

const (
	MAX_LINES = 40
)

// Returns a map of text art. Who knew capital letters were public and lowercase were private? Weird...
func LoadTexts() map[string][MAX_LINES][1]string {

	art := make(map[string][MAX_LINES][1]string)

	// Loading all of the texts onto the map...
	art["title"] = loadArtFromFile("title", "text")

	return art
}

// Returns a map of BackGround art.
func LoadBgs() map[string][MAX_LINES][1]string {

	art := make(map[string][MAX_LINES][1]string)

	// Loading all of the texts onto the map...
	art["desert"] = loadArtFromFile("desert", "bg")

	return art
}

// Prints out a piece of art.
func PrintArt(art [MAX_LINES][1]string) {

	i := 0
	for i < MAX_LINES && art[i][0] != " " {
		fmt.Println(art[i][0])
		i++
	}
}

// Opens a text file and turns the art into an array.
func loadArtFromFile(fileName string, folder string) [MAX_LINES][1]string {

	file, err := os.Open("art/" + folder + "/" + fileName + ".txt")

	if err != nil {
		fmt.Println(err)
		file.Close()
		return [MAX_LINES][1]string{{"Error, couldn't open file"}}
	}

	defer file.Close()

	// Time to read the file line by line and put it in the 2d array
	var art = [MAX_LINES][1]string{}
	i := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		art[i][0] = scanner.Text()
		i++
	}

	// Fill the rest with blank spaces
	for i = i; i < MAX_LINES; i++ {
		art[i][0] = " "
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return art
}
