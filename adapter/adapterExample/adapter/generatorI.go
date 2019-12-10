package adapter

type generatorI interface {
	StringGenerator() string
	IntGenerator() int
}
