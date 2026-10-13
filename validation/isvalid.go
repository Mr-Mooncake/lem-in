package validation

import (
	"fmt"
	output "lem-in/Ouput"
	"lem-in/handlers"
	"lem-in/readers"
	"strconv"
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

	// Mooncake: pick the best paths for all ants and print the result
	numAnts, _ := strconv.Atoi(input[0])
	paths := handlers.FindPaths(graph, rooms.Start.Name, rooms.End.Name, numAnts)
	output.PrintOut(input, paths, handlers.Distribute(paths, numAnts))
}