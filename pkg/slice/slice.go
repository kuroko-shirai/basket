package slice

func Remove[T comparable](slice []T, index int) []T {
	var empty T

	slice[index] = slice[len(slice)-1]
	slice[len(slice)-1] = empty
	slice = slice[:len(slice)-1]

	return slice
}
