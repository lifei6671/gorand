package gorand

import (
	"strings"
	"testing"
)

// TestUUID4StringIsCorrectLength tests that the length of a generated uuid4 string is 36 characters
func TestUUID4StringIsCorrectLength(t *testing.T) {
	l := len(NewUUID4().String())

	if l != 36 {
		t.Errorf("String() length was incorrect, expected 36, got %d", l)
	}
}

// TestUUID4StringContainsCorrectHyphens tests that the uuid4 string contains
// the correct hyphens in compliance with RFC 4122
func TestUUID4StringContainsCorrectHyphens(t *testing.T) {
	uuidStr := NewUUID4().String()

	if uuidStr[8] != '-' {
		t.Errorf("String()[8] was incorrect, expected '-', got %d", uuidStr[8])
	}

	if uuidStr[13] != '-' {
		t.Errorf("String()[13] was incorrect, expected '-' got %d", uuidStr[13])
	}

	if uuidStr[18] != '-' {
		t.Errorf("String()[18] was incorrect, expected '-' got %d", uuidStr[18])
	}

	if uuidStr[23] != '-' {
		t.Errorf("String()[23] was incorrect, expected '-' got %d", uuidStr[23])
	}
}

// TestUUID4BytesReturnsCorrectBytes tests that the bytes return by Bytes are correct
func TestUUID4BytesReturnsCorrectBytes(t *testing.T) {
	uuid := NewUUID4()
	bytes := uuid.Bytes()

	if len(bytes) != 16 {
		t.Errorf("Bytes() was the incorrect length, expected 16, got %d", len(bytes))
	}

	for i, b := range bytes {
		if b != uuid.bytes[i] {
			t.Errorf("Bytes()[%d] != bytes[%d]", i, i)
		}
	}
}

//TestUUID4BytesNotInternalSlice tests that Bytes() does not return the same slice used internally
func TestUUID4BytesNotInternalSlice(t *testing.T) {
	uuid := NewUUID4()
	bytes := uuid.Bytes()

	bytes[3] += 0x1

	if bytes[3] == uuid.bytes[3] {
		t.Errorf("Bytes() and bytes are referencing the same slice")
	}
}

//TestUUID4ParseStringCorrectlyParses tests that ParseString() correctly parses the string to a uuid4
func TestUUID4ParseStringCorrectlyParses(t *testing.T) {
	str := "cc2161ae-33c1-4cb1-aa53-e81000f20a30"
	uuid, err := ParseString(str)

	if err != nil {
		t.Errorf("Error Parsing the string")
		return
	}

	if uuid.String() != str {
		t.Errorf("ParseString() did not correctly parse")
	}
}

//TestUUID4ParseStringCorrectlyParsesWithoutDashes tests that ParseString() correctly parses the string to a uuid4
func TestUUID4ParseStringCorrectlyParsesWithoutDashes(t *testing.T) {
	str := "cc2161ae33c14cb1aa53e81000f20a30"
	uuid, err := ParseString(str)

	if err != nil {
		t.Errorf("Error Parsing the string")
		return
	}

	if str != strings.Replace(uuid.String(), "-", "", -1) {
		t.Errorf("ParseString() did not correctly parse")
	}
}

// TestUUID4ParseStringReturnsErrorOnIndex12NotValid tests that ParseString() gives an error when str[12] is not in compliance
// with RFC 4122
func TestUUID4ParseStringReturnsErrorOnIndex12NotValid(t *testing.T) {
	str := "cc2161ae-33c1-bcb1-aa53-e81000f20a30"
	_, err := ParseString(str)

	if err == nil {
		t.Errorf("ParseString() should have failed. str[12] is invalid.")
	}
}

// TestUUID4ParseStringReturnsErrorOnIndex16NotValid tests that ParseString() gives an error when str[16] is not in compliance
// with RFC 4122
func TestUUID4ParseStringReturnsErrorOnIndex16NotValid(t *testing.T) {
	str := "cc2161ae-33c1-4cb1-ca53-e81000f20a30"
	_, err := ParseString(str)

	if err == nil {
		t.Errorf("ParseString() should have failed. str[16] is invalid.")
	}
}

// TestUUID4ParseStringReturnsErrorOnBadLength tests that ParseString() gives an error when str is not 32 characters
func TestUUID4ParseStringReturnsErrorOnBadLength(t *testing.T) {
	str := "cc2161ae-33c1-4cb1-aa53-e81000f20a"
	_, err := ParseString(str)

	if err == nil {
		t.Errorf("ParseString() should have failed. str is invalid length")
	}
}

// TestUUID4ParseStringReturnsErrorOnInvalidHex tests that ParseString() gives an error when str is not valid hex
func TestUUID4ParseStringReturnsErrorOnInvalidHex(t *testing.T) {
	str := "cc2161ae-33c1-4cb1-aa53-e81g00f20a30"
	_, err := ParseString(str)

	if err == nil {
		t.Errorf("ParseString() should have failed. str is not hex")
	}
}
