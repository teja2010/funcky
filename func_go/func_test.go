package funcky

import (
	"testing"
)

func getRequire(t *testing.T) func(condn bool, msg string, args ...interface{}) {

	return func(condn bool, msg string, args ...interface{}) {
		if !condn {
			t.Errorf(msg, args)
		}
	}
}

func Int(i interface{}) int {
	return i.(int)
}

func TestCompose(t *testing.T) {

	require := getRequire(t)

	square := func(i IN) OUT {
		return Int(i) * Int(i)
	}

	double := func(i IN) OUT {
		return Int(i) + Int(i)
	}

	require(compose(square, double)(1) == 2, "1 squared and doubled is 2")
	require(compose(double, square)(1) == 4, "1 doubled and squared is 4")
	require(compose()(1) == 1, "no operations, so return input")
}
