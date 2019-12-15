package adapter

type stringGeneratorI interface {
	stringGenerator() string
}

type StringGenerator struct {
}

func (*StringGenerator) stringGenerator() string {
	return "Hello world"
}
