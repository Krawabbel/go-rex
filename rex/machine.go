package rex

import "fmt"

func prepare(pattern string) (*state, error) {
	tokens, err := parse(pattern)
	if err != nil {
		return nil, fmt.Errorf("compiler error: %v", err)
	}

	entry := &state{start: true}

	last := tokens.compile(entry)

	exit := &state{terminal: true}

	connect(last, empty_char, exit)

	return entry, nil
}
