package iteration

import (
	"fmt"
	"testing"
)

func TestIteration(t *testing.T) {

	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expect %q got %q", expected, repeated)
	}

}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		_ = Repeat("a", 5)
	}
}

func ExampleRepeat() {
	expected := Repeat("a", 5)
	fmt.Println(expected)
	// Output: aaaaa
}
