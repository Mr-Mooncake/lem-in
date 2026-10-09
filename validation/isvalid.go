package validation

import (
	"fmt"
	"lem-in/handlers"
	"lem-in/readers"
)

func IsValid() {
	// We read the file
	input := readers.ParseFile()
	if input == nil {
		return
	}

	// We check the ants
	if !ants(input) {
		return
	}

	// We remove the comments #
	input = readers.CommentRemover(input)

	// We check for the commands start and end
	if !readers.Commands(input) {
		return
	}

	// We read and store the rooms and links
	rooms, roomNames, ok := readers.ColonyReader(input)
	if !ok {
		return
	}

	// We check if all links connect to existing rooms
	if !readers.Links(&rooms, roomNames) {
		return
	}

	// We check for duplicate links
	if !duplicates(&rooms) {
		return
	}
	graph := handlers.CreateGraph(&rooms)

	// Find a shortest path from start to end
	path := handlers.Bfs(graph, rooms.Start.Name, rooms.End.Name)

	if path == nil {
		fmt.Println("ERROR: invalid data format, no path found")
		return
	}

	// Print the path for testing
	fmt.Println("Shortest path:", path)
}
