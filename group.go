package drx

import "strings"

type groupedForm struct {
	rxs     []Rx
	capture bool
	name    *string
}

func (g groupedForm) String() string {
	var sb strings.Builder
	sb.WriteString("(")
	// TODO: Flags.
	if g.capture && g.name != nil {
		sb.WriteString("?P<")
		sb.WriteString(*g.name)
		sb.WriteString(">")
	}
	if !g.capture {
		sb.WriteString("?:")
	}
	for _, rx := range g.rxs {
		sb.WriteString(rx.String())
	}
	sb.WriteString(")")
	return sb.String()
}

// Capture creates a capture group (submatch) for the provided Rxs.
//
// See https://golang.org/pkg/regexp/#Regexp.FindStringSubmatch and other
// 'Submatch' methods for how to get these capture groups from matches.
func Capture(rxs ...Rx) Rx {
	return groupedForm{
		rxs:     rxs,
		capture: true,
		name:    nil,
	}
}

// NamedCapture creates a named capture group for the provided Rxs.
func NamedCapture(name string, rxs ...Rx) Rx {
	return &groupedForm{
		rxs:     rxs,
		capture: true,
		name:    &name,
	}
}

// Group groups many Rxs.
func Group(rxs ...Rx) Rx {
	return groupedForm{
		rxs:     rxs,
		capture: false,
		name:    nil,
	}
}

func implicitGroup(rxs ...Rx) Rx {
	if len(rxs) == 1 {
		return rxs[0]
	}
	return Group(rxs...)
}
