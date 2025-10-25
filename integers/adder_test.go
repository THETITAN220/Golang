package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4
	if sum != expected {
		t.Errorf("expected '%d' received '%d' ", expected, sum)
	}
}

func ExampleAdd() {
	sum := Add(2, 5)
	fmt.Printf("Sum: %d", sum)
	// output: Sum: 7
}
