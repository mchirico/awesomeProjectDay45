package util

import "testing"

func TestUT_DoSomething(t *testing.T) {
	ut := NewUT()
	expecting := "Initial"
	if ut.Name != expecting {
		t.Fatalf("Expecting %s. Got: %s\n", expecting, ut.Name)
	}
}
