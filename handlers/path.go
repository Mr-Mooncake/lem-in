package handlers

import "sort"

// Mooncake: finds the set of room-disjoint paths that gets all ants to ##end in the fewest turns.

type edge struct{ to, cap, cost int }

type flowNet struct {
	edges []edge
	adj   [][]int
}

// add puts an edge and its reverse (used to undo flow) in the network
func (n *flowNet) add(from, to, cap, cost int) {
	n.adj[from] = append(n.adj[from], len(n.edges))
	n.edges = append(n.edges, edge{to, cap, cost})
	n.adj[to] = append(n.adj[to], len(n.edges))
	n.edges = append(n.edges, edge{from, 0, -cost})
}

// augment finds the cheapest route from src to sink (queue based search,
// like Bfs but it can reuse and undo earlier paths) and pushes one path through it
func (n *flowNet) augment(src, sink int) bool {
	const inf = 1 << 30
	dist := make([]int, len(n.adj))
	prev := make([]int, len(n.adj))
	inQueue := make([]bool, len(n.adj))
	for i := range dist {
		dist[i] = inf
		prev[i] = -1
	}
	dist[src] = 0
	queue := []int{src}

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		inQueue[u] = false

		for _, id := range n.adj[u] {
			e := n.edges[id]
			if e.cap > 0 && dist[u]+e.cost < dist[e.to] {
				dist[e.to] = dist[u] + e.cost
				prev[e.to] = id
				if !inQueue[e.to] {
					inQueue[e.to] = true
					queue = append(queue, e.to)
				}
			}
		}
	}

	if dist[sink] == inf {
		return false
	}

	for v := sink; v != src; v = n.edges[prev[v]^1].to {
		n.edges[prev[v]].cap--
		n.edges[prev[v]^1].cap++
	}
	return true
}

// extract reads the current flow back as room name paths (start -> end)
func (n *flowNet) extract(src, sink int, names []string) [][]string {
	var paths [][]string

	for _, id := range n.adj[src] {
		if id%2 != 0 || n.edges[id].cap != 0 {
			continue
		}

		path := []string{names[src/2]}
		cur := n.edges[id].to

		for {
			path = append(path, names[cur/2])
			if cur == sink {
				break
			}

			next := -1
			for _, oid := range n.adj[cur+1] {
				if oid%2 == 0 && n.edges[oid].cap == 0 {
					next = n.edges[oid].to
					break
				}
			}
			if next == -1 {
				break
			}
			cur = next
		}
		paths = append(paths, path)
	}
	return paths
}

// Distribute gives each ant (in order) the path where it would arrive first
func Distribute(paths [][]string, ants int) []int {
	load := make([]int, len(paths))
	assign := make([]int, ants)

	for a := range assign {
		best := 0
		for i := range paths {
			if len(paths[i])+load[i] < len(paths[best])+load[best] {
				best = i
			}
		}
		assign[a] = best
		load[best]++
	}
	return assign
}

// turnsNeeded is how many turns the last ant needs with these paths
func turnsNeeded(paths [][]string, ants int) int {
	load := make([]int, len(paths))
	for _, p := range Distribute(paths, ants) {
		load[p]++
	}

	turns := 0
	for i, l := range load {
		if l > 0 && len(paths[i])-1+l-1 > turns {
			turns = len(paths[i]) - 1 + l - 1
		}
	}
	return turns
}

// FindPaths adds one path at a time (each time the cheapest total) and keeps
// the amount of paths that gives the fewest turns for the given ants
func FindPaths(graph map[string][]string, start, end string, ants int) [][]string {
	names := make([]string, 0, len(graph))
	for name := range graph {
		names = append(names, name)
	}
	sort.Strings(names)

	index := make(map[string]int)
	for i, name := range names {
		index[name] = i
	}

	// every room is split in (2i) and out (2i+1) with capacity 1 so rooms are not shared
	net := &flowNet{adj: make([][]int, 2*len(names))}
	for i, name := range names {
		if name != start && name != end {
			net.add(2*i, 2*i+1, 1, 0)
		}
		for _, nb := range graph[name] {
			net.add(2*i+1, 2*index[nb], 1, 1)
		}
	}

	src, sink := 2*index[start]+1, 2*index[end]

	var best [][]string
	bestTurns := -1

	for net.augment(src, sink) {
		paths := net.extract(src, sink, names)
		sort.SliceStable(paths, func(a, b int) bool { return len(paths[a]) < len(paths[b]) })

		if t := turnsNeeded(paths, ants); bestTurns == -1 || t < bestTurns {
			best, bestTurns = paths, t
		}
	}
	return best
}
