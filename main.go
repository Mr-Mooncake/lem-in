package main

import (
	"fmt"
	"lem-in/handlers"
	"lem-in/validation"
	"strconv"
)

func main() {
	input, rooms, ok := validation.IsValid()
	if !ok {
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
	handlers.PrintOut(input, paths, handlers.Distribute(paths, numAnts))
}
