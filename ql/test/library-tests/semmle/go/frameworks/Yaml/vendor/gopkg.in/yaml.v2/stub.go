// Code generated by depstubber. DO NOT EDIT.
// This is a simple stub for gopkg.in/yaml.v2, strictly for use in testing.

// See the LICENSE file for information about the licensing of the original library.
// Source: gopkg.in/yaml.v2 (exports: ; functions: Marshal,NewDecoder,NewEncoder,Unmarshal,UnmarshalStrict)

// Package yaml is a stub of gopkg.in/yaml.v2, generated by depstubber.
package yaml

import (
	io "io"
)

type Decoder struct{}

func (_ *Decoder) Decode(_ interface{}) error {
	return nil
}

func (_ *Decoder) SetStrict(_ bool) {}

type Encoder struct{}

func (_ *Encoder) Close() error {
	return nil
}

func (_ *Encoder) Encode(_ interface{}) error {
	return nil
}

func Marshal(_ interface{}) ([]byte, error) {
	return nil, nil
}

func NewDecoder(_ io.Reader) *Decoder {
	return nil
}

func NewEncoder(_ io.Writer) *Encoder {
	return nil
}

func Unmarshal(_ []byte, _ interface{}) error {
	return nil
}

func UnmarshalStrict(_ []byte, _ interface{}) error {
	return nil
}
