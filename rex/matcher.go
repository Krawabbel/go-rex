package rex

func match(entry *state, msg []byte) bool {

	if entry.terminal {
		return true // terminal state
	}

	for char, nexts := range entry.tmap {
		switch char {
		case empty_char:
			for _, next := range nexts {
				if match(next, msg) {
					return true // match
				}
			}

		default:
			if len(msg) > 0 && char == msg[0] {
				for _, next := range nexts {
					if match(next, msg[1:]) {
						return true // match
					}
				}

			}
		}
	}
	return false // no match
}
