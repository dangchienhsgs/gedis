package core

type IntegerMap struct {
	Name string
	Data map[string]int
}

func (intmap IntegerMap) put(key string, value int) {
	intmap.put(key, value)
}

func NewIntegerMap(name string) IntegerMap {
	m := IntegerMap{Name: name, Data: map[string]int{}}
	return m
}
