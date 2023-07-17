package main

import (
	"fmt"
	"sort"
)

func main() {
	states := make(map[string]string) //Make a map called states
	fmt.Println(states)               // Should be empty as there is nothing in the map
	// Add keys and values to the states map
	states["AL"] = "Alabama"
	states["AK"] = "Alaska"
	states["NC"] = "North Carolina"
	states["VT"] = "Vermont"
	states["WA"] = "Washington"
	fmt.Println(states) // Prints contents of map states
	Alabama := states["AL"]
	fmt.Println(Alabama)

	delete(states, "AL") // To delete a value from a map
	fmt.Println(states)
	states["NY"] = "New York" // To add a value to the map
	fmt.Println(states)

	for k, v := range states { // A loop to iterate through states. For keys (k) and values (v) in states (range states)
		fmt.Printf("%v: %v\n", k, v) //Print %v the value of k and then %v the value of v.
	}
	keys := make([]string, len(states)) // Ceating a map called keys used for indexing and sorting
	i := 0                              //Setting a counter to 0
	for k := range states {             // A loop to iterate through the keys
		keys[i] = k // Set keys index 0 to the first element of map
		i++         // Iterate through the states
	}
	fmt.Println(keys)
	sort.Strings(keys)
	fmt.Println(keys)
	for i := range keys {
		fmt.Println(states[keys[i]])
	}
}
