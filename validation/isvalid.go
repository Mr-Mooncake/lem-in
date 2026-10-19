package validation

import (
	"fmt"
	"lem-in/handlers"
	"lem-in/readers"
	"strconv"
)

func IsValid() ([]string, handlers.Rooms, bool) {
	// We read the file
	input := readers.ParseFile()
	if input == nil {
		return nil, handlers.Rooms{}, false
	}

	// We check the ants
	if !ants(input) {
		return nil, handlers.Rooms{}, false
	}

	// We remove the comments #
	input = readers.CommentRemover(input)
	if input == nil {
		return nil, handlers.Rooms{}, false
	}

	// We check for the commands start and end
	if !readers.Commands(input) {
		return nil, handlers.Rooms{}, false
	}

	// We read and store the rooms and links
	rooms, roomNames, ok := readers.ColonyReader(input)
	if !ok {
		return nil, handlers.Rooms{}, false
	}

	// We check if all links connect to existing rooms
	if !readers.Links(&rooms, roomNames) {
		return nil, handlers.Rooms{}, false
	}

	// We check for duplicate links
	if !duplicates(&rooms) {
		return nil, handlers.Rooms{}, false
	}
	return input, rooms, true
}

func duplicates(rooms *handlers.Rooms) bool {
	linkNames := make(map[[2]string]bool)

	for _, link := range rooms.Links {
		key := [2]string{link.From, link.To}

		if key[0] > key[1] {
			key[0], key[1] = key[1], key[0]
		}

		if linkNames[key] {
			fmt.Println("ERROR: invalid data format, duplicate link")
			return false
		}

		linkNames[key] = true
	}

	return true
}

func ants(input []string) bool {
	ants, err := strconv.Atoi(input[0])
	if err != nil {
		fmt.Println("ERROR: invalid data format, invalid number of ants")
		return false
	}

	if ants <= 0 {
		fmt.Println("ERROR: invalid data format, number of ants should be greater than or equal to 1")
		return false
	}
	return true
}
