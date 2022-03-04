package utils

type params struct {
	data map[string]interface{}
}

func NewParams(data map[string]interface{}) *params {
	return &params{
		data: data,
	}
}

func (p *params) InArrayInt(key string, def int, slice []int) int {
	value, ok := p.data[key].(int)
	if !ok {
		return def
	}

	v, exist := Slice.InSliceInt(value, slice)
	if !exist {
		return def
	}

	return v
}

func (p *params) GetString(key string, def string) string {
	v := p.get(key, def)
	return v.(string)
}

func (p params) GetInt(key string, def int) int {
	v := p.get(key, def)
	return v.(int)
}

func (p *params) get(key string, def interface{}) (value interface{}) {
	value, ok := p.data[key]
	if !ok {
		return def
	}

	return value
}
