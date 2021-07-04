package funcky

type KF func(in IN) (out OUT, m Monoid)

// composition in kliesli is how we build monads
// i.e. arrows (a -> mb) and (b -> mc) in category C
// become arrows (a -> b) and (b -> c) in Kliesli category K

// >>= bind operator
func bind(k1, k2 KF) KF {
	return func(in IN) (OUT, Monoid) {
		out1, m1 := k1(in)
		out2, m2 := k2(out1)
		// use m1 & m2 in some way
		return out2, composeEffects(m1, m2)
	}
}

func composeEffects(m1, m2 Monoid) Monoid {
	// very specific to what we want to do. here ignore m2
	return m1
}
