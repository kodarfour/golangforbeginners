package bigOrSmall

import "testing"

func TestBig(t *testing.T) {
	if !Big(51) {
		t.Log("UH OH 51 should be considered small.")
		t.Fail()
	}
}
