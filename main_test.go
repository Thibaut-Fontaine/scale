package main

import "testing"

func TestShiftScale(t *testing.T) {
	if shiftScale(0x42, 0) != 0x42 {
		t.Error("expected 0x42 on 0x42, 0")
	}
	if shiftScale(0x6d2, 3) != 0x693 {
		t.Error("expected 0x693 on 0x6d2, 3")
	}
	if shiftScale(0x100, 5) != 0x2 {
		t.Error("expected 0x2 on 0x100, 5")
	}
	if shiftScale(0xfdf, 8) != 0xffd {
		t.Error("expected 0xffd on 0xfdf, 8")
	}
	if shiftScale(0x110, -6) != 0x404 {
		t.Error("expected 0x404 on 0x110, -6")
	}
}
