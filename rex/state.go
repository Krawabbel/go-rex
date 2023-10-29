package rex

type state struct {
	start, terminal bool
	tmap            map[byte][]*state
	id              int
}

func connect(entry *state, char byte, exit *state) {

	if entry == nil || exit == nil {
		panic("cannot connect nil-state(s)")
	}

	if entry.tmap == nil {
		entry.tmap = make(map[byte][]*state)
	}

	states, found := entry.tmap[char]
	if !found || states == nil {
		states = make([]*state, 0)
	}

	entry.tmap[char] = append(states, exit)

}
