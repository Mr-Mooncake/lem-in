package handlers

func Bfs(graph map[string][]string, start, end string) []string {
	// Rooms to vist
	queue := []string{start}

	// Rooms we already visited
	visited := make(map[string]bool)
	visited[start] = true

	// Remember where we came from
	previous := make(map[string]string)

	for len(queue) > 0 {
		// Take the first room from the queue
		current := queue[0]
		queue = queue[1:]

		// Stop when we reach the end
		if current == end {
			break
		}

		// Check all rooms connected to the current room
		for _, neighbor := range graph[current] {
			if visited[neighbor] {
				continue
			}

			visited[neighbor] = true
			previous[neighbor] = current

			queue = append(queue, neighbor)
		}
	}

	// No path exists
	if !visited[end] {
		return nil
	}

	// go back
	var path []string

	for current := end; current != start; current = previous[current] {
		path = append(path, current)
	}

	path = append(path, start)

	//we reverse the order from end to start so it becomes start to end
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path
}
