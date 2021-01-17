// Package drx allows for declaratively building regexps.
package drx

import (
	"regexp"
	"strings"
)

// Rx is the primary interface for building a regex string. Anything that emits
// a valid regex string can be used to declaratively build up a regex.
type Rx interface {
	// String returns a regex string for using with the various regexp package
	// compile functions.
	String() string
}

// Compile parses a declaratively built regular expression.
//
// See https://golang.org/pkg/regexp/#Compile for more details.
func Compile(r Rx) (*regexp.Regexp, error) {
	return regexp.Compile(r.String())
}

// MustCompile is like Compile but panics on error.
//
// See https://golang.org/pkg/regexp/#MustCompile for more details.
func MustCompile(r Rx) *regexp.Regexp {
	return regexp.MustCompile(r.String())
}

type builder struct {
	rxs []Rx
}

func (b builder) String() string {
	var sb strings.Builder
	for _, rx := range b.rxs {
		sb.WriteString(rx.String())
	}
	return sb.String()
}

// Build builds up a single Rx from multiple potentially nested Rxs.
func Build(rxs ...Rx) Rx {
	return builder{
		rxs: rxs,
	}
}
