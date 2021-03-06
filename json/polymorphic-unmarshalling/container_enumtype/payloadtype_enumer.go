// Code generated by "enumer -type=PayloadType -json -trimprefix=Type"; DO NOT EDIT.

package main

import (
	"encoding/json"
	"fmt"
)

const _PayloadTypeName = "EncryptedPayloadPublicPayload"

var _PayloadTypeIndex = [...]uint8{0, 16, 29}

const _PayloadTypeLowerName = "encryptedpayloadpublicpayload"

func (i PayloadType) String() string {
	if i < 0 || i >= PayloadType(len(_PayloadTypeIndex)-1) {
		return fmt.Sprintf("PayloadType(%d)", i)
	}
	return _PayloadTypeName[_PayloadTypeIndex[i]:_PayloadTypeIndex[i+1]]
}

var _PayloadTypeValues = []PayloadType{0, 1}

var _PayloadTypeNameToValueMap = map[string]PayloadType{
	_PayloadTypeName[0:16]:       0,
	_PayloadTypeLowerName[0:16]:  0,
	_PayloadTypeName[16:29]:      1,
	_PayloadTypeLowerName[16:29]: 1,
}

var _PayloadTypeNames = []string{
	_PayloadTypeName[0:16],
	_PayloadTypeName[16:29],
}

// PayloadTypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func PayloadTypeString(s string) (PayloadType, error) {
	if val, ok := _PayloadTypeNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to PayloadType values", s)
}

// PayloadTypeValues returns all values of the enum
func PayloadTypeValues() []PayloadType {
	return _PayloadTypeValues
}

// PayloadTypeStrings returns a slice of all String values of the enum
func PayloadTypeStrings() []string {
	strs := make([]string, len(_PayloadTypeNames))
	copy(strs, _PayloadTypeNames)
	return strs
}

// IsAPayloadType returns "true" if the value is listed in the enum definition. "false" otherwise
func (i PayloadType) IsAPayloadType() bool {
	for _, v := range _PayloadTypeValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for PayloadType
func (i PayloadType) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for PayloadType
func (i *PayloadType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("PayloadType should be a string, got %s", data)
	}

	var err error
	*i, err = PayloadTypeString(s)
	return err
}
