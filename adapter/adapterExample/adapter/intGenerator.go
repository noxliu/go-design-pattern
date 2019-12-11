package adapter

type IntGeneratorI interface {
	IntGenerator() int
}

type IntGenerator struct {
}

func (*IntGenerator) IntGenerator() int {
	return 19982727
}
