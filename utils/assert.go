package utils

import (
	"fmt"
	"os"
)

func Assert(cond bool, format string, formatArgs ...any) {
	if !cond {
		fmt.Fprintf(os.Stderr, format, formatArgs...)
		panic(cond)
	}
}

