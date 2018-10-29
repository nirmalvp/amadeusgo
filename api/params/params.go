package params

import "fmt"

type Params map[string]string

func (params Params) And(key string, value interface{}) Params {
	// This is lazy and wont work unless the value is a string. Use reflection to check
	switch val := value.(type) {
	case int:
		// here val has type int
		params[key] = fmt.Sprintf("%d", val)
	case string:
		params[key] = fmt.Sprintf("%s", val)
	default:
		// no match; not handling now
	}
	return params
}

func With(key string, value interface{}) Params {
	params := make(Params)
	switch val := value.(type) {
	case int:
		// here val has type int
		params[key] = fmt.Sprintf("%d", val)
	case string:
		params[key] = fmt.Sprintf("%s", val)
	default:
		// no match; not handling now
	}
	return params
}
