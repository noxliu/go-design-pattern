package adapter

type GeneratorForIntArrayI interface {
	IntArrayGenerator() int
}

type GeneratorForIntArray struct {
}

func (*GeneratorForIntArray) IntArrayGenerator() []int {
	array1 := []int{1, 2, 3, 999, 877777, 1212, 221212}
	return array1
}
