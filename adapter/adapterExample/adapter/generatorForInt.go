package adapter

type GeneratorForIntI interface {
	IntGenerator() int
}

type GeneratorForInt struct {
}

func (*GeneratorForInt) IntGenerator() int {
	return 19982727
}
