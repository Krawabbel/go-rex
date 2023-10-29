package rex

import (
	"fmt"
	"sort"
)

const (
	max_generation_depth = 1000000
)

func Generate(pattern string) ([]string, error) {
	entry, err := prepare(pattern)
	if err != nil {
		return nil, fmt.Errorf("generator error: %v", err)
	}

	samples, valid := generate(entry, max_generation_depth)
	if !valid {
		return nil, fmt.Errorf("generator error: no valid expressions found")
	}

	sort.Sort(SortByLexicographicalOrder(samples))

	return samples, nil
}

func Match(pattern string, msg []byte) (bool, error) {
	entry, err := prepare(pattern)

	if err != nil {
		return false, fmt.Errorf("match error: %v", err)
	}

	return match(entry, msg), nil
}
