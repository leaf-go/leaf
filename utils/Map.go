package utils

type IMap interface{}


// todo 支持泛型后重构
//type Map map[interface{}]interface{}

type mapIntString struct {
	data   map[int]string
	length int
}

func NewMapIntString(data map[int]string) *mapIntString {
	return &mapIntString{data: data, length: len(data)}
}

func (ism *mapIntString) Keys() (keys []int) {
	keys = make([]int, ism.length)
	index := 0
	for k, _ := range ism.data {
		keys[index] = k
		index++
	}

	return
}

func (ism mapIntString) Values() (values []string) {
	values = make([]string, ism.length)
	index := 0
	for _, v := range ism.data {
		values[index] = v
		index++
	}
	return
}

func (ism mapIntString) Convert() (data map[string]int) {
	data = make(map[string]int, ism.length)
	for k, v := range ism.data {
		data[v] = k
	}
	return
}

type mapStringInt struct {
	data   map[string]int
	length int
}

func NewMapStringInt(data map[string]int) *mapStringInt {
	return &mapStringInt{data: data, length: len(data)}
}

func (sim *mapStringInt) Keys() (keys []string) {
	keys = make([]string, sim.length)
	index := 0
	for k, _ := range sim.data {
		keys[index] = k
		index++
	}
	return keys
}

func (sim *mapStringInt) Values() (values []int) {
	values = make([]int, sim.length)
	index := 0
	for _, v := range sim.data {
		values[index] = v
		index++
	}

	return
}

func (sim mapStringInt) Convert() (data map[int]string) {
	data = make(map[int]string, sim.length)
	for k, v := range sim.data {
		data[v] = k
	}
	return
}
