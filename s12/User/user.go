package user

type myData interface {
	int | string
}

func Sum[T myData](a, b T) T {
	return a + b
}
