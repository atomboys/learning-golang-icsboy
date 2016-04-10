package main

import "fmt"
import "math/rand"
import "time"
import "sort"

func contains(slice []int, item int) bool {
	set := make(map[int]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

func main() {
	lotto := make([]int, 6)
	var new_number int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 6; i++ {
		new_number = rand.Intn(45)
		if !contains(lotto, new_number) {
			lotto[i] = new_number
		} else {
			i--
		}
	}
	sort.Ints(lotto)
	fmt.Print(lotto)
}
