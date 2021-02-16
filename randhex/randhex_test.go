package randhex

import "testing"

func TestRandHexStringIsCorrectLength(t *testing.T) {
	l := len(New().String())

	if l != 7 {
		t.Errorf("String() length was incorrect, expected 7, got %d", l)
	}
}

func TestRandHexStringStartsWithHash(t *testing.T) {
	h := New().String()

	if h[0] != '#' {
		t.Errorf("String() started with incorrect character, expected #, got %c", h[0])
	}
}

func TestRandHexBytesCorrectLength(t *testing.T) {
	l := len(New().Bytes())

	if l != 3 {
		t.Errorf("Bytes() length was incorrect, expected 3, got %d", l)
	}
}

func TestParseStringBadLengthErrors(t *testing.T) {
	_, err := ParseString("#12345")

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestParseStringValidLen3(t *testing.T) {
	r, err := ParseString("#aaa")

	if err != nil {
		t.Errorf("Expected nil, got error")
	}

	if r.String() != "#AAAAAA" {
		t.Errorf("Expected #AAAAAA, got %s", r.String())
	}
}

func TestParseStringValidLen6(t *testing.T) {
	r, err := ParseString("#aaaaaa")

	if err != nil {
		t.Errorf("Expected nil, got error")
	}

	if r.String() != "#AAAAAA" {
		t.Errorf("Expected #AAAAAA, got %s", r.String())
	}
}

func TestParseStringBadCharRange(t *testing.T) {
	_, err := ParseString("#gaa")

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestParseStringBadSymbol(t *testing.T) {
	_, err := ParseString("%aaa")

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestParseStringNoHash(t *testing.T) {
	r, err := ParseString("aaa")

	if err != nil {
		t.Errorf("Expected nil, got error")
	}

	if r.String() != "#AAAAAA" {
		t.Errorf("Expected #AAAAAA, got %s", r.String())
	}
}
