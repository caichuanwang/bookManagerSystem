package untils

import (
	"reflect"
	"testing"
)

func TestJoin(t *testing.T) {
	t.Run("testJoin", func(t *testing.T) {
		got := Join([]string{"a", "b", "d"}, "c")
		want := "acbcd"
		if got == want {
			t.Log("testJoin pass")
		} else {
			t.Errorf("got %s but want %s", got, want)
		}
	})
}

type person struct {
	name string
	age  uint
}

func TestPush(t *testing.T) {
	var arrs = []struct {
		input    []any
		ele      any
		expected []any
	}{
		{[]any{"1", "2"}, "3", []any{"1", "2", "3"}},
		{[]any{1, 2}, 3, []any{1, 2, 3}},
		{[]any{person{
			name: "test",
			age:  12,
		}, person{
			name: "test2",
			age:  13,
		}}, person{
			name: "test3",
			age:  14,
		}, []any{person{
			name: "test",
			age:  12,
		}, person{
			name: "test2",
			age:  13,
		}, person{
			name: "test3",
			age:  14,
		}}},
	}
	for _, arr := range arrs {
		got := Push[any](arr.input, arr.ele)
		if reflect.DeepEqual(got, arr.expected) {
			t.Log("pass")
		} else {
			t.Errorf("%s error", arr.input)
		}
	}
}
