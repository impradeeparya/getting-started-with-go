package main

import (
	"fmt"
	"sort"
)

func main() {

	states := make(map[string]string)
	states["DL"] = "Delhi"
	states["MH"] = "Maharashtra"
	states["UK"] = "Uttarakhand"
	fmt.Println(states)

	delhi := states["DL"]
	fmt.Println(delhi)

	delete(states, "MH")
	fmt.Println(states)

	for k, v := range states {
		fmt.Println(k, v)
	}

	states["HR"] = "Haryana"
	states["KR"] = "Karnataka"
	states["TN"] = "Tamil nadu"
	keys := make([]string, len(states))
	index := 0
	for key := range states {
		keys[index] = key
		index++
	}
	sort.Strings(keys)
	fmt.Println(keys)

	fmt.Println(states)
	for key := range keys {
		fmt.Print(states[keys[key]], " ")
	}
}
