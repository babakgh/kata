package grokking

import "golang.org/x/exp/constraints"

func Dijkstra[T comparable, K constraints.Ordered](graph map[T]map[T]K, from, to T) []T {
	return []T{}
}

// func Dijkstra[T comparable, K constraints.Ordered](graph map[T][]Edge[T, K], from, to T) []T {
// 	// var nullT T
// 	var nullK K
// 	//

// 	table := make(map[T]K, len(graph))
// 	table[to] = nullK

// 	var closest Edge[T, K]
// 	for _, edge := range graph[from] {
// 		table[edge.Vertex] = edge.Weight
// 		if edge.Weight < closest.Weight {
// 			closest = edge
// 		}
// 	}

// 	for _, edge := range graph[closest.Vertex] {
// 		if currentWeight, ok := table[edge.Vertex]; ok {
// 			if currentWeight > edge.Weight+closest.Weight {
// 				table[edge.Vertex] = edge.Weight + closest.Weight
// 			}
// 		} else {
// 			table[edge.Vertex] = edge.Weight + closest.Weight
// 		}
// 	}

// 	return []T{}
// }
