package rex

const empty_char = 0x00

type token interface {
	stringify() string
	compile(*state) *state
}

// TOKEN LITERAL

type token_literal struct {
	literal byte
}

func (l token_literal) compile(entry *state) *state {
	exit := &state{id: entry.id + 1}

	connect(entry, l.literal, exit)

	return exit
}

// TOKEN GROUP

type token_group_raw []token

func (g token_group_raw) compile(entry *state) *state {

	last := entry
	for _, t := range g {
		last = t.compile(last)
	}

	exit := new(state)

	connect(last, empty_char, exit)

	return exit
}

type token_group struct {
	tokens token_group_raw
}

func (g token_group) compile(entry *state) *state {
	return g.tokens.compile(entry)
}

type token_group_uncaptured struct {
	tokens token_group_raw
}

func (g token_group_uncaptured) compile(entry *state) *state {
	return g.tokens.compile(entry)
}

// TOKEN OR

type token_or struct {
	left, right token_group_uncaptured
}

func (o token_or) compile(entry *state) *state {
	exit_left := o.left.compile(entry)
	exit_right := o.right.compile(entry)

	connect(exit_right, empty_char, exit_left)

	return exit_left
}

// TOKEN REPEAT

type token_repeat struct {
	single   token
	min, max int
}

func (r token_repeat) compile(entry *state) *state {
	exit := entry
	for i := 0; i < r.min; i++ {
		exit = r.single.compile(exit)
	}

	switch r.max {
	case repeat_infinity:
		exit_intermediate := r.single.compile(exit)

		connect(exit_intermediate, empty_char, exit)

	default:
		n_optional := r.max - r.min

		last := exit
		exit = new(state)
		connect(last, empty_char, exit)

		for i := 0; i < n_optional; i++ {
			last = r.single.compile(last)
			connect(last, empty_char, exit)
		}

	}
	return exit
}
