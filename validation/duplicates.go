package validation

import "fmt"

func duplicates(rooms *Rooms) bool {
	linkNames := make(map[[2]string]bool)

	for _, link := range rooms.Links {
		key := [2]string{link.From, link.To}

		if key[0] > key[1] {
			key[0], key[1] = key[1], key[0]
		}

		if linkNames[key] {
			fmt.Println("ERROR: invalid data format, duplicate link")
			return false
		}

		linkNames[key] = true
	}

	return true
}
