// Code generated by 'github.com/traefik/yaegi/extract crypto/rc4'. DO NOT EDIT.

// +build go1.15,!go1.16

package stdlib

import (
	"crypto/rc4"
	"reflect"
)

func init() {
	Symbols["crypto/rc4"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"NewCipher": reflect.ValueOf(rc4.NewCipher),

		// type definitions
		"Cipher":       reflect.ValueOf((*rc4.Cipher)(nil)),
		"KeySizeError": reflect.ValueOf((*rc4.KeySizeError)(nil)),
	}
}
