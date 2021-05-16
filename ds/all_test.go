package ds

import (
	"fmt"
	"reflect"
	"testing"
)

func ExampleRepeat() {
	res := Repeat("b")
	fmt.Println(res)
}

func TestAll(t *testing.T) {

	t.Run("running Integers Add test", func(t *testing.T) {
		sum := Add(1, 2)
		expected := 3
		if sum != 3 {
			t.Errorf("expected '%d' but got '%d'", expected, sum)
		}
	})

	t.Run("running Integers subtract test", func(t *testing.T) {
		diff := Sub(4, 5)
		expected := -1

		if diff != expected {
			t.Errorf("expected '%d' but got '%d'", expected, diff)
		}
	})

	t.Run("running String repeat test", func(t *testing.T) {
		res := Repeat("b")
		expected := "bbbbb"
		if res != expected {
			t.Errorf("expected '%s' but got '%s'", expected, res)
		}
	})

	t.Run("test Array sum", func(t *testing.T) {
		res := SumList([]int{1, 2, 3, 4})
		expected := 10
		if res != expected {
			t.Errorf("expected '%d' but got '%d'", expected, res)
		}
	})

	t.Run("test multiple array sum", func(t *testing.T) {
		res := SumAllList([]int{1, 2, 3, 4}, []int{4, 4, 4, 4, 4}, []int{5, 5, 5})
		expected := []int{10, 20, 15}

		if !reflect.DeepEqual(res, expected) {
			t.Errorf("expected '%v' but got '%v'", expected, res)
		}

	})
}