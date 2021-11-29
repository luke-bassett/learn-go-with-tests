package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 3)
	expected := "aaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func TestReplace(t *testing.T) {
	replaced := Replace("hip", "i", "o")
	expected := "hop"

	if replaced != expected {
		t.Errorf("expected %q but got %q", expected, replaced)
	}

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("ha", 3)
	fmt.Println(repeated)
	// Output: hahaha
}
