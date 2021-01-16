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

func And(rxs ...Rx) Rx {
	return &compositeForm{
		rxs:    rxs,
		joinBy: "",
	}
}

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

func ZeroOrMore(greedy bool, rxs ...Rx) Rx {
	return &repetitionForm{
		rx:     implicitGroup(rxs...),
		op:     opZeroOrMore,
		greedy: greedy,
	}
}

func OneOrMore(greedy bool, rxs ...Rx) Rx {
	return &repetitionForm{
		rx:     implicitGroup(rxs...),
		op:     opOneOrMore,
		greedy: greedy,
	}
}

func ZeroOrOne(greedy bool, rxs ...Rx) Rx {
	return &repetitionForm{
		rx:     implicitGroup(rxs...),
		op:     opZeroOrOne,
		greedy: greedy,
	}
}

type boundedRepetitionForm struct {
	rx     Rx
	min    int64
	max    *int64
	greedy bool
}

func (b *boundedRepetitionForm) String() string {
	s := fmt.Sprintf("%s{%d", b.rx.String(), b.min)
	// TODO: Need min or more
	if b.max != nil {
		s = fmt.Sprintf("%s,%d}", s, *b.max)
	} else {
		s = s + "}"
	}
	if !b.greedy {
		s = s + "?"
	}

	return s
}

func NTimes(n int64, greedy bool, rxs ...Rx) Rx {
	return &boundedRepetitionForm{
		rx:     implicitGroup(rxs...),
		min:    n,
		greedy: greedy,
	}
}

func NtoMTimes(n int64, m int64, greedy bool, rxs ...Rx) Rx {
	return &boundedRepetitionForm{
		rx:     implicitGroup(rxs...),
		min:    n,
		max:    &m,
		greedy: greedy,
	}
}
