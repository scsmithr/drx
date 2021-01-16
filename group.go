package drx

import "strings"

type groupedForm struct {
	rxs     []Rx
	capture bool
}

func (g groupedForm) String() string {
	var sb strings.Builder
	sb.WriteString("(")
	if !g.capture {
		sb.WriteString("?:")
	}
	for _, rx := range g.rxs {
		sb.WriteString(rx.String())
	}
	sb.WriteString(")")
	return sb.String()
}

func Capture(rxs ...Rx) Rx {
	return groupedForm{
		rxs:     rxs,
		capture: true,
	}
}

func Group(rxs ...Rx) Rx {
	return groupedForm{
		rxs:     rxs,
		capture: false,
	}
}

func implicitGroup(rxs ...Rx) groupedForm {
	return groupedForm{
		rxs:     rxs,
		capture: false,
	}
}
