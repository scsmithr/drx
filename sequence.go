package drx

import (
	"fmt"
	"strings"
)

type compositeForm struct {
	rxs    []Rx
	joinBy string
}

func (c *compositeForm) String() string {
	elems := make([]string, len(c.rxs))
	for i, rx := range c.rxs {
		elems[i] = rx.String()
	}
	return strings.Join(elems, c.joinBy)
}

// And creates an Rx for matching all provided Rxs.
func And(rxs ...Rx) Rx {
	return &compositeForm{
		rxs:    rxs,
		joinBy: "",
	}
}

// Or creates an Rx for matching any Rxs, preferring to match left to right.
func Or(rxs ...Rx) Rx {
	return &compositeForm{
		rxs:    rxs,
		joinBy: "|",
	}
}

type repetitionOp string

const (
	opZeroOrMore repetitionOp = "*"
	opOneOrMore  repetitionOp = "+"
	opZeroOrOne  repetitionOp = "?"
)

type repetitionForm struct {
	rx     Rx
	op     repetitionOp
	greedy bool
}

func (r *repetitionForm) String() string {
	s := r.rx.String() + string(r.op)
	if !r.greedy {
		s = s + "?"
	}
	return s
}

// ZeroOrMore creates an Rx for matching zero or more Rxs. Rxs will be grouped
// (non-capture).
func ZeroOrMore(greedy bool, rxs ...Rx) Rx {
	return &repetitionForm{
		rx:     implicitGroup(rxs...),
		op:     opZeroOrMore,
		greedy: greedy,
	}
}

// OneOrMore creates an Rx for matching one or more Rxs. Rxs will be grouped
// (non-capture).
func OneOrMore(greedy bool, rxs ...Rx) Rx {
	return &repetitionForm{
		rx:     implicitGroup(rxs...),
		op:     opOneOrMore,
		greedy: greedy,
	}
}

// ZeroOrOne creates an Rx for matching zero or one Rxs. Rxs will be grouped
// (non-capture).
func ZeroOrOne(greedy bool, rxs ...Rx) Rx {
	return &repetitionForm{
		rx:     implicitGroup(rxs...),
		op:     opZeroOrOne,
		greedy: greedy,
	}
}

type boundedRepetitionForm struct {
	rx           Rx
	min          int64
	max          *int64
	unboundedMax bool
	greedy       bool
}

func (b *boundedRepetitionForm) String() string {
	s := fmt.Sprintf("%s{%d", b.rx.String(), b.min)
	if b.unboundedMax {
		s = fmt.Sprintf("%s,}", s)
	} else if b.max != nil {
		s = fmt.Sprintf("%s,%d}", s, *b.max)
	} else {
		s = s + "}"
	}
	if !b.greedy {
		s = s + "?"
	}

	return s
}

// NTimes creates an Rx for matching Rxs exactly n times. Rxs will be grouped
// (non-capture).
func NTimes(n int64, greedy bool, rxs ...Rx) Rx {
	return &boundedRepetitionForm{
		rx:           implicitGroup(rxs...),
		min:          n,
		max:          nil,
		unboundedMax: false,
		greedy:       greedy,
	}
}

// NOrMoreTimes creates an Rx for matching Rxs n or more times. Rxs will be
// grouped (non-capture).
func NOrMoreTimes(n int64, greedy bool, rxs ...Rx) Rx {
	return &boundedRepetitionForm{
		rx:           implicitGroup(rxs...),
		min:          n,
		max:          nil,
		unboundedMax: true,
		greedy:       greedy,
	}
}

// NTimes creates an Rx for matching Rxs n to m times. Rxs will be grouped
// (non-capture).
func NtoMTimes(n int64, m int64, greedy bool, rxs ...Rx) Rx {
	return &boundedRepetitionForm{
		rx:           implicitGroup(rxs...),
		min:          n,
		max:          &m,
		unboundedMax: false,
		greedy:       greedy,
	}
}
