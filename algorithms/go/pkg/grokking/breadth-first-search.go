package grokking

func BFS[T comparable](graph map[T][]T, from, to T) []T {
	var null T

	if _, ok := graph[from]; !ok {
		return []T{}
	}

	queue := append([]T{from}, graph[from]...)

	visits := make(map[T]T, len(graph))
	visits[from] = null

	for len(queue) > 0 {
		var current T
		current, queue = queue[0], queue[1:]

		if current == to {
			path := []T{}
			for prev := to; prev != null; prev = visits[prev] {
				path = append([]T{prev}, path...)
			}
			return path
		}

		for _, neighbor := range graph[current] {
			if _, visited := visits[neighbor]; !visited {
				visits[neighbor] = current
				queue = append(queue, neighbor)
			}
		}
	}

	return []T{}
}
