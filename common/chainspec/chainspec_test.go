package chainspec

import "testing"

func TestKrumelanke(t *testing.T) {
	if Krummelanke.Name != "Krumme Lanke" {
		t.Fail()
	}

	first, ok := Krummelanke.Genesis.Raw["0x9768f3cbdd14c1a63474dfbdbe052f42"]

	if !ok {
		t.Fail()
	}
	if first != "0x80f4030000000000" {
		t.Fail()
	}
}
