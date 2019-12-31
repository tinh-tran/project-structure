package convert

import "fmt"

func MustBool(value interface{}) bool {
	if res, err := ToBool(value); err == nil {
		return res
	} else {
		panic(err)
	}
}

func ToBool(value interface{}) (res bool, err error) {
	res = false
	err = nil
	switch value.(type) {
	case string:
		{
			str := MustString(value)
			switch str {
			case "1", "t", "T", "true", "TRUE", "True":
				{
					res = true
				}
			case "0", "f", "F", "false", "FALSE", "False", "":
				{
					res = false
				}
			default:
				err = fmt.Errorf("convert: %v to boolean failed.", value)
			}
		}
	case bool:
		{
			res = value.(bool)
		}
	default:
		valueStr := MustString(value)
		res, err = ToBool(valueStr)
	}
	return
}
