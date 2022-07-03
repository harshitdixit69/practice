package main

import (
	"math"
	"sort"
)

// func main() {
// 	start := time.Now()
// 	elapsed := time.Since(start)
// 	points := [][]int{{1, 3}, {2, 9}, {3, 7}, {3, 7}, {3, 7}}

// 	vector := []int{2, 2}
// 	arr := []int{}
// 	c := 4
// 	keys, finaalMap := newFunction(points, vector, arr, c)
// 	for _, k := range keys {
// 		if c != 0 {
// 			fmt.Println(finaalMap[k])
// 			c--
// 		}
// 	}
// 	log.Printf("Binomial took %s", elapsed)
// }

func newFunction(points [][]int, vector []int, arr []int, c int) ([]int, map[int][]int) {
	for i := 0; i < len(points); i++ {
		dis := math.Sqrt(float64(math.Pow(float64(vector[1]-vector[0]), 2)) + float64(math.Pow(float64(points[i][1]-points[i][0]), 2)))
		arr = append(arr, int(dis))
	}
	slcMap := make(map[int][]int)
	for i := 0; i < len(arr); i++ {
		slcMap[arr[i]] = points[i]
	}
	keys := make([]int, 0, len(slcMap))

	for k := range slcMap {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys, slcMap
}
