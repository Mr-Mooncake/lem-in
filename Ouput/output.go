package ouput

import (
	"fmt"
	"strings"
)

// Mooncake: prints the colony file, a blank line, then the ant moves turn by turn.
// paths are the chosen start->end room paths, antPath[i] is the path index of ant i+1.
func PrintOut(input []string, paths [][]string, antPath []int) {
	for _, line := range input {
		fmt.Println(line)
	}
	fmt.Println()

	pos := make([]int, len(antPath)) // current index of each ant on its path
	occupied := make(map[string]bool)
	finished := 0

	for finished < len(antPath) {
		var moves []string

		// Ants are handled in order so the front ant on a path moves first
		for a := range antPath {
			p := paths[antPath[a]]
			if pos[a] == len(p)-1 {
				continue
			}

			next := p[pos[a]+1]
			isEnd := pos[a]+1 == len(p)-1

			// A room (other than the end) can hold one ant only
			if !isEnd && occupied[next] {
				continue
			}

			if pos[a] > 0 {
				occupied[p[pos[a]]] = false
			}
			pos[a]++

			if isEnd {
				finished++
			} else {
				occupied[next] = true
			}
			moves = append(moves, fmt.Sprintf("L%d-%s", a+1, next))
		}

		fmt.Println(strings.Join(moves, " "))
	}
}
