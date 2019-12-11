package adapter

type GeneratorForString struct {
}

func (*GeneratorForString) StringGenerator() string {
	return "Hello world"
}
