package adapter

type generatorI interface {
	StringGenerator() string
}

type GeneratorForString struct {
}

func (*GeneratorForString) StringGenerator() string {
	return "Hello world"
}
