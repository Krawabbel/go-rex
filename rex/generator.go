package rex

func generate(entry *state, max_depth int) ([]string, bool) {

	if entry == nil {
		return nil, false
	}

	if max_depth < 0 {
		return nil, false
	}

	if entry.terminal {
		return []string{""}, true
	}

	ret_samples := make([]string, 0)
	ret_valid := false

	for char, nexts := range entry.tmap {

		symbol := string(char)
		if char == empty_char {
			symbol = ""
		}

		for _, next := range nexts {
			remainders, valid := generate(next, max_depth-1)
			if valid {
				for _, remainder := range remainders {
					ret_samples = append(ret_samples, symbol+remainder)
					ret_valid = true
				}
			}
		}
	}

	return ret_samples, ret_valid
}

type SortByLexicographicalOrder []string

func (a SortByLexicographicalOrder) Len() int           { return len(a) }
func (a SortByLexicographicalOrder) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByLexicographicalOrder) Less(i, j int) bool { return a[i] < a[j] }
