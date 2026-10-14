package readers

import (
	"fmt"
	"lem-in/handlers"
	"os"
	"strconv"
	"strings"
)

const usage = "Usage: go run . <filename.txt>"

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

func ParseFile() []string {
	args := os.Args
	if len(args) != 2 {
		fmt.Println(usage)
		return nil
	}

	filename := args[1]

	if !strings.HasSuffix(filename, ".txt") {
		fmt.Println(usage)
		return nil
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(usage)
		return nil
	}

	content := strings.ReplaceAll(string(data), "\r\n", "\n")
	content = strings.TrimSuffix(content, "\n")

	input := strings.Split(content, "\n")
	return input
}

func LinkReader(line string, rooms *handlers.Rooms) bool {
	link := strings.Split(line, "-")

	if len(link) != 2 || link[0] == "" || link[1] == "" {
		fmt.Println("ERROR: invalid data format, invalid link")
		return false
	}

	if link[0] == link[1] {
		fmt.Println("ERROR: invalid data format, room cannot link to itself")
		return false
	}

	rooms.Links = append(rooms.Links, handlers.Link{From: link[0], To: link[1]})
	return true
}

func RoomReader(line string, rooms *handlers.Rooms, roomNames map[string]bool, nextRoom string) bool {
	fields := strings.Fields(line)

	if len(fields) != 3 {
		fmt.Println("ERROR: invalid data format, invalid room")
		return false
	}

	if strings.HasPrefix(fields[0], "L") || strings.HasPrefix(fields[0], "#") {
		fmt.Println("ERROR: invalid data format, invalid room name")
		return false
	}

	x, err := strconv.Atoi(fields[1])
	if err != nil {
		fmt.Println("ERROR: invalid data format, room coordinates must be integers")
		return false
	}

	y, err := strconv.Atoi(fields[2])
	if err != nil {
		fmt.Println("ERROR: invalid data format, room coordinates must be integers")
		return false
	}

	room := handlers.RoomDetails{Name: fields[0], X: x, Y: y}

	if roomNames[room.Name] {
		fmt.Println("ERROR: invalid data format, duplicate room")
		return false
	}

	roomNames[room.Name] = true
	rooms.All = append(rooms.All, room)

	switch nextRoom {
	case "##start":
		rooms.Start = room
	case "##end":
		rooms.End = room
	}

	return true
}
