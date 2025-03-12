package util

func When[T any](cond bool, tv, fv T) T {
	if cond {
		return tv
	}
	return fv
}
