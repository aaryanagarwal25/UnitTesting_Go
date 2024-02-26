package Hello

import (
	"testing"
)

func TestAdd(t *testing.T) {
	if Add(4, 5) != 9 {
		t.Error("Incorrect ans with input 4 and 5")
	}
	if Add(-4, 5) != 1 {
		t.Error("Incorrect ans with input -4 and 5")
	}
	if Add(4, -5) != -1 {
		t.Error("Incorrect ans with input 4 and -5")
	}
	if Add(-4, -5) != -9 {
		t.Error("Incorrect ans with input -4 and -5")
	}
	if Add(0, 5) != 5 {
		t.Error("Incorrect ans with input 0 and 5")
	}
}
