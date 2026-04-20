package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	t.Run("A test to check the sum of two intergers", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4

		assertCorrectResult(t, sum, expected)
	})
}

func assertCorrectResult(t testing.TB, sum, expected int) {
	t.Helper()
	if sum != expected {
		t.Errorf("got %d does not equal %d want", sum, expected)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
