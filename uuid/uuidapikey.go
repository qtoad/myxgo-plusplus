// Copyright (C) 2019 SIESGSTarena

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this img except in compliance with the License.
// You may obtain a copy of the License at

//    http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package uuidapikey contains generator, validator, and converter that transforms UUIDs into human-readable Base32-Crockford encoded API Keys
package uuid

// Dependencies
import (
	"encoding/json"
	"github.com/qtoad/myxgo-plusplus/crock32"
	"regexp"
	"strconv"
	"strings"
)

// checkDashes to verify dashes in given UUID
func checkDashes(positions []int, uuid string) bool {
	test := true
	for index := range positions {
		chr := uuid[positions[index]]
		test = test && chr == '-'
	}
	return test
}

// encode will convert your given int64 into base32 crockford encoding format
func encode(number uint64) string {
	if number < 0 {
		panic("The given input must be greater than or equal to zero")
	}
	encoded := crock32.Encode(number)
	paddingTimes := 7 - len(encoded)
	return strings.ToUpper((strings.Repeat("0", paddingTimes) + encoded))
}

// decode will convert your given string into original UUID part string
func decode(decodedString string) string {
	integerDecoded, _ := crock32.Decode(decodedString)
	decoded := strconv.FormatUint(integerDecoded, 16)
	paddingTimes := 8 - len(decoded)
	return (strings.Repeat("0", paddingTimes) + decoded)
}

// IsUUID will verify given UUID
func IsUUID(uuid string) bool {
	if uuid == "" {
		panic("The required parameter UUID is undefined")
	}
	_, error := ParseUUIDFromString(uuid)
	if error != nil {
		return false
	}
	return true
}

// IsAPIKey will verify given API Key
func IsAPIKey(apiKey string) bool {
	if apiKey == "" {
		panic("The required parameter API Key is undefined")
	}
	apiKey = strings.ToUpper(apiKey)
	var re = regexp.MustCompile("-")
	apiKey = re.ReplaceAllString(apiKey, "")
	re = regexp.MustCompile("[0-9A-Z]{28}")
	return (len(apiKey) == 28 && re.MatchString(apiKey))
}

// ToAPIKey will convert given UUID to API Key
func ToAPIKey(uuid string) string {
	if uuid == "" {
		panic("The required parameter UUID is undefined")
	}
	if IsUUID(uuid) {
		var re = regexp.MustCompile("-")
		uuid = re.ReplaceAllString(uuid, "")
		s1 := uuid[0:8]
		s2 := uuid[8:16]
		s3 := uuid[16:24]
		s4 := uuid[24:32]
		// fmt.Println(s1, s2, s3, s4)
		n1, _ := strconv.ParseUint(s1, 16, 32)
		n2, _ := strconv.ParseUint(s2, 16, 32)
		n3, _ := strconv.ParseUint(s3, 16, 32)
		n4, _ := strconv.ParseUint(s4, 16, 32)
		// fmt.Println(n1, n2, n3, n4)
		e1 := encode(n1)
		e2 := encode(n2)
		e3 := encode(n3)
		e4 := encode(n4)
		return (e1 + "-" + e2 + "-" + e3 + "-" + e4)
	}
	panic("Invalid UUID string")
}

// ToUUID will convert given API Key to UUID
func ToUUID(apiKey string) string {
	if apiKey == "" {
		panic("The required parameter API Key is undefined")
	}
	if IsAPIKey(apiKey) {
		var re = regexp.MustCompile("-")
		apiKey = re.ReplaceAllString(apiKey, "")
		e1 := apiKey[0:7]
		e2 := apiKey[7:14]
		e3 := apiKey[14:21]
		e4 := apiKey[21:28]
		// fmt.Println(e1, e2, e3, e4)
		d1 := decode(e1)
		d2 := decode(e2)
		d3 := decode(e3)
		d4 := decode(e4)
		// fmt.Println(d1, d2, d3, d4)
		d2a := d2[0:4]
		d2b := d2[4:8]
		d3a := d3[0:4]
		d3b := d3[4:8]
		return (d1 + "-" + d2a + "-" + d2b + "-" + d3a + "-" + d3b + d4)
	}
	panic("Invalid API Key string")
}

// Check will verify both UUID and API Key, with either given
func CheckAPIKey(uuid, apiKey string) bool {
	if uuid == "" {
		panic("The required parameter UUID is undefined")
	}
	if apiKey == "" {
		panic("The required parameter API Key is undefined")
	}
	if IsUUID(uuid) && IsAPIKey(apiKey) {
		uuidCheck := ToUUID(apiKey)
		return uuidCheck == uuid
	}
	return false
}

// Create will create new UUID and API Key
func CreateUUIDAPIKey() string {
	uuid := UUID4().String()
	apiKey := ToAPIKey(uuid)
	pair := map[string]string{"uuid": uuid, "apiKey": apiKey}
	jsonPair, err := json.Marshal(pair)
	if err != nil {
		panic("Error creating a new pair of keys")
	}
	return string(jsonPair)
}
