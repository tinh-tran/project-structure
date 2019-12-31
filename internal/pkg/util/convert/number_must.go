package convert

func MustInt(value interface{}) (res int) {
	if res, err := ToInt(value); err == nil {
		return res
	} else {
		panic(err)
	}
}
