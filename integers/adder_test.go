package integers

import (
	"testing"
	"fmt"
)

func TestAdd(t *testing.T) {
	sum := Add(3, 2)
	expected := 5

	if sum != expected {
		t.Errorf("expected '%d' got '%d'", expected, sum)
	}
}

func ExampleAdd() {
	sum := Add(3, 4)
	fmt.Println(sum)
	// Output: 7
}