package handlers

import (
	"lem-in/structs"
)

func CreateGraph(rooms *structs.Rooms) map[string][]string {
	graph := make(map[string][]string)

	// Add all rooms to the graph
	for _, room := range rooms.All {
		graph[room.Name] = []string{}
	}

	// Connect the rooms using their links
	for _, link := range rooms.Links {
		graph[link.From] = append(graph[link.From], link.To)
		graph[link.To] = append(graph[link.To], link.From)
	}
	return graph
}
