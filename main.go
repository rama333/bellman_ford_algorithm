package main

import "fmt"

type Graph struct {
	a    int
	b    int
	edge int
}

func main() {

	inf := 1 << 31
	v := 0

	graphs := []Graph{}

	graphs = append(graphs, Graph{a: 0, b: 1, edge: 7})
	graphs = append(graphs, Graph{a: 1, b: 2, edge: -5})
	graphs = append(graphs, Graph{a: 0, b: 2, edge: 6})
	graphs = append(graphs, Graph{a: 2, b: 3, edge: 4})
	graphs = append(graphs, Graph{a: 3, b: 1, edge: 8})

	d := []int{}
	m := len(graphs)
	n := len(graphs)

	for i := 0; i < n-1; i++ {
		d = append(d, inf)
	}

	d[v] = 0

	for i := 0; i < n-1; i++ {
		for j := 0; j < m; j++ {
			if d[graphs[j].a] < inf {
				d[graphs[j].b] = min(d[graphs[j].b], d[graphs[j].a]+graphs[j].edge)
				fmt.Println(i, j, d[graphs[j].b], d[graphs[j].a]+graphs[j].edge)
			}

		}
	}

	fmt.Println(d) //return [0 7 2 6]

}

func min(a int, b int) int {

	if a < b {
		return a
	}
	return b

}
