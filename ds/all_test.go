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

	t.Run("test Perimeter Rectangle", func(t *testing.T) {
		rec := Rectangle{10, 6}
		perm := rec.Perimeter()
		requiredPerm := 32.0

		if !reflect.DeepEqual(perm, requiredPerm) {
			t.Errorf("expected '%f' but got '%f'", requiredPerm, perm)
		}
	})

	t.Run("test Area Rectangle", func(t *testing.T) {
		rec := Rectangle{10, 6}
		area := rec.Area()
		requiredArea := 60.0

		if !reflect.DeepEqual(area, requiredArea) {
			t.Errorf("expected '%f' but got '%f'", requiredArea, area)
		}

	})

	t.Run("Wallet system using pointers", func(t *testing.T) {
		w := Wallet{}
		// explicitly making it call by reference.
		deposit(&w, 20)
		// explicitly making it call by reference.
		bal := balance(&w)
		req := 20

		if bal != req {
			t.Errorf("expected '%d' but got '%d'", req, bal)
		}
	})
}
