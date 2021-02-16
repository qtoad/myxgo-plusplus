/*
Example usage:
	package main

	import (
		"fmt"

	)

	func main() {
		UUID4 := uuid.UUID4()
		fmt.Println("UUID4 is", UUID4)

		UUID5 := uuid.UUID5(uuid.NamespaceDNS, "SomeName")
		fmt.Println("UUID5", UUID5)
	}
*/
package uuid

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
)

const (
	reservedNcs       byte = 0x80 //Reserved for NCS compatibility
	rfc4122           byte = 0x40 //Specified in RFC 4122
	reservedMicrosoft byte = 0x20 //Reserved for Microsoft compatibility
	reservedFuture    byte = 0x00 // Reserved for future definition.
)

//The following standard UUIDs are for use with UUID3() or UUID5().
var (
	NamespaceDNS  = UUID{107, 167, 184, 16, 157, 173, 17, 209, 128, 180, 0, 192, 79, 212, 48, 200}
	NamespaceURL  = UUID{107, 167, 184, 17, 157, 173, 17, 209, 128, 180, 0, 192, 79, 212, 48, 200}
	NamespaceOID  = UUID{107, 167, 184, 18, 157, 173, 17, 209, 128, 180, 0, 192, 79, 212, 48, 200}
	NamespaceX500 = UUID{107, 167, 184, 20, 157, 173, 17, 209, 128, 180, 0, 192, 79, 212, 48, 200}
)

// UUID represents a UUID
type UUID [16]byte

func (u *UUID) setVariant(variant byte) {
	switch variant {
	case reservedNcs:
		u[8] &= 0x7F
	case rfc4122:
		u[8] &= 0x3F
		u[8] |= 0x80
	case reservedMicrosoft:
		u[8] &= 0x1F
		u[8] |= 0xC0
	case reservedFuture:
		u[8] &= 0x1F
		u[8] |= 0xE0
	}
}

func (u *UUID) setVersion(version byte) {
	u[6] = (u[6] & 0x0F) | (version << 4)
}

// UUID3 generates a version 3 UUID
func UUID3(namespace UUID, name string) UUID {
	var uuid UUID
	var version byte = 3
	hasher := md5.New()
	hasher.Write(namespace[:])
	hasher.Write([]byte(name))
	sum := hasher.Sum(nil)
	copy(uuid[:], sum[:len(uuid)])

	uuid.setVariant(rfc4122)
	uuid.setVersion(version)
	return uuid
}

// UUID4 generates a version 4 UUID
func UUID4() UUID {

	var uuid UUID
	var version byte = 4

	// Read is a helper function that calls io.ReadFull.
	_, err := rand.Read(uuid[:])
	if err != nil {
		panic(err)
	}

	uuid.setVariant(rfc4122)
	uuid.setVersion(version)
	return uuid
}

// UUID5 generates a version 5 UUID
func UUID5(namespace UUID, name string) UUID {
	var uuid UUID
	var version byte = 5
	hasher := sha1.New()
	hasher.Write(namespace[:])
	hasher.Write([]byte(name))
	sum := hasher.Sum(nil)
	copy(uuid[:], sum[:len(uuid)])

	uuid.setVariant(rfc4122)
	uuid.setVersion(version)
	return uuid
}

// String provides the uuid in a format compliant with RFC 4122
func (u UUID) String() string {
	return fmt.Sprintf("%x-%x-%x-%x-%x",
		u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
}

// ParseUUIDFromString parses a RFC 4122 compliant string representation of a uuid into a UUID4
func ParseUUIDFromString(str string) (uuid UUID, err error) {
	noDash := strings.Replace(str, "-", "", -1)
	noDash = strings.ToLower(noDash)
	var val UUID

	if len(noDash) != 32 {
		return val, errors.New(str + " is not a valid UUID4. The unhyphenated string representation should be 32 characters in length")
	}

	if noDash[12] != '4' {
		return val, errors.New(str + " is not a valid UUID4. character 13 should be '4'.")
	}

	if noDash[16] != '8' && noDash[16] != '9' && noDash[16] != 'a' && noDash[16] != 'b' {
		return val, errors.New(str + " is not a valid UUID4. character 17 should be '8', '9', 'a', or 'b'")
	}

	bytes := make([]byte, 16)
	bytes, err = hex.DecodeString(noDash)

	if err != nil {
		return val, err
	}

	//var i int
	for i := 0; i < 16; i++ {
		val[i] = bytes[i]
	}

	return val, nil
}
