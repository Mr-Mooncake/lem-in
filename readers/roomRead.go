package readers

import (
	"fmt"
	"lem-in/structs"
	"strconv"
	"strings"
)

func RoomReader(line string, rooms *structs.Rooms, roomNames map[string]bool, nextRoom string) bool {
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

	room := structs.Room{Name: fields[0], X: x, Y: y}

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
