package validation

import (
	"fmt"
	"os"
	"strings"
)

func parseFile() []string {
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
