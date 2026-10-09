package validation

import "fmt"

const usage = "Usage: go run . <filename.txt>"

func IsValid() {
	// We read the file
	input := parseFile()
	if input == nil {
		return
	}

	// We check the ants
	if !ants(input) {
		return
	}

	// We remove the comments #
	input = commentRemover(input)

	// We check for the commands start and end
	if !commands(input) {
		return
	}

	// We read and store the rooms and links
	rooms, roomNames, ok := colonyReader(input)
	if !ok {
		return
	}

	// We check if all links connect to existing rooms
	if !links(&rooms, roomNames) {
		return
	}

	// We check for duplicate links
	if !duplicates(&rooms) {
		return
	}

	// We print the rooms and links
	fmt.Println(rooms)
}
