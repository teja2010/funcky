package funcky

// use closures to represent tuples, since go does not allow it

type T interface{}

func tuple(a, b T) (T, T) {
	return a, b
}

func triple(a, b, c T) (T, T, T) {
	return a, b, c
}
