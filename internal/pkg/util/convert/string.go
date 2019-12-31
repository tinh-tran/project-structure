package convert

import (
	"fmt"
	"reflect"
)

func MustString(value interface{}) string {
	switch v := value.(type) {
	case fmt.Stringer:
		return v.String()
	case string:
		return v
	default:
		return fmt.Sprintf("%v", value)
	}
}

func ToString(value interface{}) (string, error) {
	return MustString(value), nil
}

func MustStringArray(array interface{}) (resArray []string) {
	t := reflect.TypeOf(array)
	switch t.Kind() {
	case reflect.Array, reflect.Slice:
		{
			v := reflect.ValueOf(array)
			resArray = make([]string, v.Len())
			for index, _ := range resArray {
				resArray[index] = MustString(v.Index(index).Interface())
			}
			return
		}
	}
	return []string{MustString(array)}
}

func ToStringArray(value interface{}) ([]string, error) {
	return MustStringArray(value), nil
}
