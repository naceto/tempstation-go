package sdk

// IIf -
func IIf[T any](cond bool, first T, second T) T {
	if cond {
		return first
	}

	return second
}

func Ptr[T any](t T) *T {
	return &t
}
