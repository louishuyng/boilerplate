package helpers

import "strconv"

func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}

	return v
}
func StrToInt(v string) int {
	i, err := strconv.Atoi(v)
	if err != nil {
		panic(err)
	}

	return i
}
