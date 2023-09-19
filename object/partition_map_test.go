package object

import (
	"fmt"
	"maps"
)

// go 1.21
// https://medium.com/@goel.yash143/go-maps-magic-unveiling-the-maps-package-in-go-v1-21-c0bb4ebbf9e6
func TestMaps_shallow_copy(t *testing.T) {
	orig := map[string]int{"a": 1, "b": 2, "c": 3}
	clone := maps.Clone(orig)
	fmt.Println("Original Map:", originalMap)
	fmt.Println("Cloned Map:", clonedMap)
}
