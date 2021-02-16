package randhex

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"strings"
)

// RandHex is a container for a random hexadecimal color code
type RandHex struct {
	bytes []byte
}

// New generates a new random hexadecimal color code
func New() RandHex {
	bytes := make([]byte, 3)
	rand.Read(bytes)

	return RandHex{bytes: bytes}
}

// String provides the hexadecimal color code in string format
func (randhex RandHex) String() string {
	return strings.ToUpper("#" + hex.EncodeToString(randhex.bytes))
}

// Bytes provides the bytes of the hexadecimal color code
func (randhex RandHex) Bytes() []byte {
	val := make([]byte, 3)
	copy(val, randhex.bytes)
	return val
}

// ParseString parses a string representation of a hexcode into a RandHex
func ParseString(str string) (randhex RandHex, err error) {
	noHash := strings.Replace(str, "#", "", -1)
	noHash = strings.ToLower(noHash)

	if len(noHash) == 3 {
		noHash = string(noHash[0]) + string(noHash[0]) + string(noHash[1]) + string(noHash[1]) + string(noHash[2]) + string(noHash[2])
	}

	if len(noHash) != 6 {
		return RandHex{}, errors.New(str + " is not a valid hexadecimal color code. Hexadecimal color codes must be 3 or 6 digits.")
	}

	bytes, err := hex.DecodeString(noHash)

	if err != nil {
		return RandHex{}, err
	}

	return RandHex{bytes: bytes}, nil
}
