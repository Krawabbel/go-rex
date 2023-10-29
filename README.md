# go-rex
a regular expression matcher written in Go (Golang)

[![Go](https://github.com/Krawabbel/go-rex/actions/workflows/go.yml/badge.svg)](https://github.com/Krawabbel/go-rex/actions/workflows/go.yml)

[![Go Coverage](https://github.com/Krawabbel/go-rex/wiki/coverage.svg)](https://raw.githack.com/wiki/Krawabbel/go-rex/coverage.html)

The current version is a non-optimized default approach, i.e. it relies heavily on recursion. In particular, it does not (yet) implement the much faster Thompson NFA approach.

## Resources 
* https://rhaeguard.github.io/posts/regex/
* https://swtch.com/~rsc/regexp/regexp1.html
