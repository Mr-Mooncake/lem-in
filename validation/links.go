package validation

import (
	"fmt"
	"strings"
)

func links(rooms *Rooms, roomNames map[string]bool) bool {
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

func linkReader(line string, rooms *Rooms) bool {
	link := strings.Split(line, "-")

	if len(link) != 2 || link[0] == "" || link[1] == "" {
		fmt.Println("ERROR: invalid data format, invalid link")
		return false
	}

	if link[0] == link[1] {
		fmt.Println("ERROR: invalid data format, room cannot link to itself")
		return false
	}

	rooms.Links = append(rooms.Links, Link{From: link[0], To: link[1]})

	return true
}
