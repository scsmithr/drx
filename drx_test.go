package drx_test

import (
	"regexp"
	"testing"

	"github.com/scsmithr/drx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSimple(t *testing.T) {
	type testCase struct {
		name     string
		testRx   drx.Rx
		expected string
	}

	tcs := []testCase{
		{
			name: "BasicLiteralForms",
			testRx: drx.Build(
				drx.Bol,
				drx.Literal("hello"),
				drx.Eol,
			),
			expected: "^hello$",
		},
		{
			name: "BasicCompositeAnd",
			testRx: drx.Build(
				drx.Bol,
				drx.And(
					drx.Literal("hello"),
					drx.Space,
					drx.Literal("world"),
				),
				drx.Eol,
			),
			expected: "^hello[[:space:]]world$",
		},
		{
			name: "BasicCompositeOr",
			testRx: drx.Build(
				drx.Bol,
				drx.Or(
					drx.Literal("green"),
					drx.Literal("blue"),
				),
				drx.Eol,
			),
			expected: "^green|blue$",
		},
		{
			name:     "BasicRepetition",
			testRx:   drx.Build(drx.Bol, drx.ZeroOrMore(drx.Literal("hello"), false)),
			expected: "^hello*?",
		},
		{
			name: "BasicNtoMTimes",
			testRx: drx.Build(
				drx.Bol,
				drx.NtoMTimes(3, 5, true,
					drx.Literal("hello"),
					drx.Space,
					drx.Literal("world"),
				),
				drx.Eol,
			),
			expected: "^(?:hello[[:space:]]world){3,5}$",
		},
		{
			name: "BasicCapture",
			testRx: drx.Build(
				drx.Capture(
					drx.Or(
						drx.Literal("cat"),
						drx.Literal("dog"),
						drx.Literal("mouse"),
					),
				),
			),
			expected: "(cat|dog|mouse)",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			rx, err := drx.Compile(tc.testRx)
			require.NoError(t, err)

			stdRx, err := regexp.Compile(tc.expected)
			require.NoError(t, err)

			assert.Equal(t, stdRx.String(), rx.String())
		})
	}
}
