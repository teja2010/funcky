package funcky

// f2 ֯֯o f1 : f2 after f1
func compose(fs ...F) F {
	return func(in IN) OUT {
		for _, f := range fs {
			in = f(in)
		}
		return in
	}
}
