package iteration

import (
	"fmt"
	"math/rand"
	"testing"
	"strings"
)

func TestRepeat(t *testing.T) {
	times := 7
    repeated := Repeat("a", times)
    expected := strings.Repeat("a", times)

	if repeated != expected {
		t.Errorf("expected '%s' but got '%s'", expected, repeated)
	}
}
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", rand.Intn(100000))
	}
}
func ExampleRepeat() {
	repeated := Repeat("a", 8)
	fmt.Println(repeated)
	// Output: aaaaaaaa
}