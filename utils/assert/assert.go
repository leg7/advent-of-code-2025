package assert

import (
	"log"
)

func True(cond bool, format string, formatArgs ...any) {
	if !cond {
		log.Panicf("assert.True: " + format, formatArgs...)
		panic(cond)
	}
}

func Equals[T comparable](actual T, expected T) {
	if actual != expected {
		log.Panicf("assert.Equals: Expected value %v but got actual value %v\n", expected, actual)
	}
}

func NoError(err error) {
	if err != nil {
		log.Panic(err)
	}
}
