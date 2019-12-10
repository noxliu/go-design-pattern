package adapter

type generatorI interface {
	StringGenerator() string
	IntegerGenerator() int
}
