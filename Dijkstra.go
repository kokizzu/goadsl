package main

import "fmt"
import "sort"

// O(n)
func MinCost(cost []int, dist []int) (int, int) {
	minn := 100000000
	to := 100000000
	for j:=0; j<len(cost); j++ {
		if cost[j] != 0 && dist[j] != 1 {
			if cost[j] < minn {
				minn = cost[j]
				to = j
			} 
		}
	}
	return minn, to
}

// O(n^2)
func Dijkstra(graph [][]int, vertex int) (map[string]int, int) {
	min := 0
	path := make(map[string]int)
	count := 0
	next := 0
	path_temp := ""
	dist := make([]int, vertex)
	for len(path) < vertex-1 {
		dist[next] = 1
		a, b := MinCost(graph[next], dist) 
		path_temp = string(count+'1')+". "+string(next+'A')+"->"+string(b+'A')
		path[path_temp] = a
		min += a
		next = b
		count += 1
	}
	return path, min
}

func main(){
	vertex := 5
	/*
	      2    3
       A----B----C
       |   / \   |
      6| 8/   \5 |7
       | /     \ |
	   D---------E
            9        
	*/
	graph := [][]int{
					{0, 2, 0, 6, 0},
                    {2, 0, 3, 8, 5},
                    {0, 3, 0, 0, 7},
                    {6, 8, 0, 0, 9},
                    {0, 5, 7, 9, 0},
			}
	path, min := Dijkstra(graph, vertex)
	fmt.Println("Shortest path : ")
	var keys []string
    for k := range path {
        keys = append(keys, k)
    }
    sort.Strings(keys)
	for _, k := range keys {
        fmt.Println(k, ":", path[k])
    }
	fmt.Println("Total cost :", min)
}
