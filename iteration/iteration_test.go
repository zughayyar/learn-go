package iteration

import (
	"fmt"
	"testing"
)

func TestIteration(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("Expected %s, got %s", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a")
	}
}

func ExampleRepeat() {
	repeated := Repeat("a")
	fmt.Println(repeated)
	// Output: aaaaa
}
