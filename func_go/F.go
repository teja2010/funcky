package funcky

type IN interface {
}

type OUT interface {
}

// a simple function that takes one in and one out
type F (func(in IN) OUT)
