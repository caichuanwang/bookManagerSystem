package untils

import "github.com/samber/lo"

func Push[T any](slice []T, ele T) []T {
	if slice == nil {
		panic("slice is nil")
	}
	return append(slice, ele)
}

func Join(slice []string, ele string) string {
	return lo.Reduce[string, string](slice, func(agg string, item string, i int) string {
		if i == 0 {
			return agg + item
		} else {
			return agg + ele + item
		}
	}, "")

}
