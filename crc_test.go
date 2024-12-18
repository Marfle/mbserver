package mbserver

import "testing"

func TestCRC(t *testing.T) {
	got := crcModbus([]byte{0x01, 0x04, 0x02, 0xFF, 0xFF})
	expect := 0x80B8
	if !isEqual(expect, got) {
		t.Errorf("expected %x, got %x", expect, got)
	}
}

func TestCRC2(t *testing.T) {
	got := crcModbus([]byte{0x01, 0x03, 0x00, 0x35, 0x00, 0x18})
	expect := 0xce55 // flipped in intel
	if !isEqual(expect, got) {
		t.Errorf("expected %x, got %x", expect, got)
	}
}

func TestCRC3(t *testing.T) {
	got := crcModbus([]byte{0x01, 0x02, 0x00, 0x00, 0x00, 0x10})
	expect := 0xc679 // flipped in intel
	if !isEqual(expect, got) {
		t.Errorf("expected %x, got %x", expect, got)
	}
}
func TestCRC4(t *testing.T) {
	got := crcModbus([]byte{0x01, 0x02, 0x00, 0x40, 0x00, 0x20})
	expect := 0x0678 // flipped in intel
	if !isEqual(expect, got) {
		t.Errorf("expected %x, got %x", expect, got)
	}
}

func TestCRC5(t *testing.T) {
	got := crcModbus([]byte{0x01, 0x02, 0x00, 0xd0, 0x00, 0x47})
	expect := 0xc139 // flipped in intel
	if !isEqual(expect, got) {
		t.Errorf("expected %x, got %x", expect, got)
	}
}

func TestCRC6(t *testing.T) {
	got := crcModbus([]byte{0x02, 0x03, 0x00, 0x00, 0x00})
	expect := 0x3404 // flipped in intel
	if !isEqual(expect, got) {
		t.Errorf("expected %x, got %x", expect, got)
	}
}

func TestCRC7(t *testing.T) {
	got := crcModbus([]byte{0x01, 0x02, 0x00, 0x00, 0x00, 0x10})
	expect := 0xc619 // flipped in intel
	if !isEqual(expect, got) {
		t.Errorf("expected %x, got %x", expect, got)
	}
}