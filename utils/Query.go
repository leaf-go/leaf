package utils

import (
	"fmt"
	"net/url"
)

var (
	Query = &query{}
)

type query struct {
}

func (q *query) Bind(m map[string]interface{} , format string) string  {
	values := url.Values{}
	var nk, ret string
	for k, v := range m {
		if len(format) != 0 {
			nk = fmt.Sprintf(format, k)
		} else {
			nk = k
		}

		switch v.(type) {
		case string:
			values.Add(nk, v.(string))
			break
		case HashBytes:
			values.Add(nk, v.(HashBytes).String())
			break
		case map[string]interface{}:
			ret += q.Bind(v.(map[string]interface{}), nk+"[%s]")
			ret += "&"
		case int64, int32, int16, int8, int, uint64, uint32, uint16, uint8, uint:
			values.Add(nk, fmt.Sprintf("%d", v))
		}
	}

	ret += values.Encode()
	return ret
}