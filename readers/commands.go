package readers

import (
	"fmt"
	"strings"
)

func CommentRemover(input []string) []string {
	var cleaned []string

	for _, line := range input {
		if strings.HasPrefix(line, "##") && line != "##start" && line != "##end" {
			fmt.Println("ERROR: invalid data format, other attemps of commands that are not start or end are unallowed")
			return nil
		}
		if strings.HasPrefix(line, "#") && line != "##start" && line != "##end" {
			continue
		}

		cleaned = append(cleaned, line)
	}

	input = cleaned
	return cleaned
}

func Commands(input []string) bool {
	foundStart := false
	foundEnd := false

	for _, v := range input {

		if v == "##start" {
			if foundStart {
				fmt.Println("ERROR: invalid data format, only 1 start command allowed")
				return false
			}
			foundStart = true
		}

		if v == "##end" {
			if foundEnd {
				fmt.Println("ERROR: invalid data format, only 1 end command allowed")
				return false
			}
			foundEnd = true
		}
	}

	if !foundStart || !foundEnd {
		fmt.Println("ERROR: invalid data format, there should be a start and end command in the file")
		return false
	}
	return true
}
