package rex

import (
	"fmt"
	"strings"
)

func (l token_literal) stringify() string {
	return string(l.literal)
}

func (g token_group_raw) stringify() string {
	s := make([]string, len(g))
	for i, t := range g {
		s[i] = t.stringify()
	}
	return strings.Join(s, "")
}

func (g token_group) stringify() string {
	return "(" + g.tokens.stringify() + ")"
}

func (g token_group_uncaptured) stringify() string {
	return "<" + g.tokens.stringify() + ">"
}

func (o token_or) stringify() string {
	return fmt.Sprintf("(%s|%s)", o.left.stringify(), o.right.stringify())
}

func (r token_repeat) stringify() string {
	return fmt.Sprintf("(%s){%d,%d}", r.single.stringify(), r.min, r.max)
}
