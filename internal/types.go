package internal

// Тип вынесен в интерфейс, так как в ряде случаев может быть необходим целочисленный тип. В иных случаях --
// вещественный
type TypeOfDistance interface {
	More(TypeOfDistance) bool
	Plus(TypeOfDistance) TypeOfDistance
	Minus(TypeOfDistance) TypeOfDistance
}

type Int int
type Float float32

func (baseObject Int) More(object TypeOfDistance) bool {
	return int(baseObject) > int(object.(Int))
}

func (baseObject Int) Plus(object TypeOfDistance) TypeOfDistance {
	return Int(int(baseObject) + int(object.(Int)))
}

func (baseObject Int) Minus(object TypeOfDistance) TypeOfDistance {
	return Int(int(baseObject) - int(object.(Int)))
}

func (baseObject Float) More(object TypeOfDistance) bool {
	return float32(baseObject) > float32(object.(Float))
}

func (baseObject Float) Plus(object TypeOfDistance) TypeOfDistance {
	return Float(float32(baseObject) + float32(object.(Float)))
}

func (baseObject Float) Minus(object TypeOfDistance) TypeOfDistance {
	return Float(float32(baseObject) - float32(object.(Float)))
}
