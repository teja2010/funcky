package funcky

func ListFmap(list []IN, f func(IN) OUT) []OUT {
	out := []OUT{}
	for _, i := range list {
		out = append(out, f(i))
	}
	return out
}
