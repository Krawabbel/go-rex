package rex

import "fmt"

const repeat_infinity = -1

type parse_context struct {
	pattern string
	pos     int
	tokens  []token
}

func parse(pattern string) (token_group_raw, error) {

	ctx := &parse_context{
		pattern: pattern,
		pos:     0,
		tokens:  make([]token, 0),
	}

	for ctx.pos < len(pattern) {
		if err := ctx.parse_next(); err != nil {
			return nil, err
		}
		ctx.pos++
	}

	return ctx.tokens, nil
}

func (ctx *parse_context) parse_next() error {
	ch := ctx.pattern[ctx.pos]
	switch ch {
	case '(':
		return ctx.parse_group()
	case '|':
		return ctx.parse_or()
	case '*', '+', '?':
		return ctx.parse_repeat(ch)
	case '\\':
		return ctx.parse_escape()
	case ')':
		return fmt.Errorf("parse error: unexpected ')'")
	default:
		return ctx.parse_literal(ch)
	}
}

func (ctx *parse_context) parse_group() error {

	group_ctx := &parse_context{
		pattern: ctx.pattern,
		pos:     ctx.pos + 1,
		tokens:  make([]token, 0),
	}

	for group_ctx.pattern[group_ctx.pos] != ')' {
		if err := group_ctx.parse_next(); err != nil {
			return err
		}
		group_ctx.pos++
		if group_ctx.pos == len(group_ctx.pattern) {
			return fmt.Errorf("non-matched '('")
		}
	}

	ctx.tokens = append(ctx.tokens, token_group{tokens: group_ctx.tokens})
	ctx.pos = group_ctx.pos

	return nil
}

func (ctx *parse_context) parse_or() error {

	rhs_ctx := &parse_context{
		pattern: ctx.pattern,
		pos:     ctx.pos + 1,
		tokens:  make([]token, 0),
	}

	for rhs_ctx.pos < len(rhs_ctx.pattern) && rhs_ctx.pattern[rhs_ctx.pos] != ')' {
		if err := rhs_ctx.parse_next(); err != nil {
			return err
		}
		rhs_ctx.pos++
	}

	left := token_group_uncaptured{tokens: ctx.tokens}

	right := token_group_uncaptured{tokens: rhs_ctx.tokens}

	ctx.pos = rhs_ctx.pos - 1
	ctx.tokens = []token{token_or{left: left, right: right}}

	return nil
}

func (ctx *parse_context) parse_repeat(ch byte) error {

	var min, max int

	switch ch {
	case '*':
		min = 0
		max = repeat_infinity
	case '?':
		min = 0
		max = 1
	case '+':
		min = 1
		max = repeat_infinity
	}

	if len(ctx.tokens) == 0 {
		return fmt.Errorf("missing argument to repetition operator: '%s'", string(ch))
	}
	last_token := ctx.tokens[len(ctx.tokens)-1]
	ctx.tokens[len(ctx.tokens)-1] = token_repeat{single: last_token, min: min, max: max}

	return nil
}

func (ctx *parse_context) parse_literal(ch byte) error {
	ctx.tokens = append(ctx.tokens, token_literal{literal: ch})
	return nil
}

func (ctx *parse_context) parse_escape() error {
	ctx.pos++
	ch := ctx.pattern[ctx.pos]
	switch ch {
	case '*', '?', '+', '\\', '(', ')':
		ctx.parse_literal(ch)
	default:
		return fmt.Errorf("unexpected escape character '\\%s'", string(ctx.pattern[ctx.pos]))
	}
	return nil
}
