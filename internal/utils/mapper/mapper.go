package mapper

// Map applies function to every item of an iterable and returns a slices of the results
func Map[T, M any](a []T, f func(*T) M) []M {
	n := make([]M, len(a))
	for i := range a {
		n[i] = f(&a[i])
	}

	return n
}
