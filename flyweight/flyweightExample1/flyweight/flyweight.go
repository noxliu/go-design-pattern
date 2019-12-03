package flyweight

type FlyWeight struct {
	data string
}

var F = make(map[string]*FlyWeight)

func GetFLyWeight(fileName string) *FlyWeight {
	image := F[fileName]
	if image == nil {
		image = &FlyWeight{data: fileName}
		F[fileName] = image
	}

	return image
}
