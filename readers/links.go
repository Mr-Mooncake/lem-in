package readers

import (
	"fmt"
	"lem-in/structs"
	"strings"
)

func Links(rooms *structs.Rooms, roomNames map[string]bool) bool {
	if len(rooms.Links) == 0 {
		fmt.Println("ERROR: invalid data format, no links found")
		return false
	}

	// Check that every linked room exists
	for _, link := range rooms.Links {
		if !roomNames[link.From] || !roomNames[link.To] {
			fmt.Println("ERROR: invalid data format, unknown room in link")
			return false
		}
	}

	return true
}

func LinkReader(line string, rooms *structs.Rooms) bool {
	link := strings.Split(line, "-")

	if len(link) != 2 || link[0] == "" || link[1] == "" {
		fmt.Println("ERROR: invalid data format, invalid link")
		return false
	}

	if link[0] == link[1] {
		fmt.Println("ERROR: invalid data format, room cannot link to itself")
		return false
	}

	rooms.Links = append(rooms.Links, structs.Link{From: link[0], To: link[1]})
	return true
}
