# drx [![Go Reference](https://pkg.go.dev/badge/github.com/scsmithr/drx.svg)](https://pkg.go.dev/github.com/scsmithr/drx)

`drx` is a Go library for declaratively building regular expressions.

## Example

(API definitely subject to change)

```go
package main

import (
	"fmt"

	"github.com/scsmithr/drx"
)

func main() {
	const s = "The quick brown fox jumps over the lazy dog"

	rx := drx.MustCompile(drx.Build(
		drx.Bol,
		drx.ZeroOrMore(false, drx.Any),
		drx.Capture(
			drx.Or(
				drx.Literal("green"),
				drx.Literal("brown"),
			),
		),
		drx.ZeroOrMore(false, drx.Any),
		drx.Literal("lazy"),
		drx.Space,
		drx.Capture(
			drx.ZeroOrMore(false, drx.Alpha),
		),
		drx.Eol,
	))

	fmt.Printf("%q\n", rx.FindStringSubmatch(s))
	// ["The quick brown fox jumps over the lazy dog" "brown" "dog"]
}
```

