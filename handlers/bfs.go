package handlers

func Bfs(graph map[string][]string, start, end string) []string {
	queue := []string{start}
	visited := map[string]bool{start: true}
	previous := make(map[string]string)

	// Go through the queue
	for i := 0; i < len(queue); i++ {
		current := queue[i]

		// Stop when we find the end
		if current == end {
			break
		}

		// Check connected rooms
		for _, neighbor := range graph[current] {
			if visited[neighbor] {
				continue
			}

			visited[neighbor] = true
			previous[neighbor] = current
			queue = append(queue, neighbor)
		}
	}

	// No path found
	if !visited[end] {
		return nil
	}

	// Build the path backwards
	path := []string{}

	for room := end; room != start; room = previous[room] {
		path = append(path, room)
	}

	path = append(path, start)

	// Reverse the path
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path
}
