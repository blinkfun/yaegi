// Code generated by 'github.com/containous/yaegi/extract crypto/ecdsa'. DO NOT EDIT.

// +build go1.13,!go1.14

package stdlib

import (
	"crypto/ecdsa"
	"reflect"
)

func init() {
	Symbols["crypto/ecdsa"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"GenerateKey": reflect.ValueOf(ecdsa.GenerateKey),
		"Sign":        reflect.ValueOf(ecdsa.Sign),
		"Verify":      reflect.ValueOf(ecdsa.Verify),

		// type definitions
		"PrivateKey": reflect.ValueOf((*ecdsa.PrivateKey)(nil)),
		"PublicKey":  reflect.ValueOf((*ecdsa.PublicKey)(nil)),
	}
}
