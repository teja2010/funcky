package funcky

import (
	"fmt"
)

func full(a, b, c IN) OUT {
	fmt.Println(a, b, c)
	return nil
}

// closures to apply partially
func part(a IN) func(IN, IN) OUT {
	return func(b, c IN) OUT {
		return full(a, b, c)
	}
}
