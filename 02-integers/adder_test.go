package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func ExampleAdd() {
	// In order for this function to run, you need to run "go test -v"
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
