package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	// Initilization seed
	rand.Seed(time.Now().UnixNano())

	var scoreMap = make(map[string]int, 100)

	// Fill data on map
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100)
		scoreMap[key] = value
	}

	// Get all keys
	var keys = make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}

	// Sort slice
	sort.Strings(keys)
	// Traverse map by sort key
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
}
