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

type StringMap struct {
	Name string
	Data map[string]string
}

func (strmap StringMap) put(key string, value string) {
	strmap.put(key, value)
}

func NewStringMap(name string) StringMap {
	m := StringMap{Name: name, Data: map[string]string{}}
	return m
}
