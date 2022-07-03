package main

import (
	"fmt"
	"log"
	"math"
	"sort"
	"sync"
	"time"
)

var arr []float64

func main() {
	var wg sync.WaitGroup
	start := time.Now()
	elapsed := time.Since(start)

	p4 := point{
		0, 2,
	}
	p5 := point{
		2, 0,
	}

	p6 := point{
		1, 0,
	}

	p7 := point{
		3, 0,
	}

	points := []point{p4, p5, p6, p7}
	vector := []int{0, 0}
	c := 3
	ch := make(chan float64, 10)

	for _, v := range points {
		wg.Add(1)
		go findDis(v, vector, c, ch, &wg)
		wg.Wait()
		// val := <-ch
	}
	close(ch)
	for v := range ch {
		arr = append(arr, v)
	}

	slcMap := make(map[float64][]point)
	for i := 0; i < len(arr); i++ {
		if b, ok := slcMap[arr[i]]; ok {
			b = append(b, points[i])
			slcMap[arr[i]] = b
		} else {
			slcMap[arr[i]] = []point{points[i]}
		}
	}
	keys := make([]float64, 0, len(slcMap))

	for k := range slcMap {
		keys = append(keys, k)
	}
	sort.Float64s(keys)
	for _, k := range keys {
		if c != 0 {
			fmt.Println(k, slcMap[k])
			c--
		}
	}
	log.Printf("Binomial took %s", elapsed)
}
func findDis(v point, vector []int, c int, ch chan float64, wg *sync.WaitGroup) {
	dis := math.Sqrt(float64(math.Pow(float64(vector[1]-vector[0]), 2)) + float64(math.Pow(float64(v.y-v.x), 2)))
	ch <- dis
	wg.Done()
}

type point struct {
	x int
	y int
}
