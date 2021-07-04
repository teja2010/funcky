package funcky

import (
	"fmt"
	"testing"
)

// product: best way to HOLD two values without any noise. Any other
// combination of the value scan be derived from this value.
// canonical example: tuple
// corresponds to multiplication in algebra.
// increasing state space
// (a,b) == (b, a) order does not matter.

// a closure to hold two values. they are returned when it is evaluated
// a closure is a stricter product that a struct type wise. A struct can
// be constructed without all of its elements. this closure cant.
type intTuple (func() (int, int))

func getIntTuple(a, b int) intTuple {
	return func() (int, int) {
		return a, b
	}
}

func fst(t intTuple) int {
	f, _ := t()
	return f
}

func snd(t intTuple) int {
	_, s := t()
	return s
}

func TestProduct(t *testing.T) {
	x := getIntTuple(1, 2)
	require := getRequire(t)

	require(fst(x) == 1, "fst check")
	require(snd(x) == 2, "snd check")
}

// sum: best to hold exactly one of them without any noise.
// any other way can be derived from this value
// canonical example: Either a b
// implemented by an interface.
// corresponds to addition in algebra.
// transforming state space, but no increase.
// order matters, choosing a & choosing b are completely different.
type either interface {
	Val() string
}

type eitherInt struct{ i int32 }

func (e eitherInt) Val() string { return string(e.i) }

type eitherFloat struct{ f float32 }

func (e eitherFloat) Val() string { return fmt.Sprintf("%f", e.f) }
