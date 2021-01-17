package drx

import "regexp"

type literalForm struct {
	literal string
}

func (l literalForm) String() string {
	return l.literal
}

// Literal creates an Rx for the literal string. No escaping is done.
func Literal(s string) Rx {
	return literalForm{
		literal: s,
	}
}

// QuoteLiteral creates an Rx that will match the literal string, escaping all
// regex meta characters.
//
// See https://golang.org/pkg/regexp/#QuoteMeta for more details.
func QuoteLiteral(s string) Rx {
	return literalForm{
		literal: regexp.QuoteMeta(s),
	}
}

type StringLiteral string

func (s StringLiteral) String() string {
	return string(s)
}

const (
	Any    StringLiteral = "."
	Alnum  StringLiteral = "[[:alnum:]]"
	Alpha  StringLiteral = "[[:alpha:]]"
	Ascii  StringLiteral = "[[:ascii:]]"
	Blank  StringLiteral = "[[:blank:]]"
	Space  StringLiteral = "[[:space:]]"
	Digit  StringLiteral = "[[:digit:]]"
	Lower  StringLiteral = "[[:lower:]]"
	Upper  StringLiteral = "[[:upper:]]"
	Word   StringLiteral = "[[:word:]]"
	Xdigit StringLiteral = "[[:xdigit:]]"
	Hex                  = Xdigit
)

const (
	LineStart       StringLiteral = "^"
	Bol                           = LineStart
	LineEnd         StringLiteral = "$"
	Eol                           = LineEnd
	WordBoundary    StringLiteral = `\b`
	NotWordBoundary StringLiteral = `\B`
)

type characterClassForm struct {
	charset string
	negate  bool
}

func (c characterClassForm) String() string {
	if c.negate {
		return "[^" + c.charset + "]"
	}
	return "[" + c.charset + "]"
}

// Charset creates an Rx for for matching any character in the provided charset.
func Charset(charset string) Rx {
	return characterClassForm{
		charset: charset,
		negate:  false,
	}
}

// NotCharset creates an Rx for for matching any character not in the provided
// charset.
func NotCharset(charset string) Rx {
	return characterClassForm{
		charset: charset,
		negate:  true,
	}
}
