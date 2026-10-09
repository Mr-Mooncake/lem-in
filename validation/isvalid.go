package validation

import "fmt"

const usage = "Usage: go run . <filename.txt>"

func IsValid() {
	//we read the file
	input := parseFile()
	if input == nil {
		return
	}
	//we check the ants
	if !ants(input) {
		return
	}
	//we remove the comments #
	input = commentRemover(input)

	//we check for the commands start and end
	if !commands(input) {
		return
	}

	rooms, roomNames, ok := colonyReader(input)
	if !ok {
		return
	}

	if !links(&rooms, roomNames) {
		return
	}

	if !duplicates(&rooms) {
		return
	}

	fmt.Println(rooms)
}
