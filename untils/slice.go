package untils

func Push(slice []interface{}, ele interface{}) []interface{} {
	if slice == nil {
		panic("slice is nil")
	}
	if ele == nil {
		panic("element is nil")
	}
	return append(slice, ele)
}

