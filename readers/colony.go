package readers

import (
	"fmt"
	"strings"
	"lem-in/handlers"
)

func ColonyReader(input []string) (handlers.Rooms, map[string]bool, bool) {
	var rooms handlers.Rooms
	roomNames := make(map[string]bool)

	noMoreRooms := false
	nextRoom := ""

	for i := 1; i < len(input); i++ {
		v := input[i]

		if strings.TrimSpace(v) == "" {
			fmt.Println("ERROR: invalid data format, empty line")
			return rooms, roomNames, false
		}

		// Detect the start and end commands
		if v == "##start" || v == "##end" {
			if noMoreRooms || nextRoom != "" {
				fmt.Println("ERROR: invalid data format, invalid command position")
				return rooms, roomNames, false
			}

			nextRoom = v
			continue
		}

		fields := strings.Fields(v)

		// The room after ##start or ##end
		if nextRoom != "" {
			if !RoomReader(v, &rooms, roomNames, nextRoom) {
				return rooms, roomNames, false
			}

			nextRoom = ""
			continue
		}

		// Normal rooms
		if len(fields) == 3 {
			if noMoreRooms {
				fmt.Println("ERROR: invalid data format, room after links")
				return rooms, roomNames, false
			}

			if !RoomReader(v, &rooms, roomNames, "") {
				return rooms, roomNames, false
			}

			continue
		}

		// Links
		if len(fields) == 1 {
			if !LinkReader(v, &rooms) {
				return rooms, roomNames, false
			}

			noMoreRooms = true
			continue
		}

		fmt.Println("ERROR: invalid data format, invalid room or link")
		return rooms, roomNames, false
	}

	// Check if a command has no room after it
	if nextRoom != "" {
		fmt.Println("ERROR: invalid data format, missing room")
		return rooms, roomNames, false
	}

	return rooms, roomNames, true
}
