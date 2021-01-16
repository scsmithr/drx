// Package drx allows for declaratively building regexps.
package drx

import (
	"regexp"
	"strings"
)

type Rx interface {
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

func Build(rxs ...Rx) Rx {
	return builder{
		rxs: rxs,
	}
}
