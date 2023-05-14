package dog

import "testing"

func TestYears(t *testing.T) {
	v := Years(2)
	if v != 20 {
		t.Error("Expected : ", 20, " got ", v)
	}

}
