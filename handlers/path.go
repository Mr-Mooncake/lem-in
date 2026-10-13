package handlers

import (
	"sort"
)

// Mooncake: finds the set of room-disjoint paths that gets all ants to ##end in the fewest turns.

// add puts an edge and its reverse (used to undo flow) in the network
func add(n *FlowNet, from, to, cap, cost int) {
	n.Adj[from] = append(n.Adj[from], len(n.Edges))
	n.Edges = append(n.Edges, Edge{To: to, Cap: cap, Cost: cost})
	n.Adj[to] = append(n.Adj[to], len(n.Edges))
	n.Edges = append(n.Edges, Edge{To: from, Cap: 0, Cost: -cost})
}

// augment finds the cheapest route from src to sink (queue based search,
// like Bfs but it can reuse and undo earlier paths) and pushes one path through it
func augment(n *FlowNet, src, sink int) bool {
	const inf = 1 << 30
	dist := make([]int, len(n.Adj))
	prev := make([]int, len(n.Adj))
	inQueue := make([]bool, len(n.Adj))
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

		for _, id := range n.Adj[u] {
			e := n.Edges[id]
			if e.Cap > 0 && dist[u]+e.Cost < dist[e.To] {
				dist[e.To] = dist[u] + e.Cost
				prev[e.To] = id
				if !inQueue[e.To] {
					inQueue[e.To] = true
					queue = append(queue, e.To)
				}
			}
		}
	}

	if dist[sink] == inf {
		return false
	}

	for v := sink; v != src; v = n.Edges[prev[v]^1].To {
		n.Edges[prev[v]].Cap--
		n.Edges[prev[v]^1].Cap++
	}
	return true
}

// extract reads the current flow back as room name paths (start -> end)
func extract(n *FlowNet, src, sink int, names []string) [][]string {
	var paths [][]string

	for _, id := range n.Adj[src] {
		if id%2 != 0 || n.Edges[id].Cap != 0 {
			continue
		}

		path := []string{names[src/2]}
		cur := n.Edges[id].To

		for {
			path = append(path, names[cur/2])
			if cur == sink {
				break
			}

			next := -1
			for _, oid := range n.Adj[cur+1] {
				if oid%2 == 0 && n.Edges[oid].Cap == 0 {
					next = n.Edges[oid].To
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
	net := &FlowNet{Adj: make([][]int, 2*len(names))}
	for i, name := range names {
		if name != start && name != end {
			add(net, 2*i, 2*i+1, 1, 0)
		}
		for _, nb := range graph[name] {
			add(net, 2*i+1, 2*index[nb], 1, 1)
		}
	}

	src, sink := 2*index[start]+1, 2*index[end]

	var best [][]string
	bestTurns := -1

	for augment(net, src, sink) {
		paths := extract(net, src, sink, names)
		sort.SliceStable(paths, func(a, b int) bool { return len(paths[a]) < len(paths[b]) })

		if t := turnsNeeded(paths, ants); bestTurns == -1 || t < bestTurns {
			best, bestTurns = paths, t
		}
	}
	return best
}
