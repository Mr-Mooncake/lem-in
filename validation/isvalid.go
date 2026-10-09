package validation

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const usage = "Usage: go run . <filename.txt>"

func IsValid() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println(usage)
		return
	}
	// noMoreRooms := false
	filename := args[1]

	if !strings.HasSuffix(filename, ".txt") {
		fmt.Println(usage)
		return
	}
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(usage)
		return
	}
	content := strings.ReplaceAll(string(data), "\r\n", "\n")
	content = strings.TrimSuffix(content, "\n")

	input := strings.Split(content, "\n")

	// fmt.Println(input[0])

	ants, err := strconv.Atoi(input[0])
	if err != nil {
		fmt.Println("ERROR: invalid data format, invalid number of ants")
		return
	}
	if ants <= 0 {
		fmt.Println("ERROR: invalid data format, number of ants should be greater than or equal to 1")
		return
	}

	var cleaned []string

	for _, line := range input {
		if strings.HasPrefix(line, "#") &&
			line != "##start" && line != "##end" {
			continue
		}

		cleaned = append(cleaned, line)
	}

	input = cleaned

	foundStart := false
	foundEnd := false

	for _, v := range input {
		if v == "##start" {
			if foundStart {
				fmt.Println("ERROR: invalid data format, only 1 start command allowed")
				return
			}
			foundStart = true
		}

		if v == "##end" {
			if foundEnd {
				fmt.Println("ERROR: invalid data format, only 1 end command allowed")
				return
			}
			foundEnd = true
		}
	}
	if !foundStart || !foundEnd {
		fmt.Println("ERROR: invalid data format, there should a start and end command in the file")
		return
	}
	var rooms Rooms
	roomNames := make(map[string]bool)
	for i, v := range input {

		if i == 0 || strings.TrimSpace(v) == "" {
			continue
		}
		if v == "##start" || v == "##end" {
			if i+1 >= len(input) {
				fmt.Println("ERROR: invalid data format, missing room")
				return
			}

			fields := strings.Fields(input[i+1])
			if len(fields) != 3 {
				fmt.Println("ERROR: invalid data format, invalid room")
				return
			}
			if strings.Contains(fields[0], "L") || strings.Contains(fields[0], "#") {
				fmt.Println("ERROR: invalid data format, invalid room name")
				return
			}

			x, err := strconv.Atoi(fields[1])
			if err != nil {
				fmt.Println("ERROR: invalid data format, room coordinates must be integers")
				return
			}

			y, err := strconv.Atoi(fields[2])
			if err != nil {
				fmt.Println("ERROR: invalid data format, room coordinates must be integers")
				return
			}

			room := Room{Name: fields[0], X: x, Y: y}
			if input[i] == "##start" {
				rooms.Start = room
			} else {
				rooms.End = room
			}
			continue
		}

		fields := strings.Fields(v)

		if len(fields) == 1 {
			link := strings.Split(fields[0], "-")
			if len(link) != 2 || link[0] == "" || link[1] == "" || link[0] == link[1] {
				fmt.Println("ERROR: invalid data format, invalid link")
				return
			}

			rooms.Links = append(rooms.Links, Link{From: link[0], To: link[1]})
			// noMoreRooms = true
			continue
		}

		if len(fields) != 3 {
			fmt.Println("ERROR: invalid data format, invalid room")
			return
		}
		if strings.HasPrefix(fields[0], "L") || strings.HasPrefix(fields[0], "#") {
			fmt.Println("ERROR: invalid data format, invalid room name")
			return
		}

		x, err := strconv.Atoi(fields[1])
		if err != nil {
			fmt.Println("ERROR: invalid data format, room coordinates must be integers")
			return
		}

		y, err := strconv.Atoi(fields[2])
		if err != nil {
			fmt.Println("ERROR: invalid data format, room coordinates must be integers")
			return
		}

		room := Room{Name: fields[0], X: x, Y: y}
		if roomNames[room.Name] {
			fmt.Println("ERROR: invalid data format, duplicate room")
			return
		}

		roomNames[room.Name] = true
		rooms.All = append(rooms.All, room)
	}

	for _, link := range rooms.Links {
		if !roomNames[link.From] || !roomNames[link.To] {
			fmt.Println("ERROR: invalid data format, unknown room in link")
			return
		}
	}
	if len(rooms.Links) == 0 {
		fmt.Println("ERROR: invalid data format, no links found")
		return
	}
	linkNames := make(map[[2]string]bool)

	for _, link := range rooms.Links {
		key := [2]string{link.From, link.To}

		if key[0] > key[1] {
			key[0], key[1] = key[1], key[0]
		}

		if linkNames[key] {
			fmt.Println("ERROR: invalid data format, duplicate link")
			return
		}

		linkNames[key] = true
	}

	// fmt.Println(rooms)
}
