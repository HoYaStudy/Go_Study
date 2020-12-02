package calc

import "testing"

func TestAdder(t *testing.T) {
	testTable := []struct {
		x int
		y int
		z int
	}{
		{1, 2, 3},
		{0, 5, 5},
		{3, 4, 7},
		{4, -2, 2},
	}
	for _, elem := range testTable {
		result := Adder(elem.x, elem.y)
		if result != elem.z {
			t.Log("Expected: ", elem.z, "Got: ", result)
			t.Fail()
		} else {
			t.Log("SUCCESS")
		}
	}
}
