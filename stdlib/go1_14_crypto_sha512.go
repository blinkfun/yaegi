// Code generated by 'github.com/containous/yaegi/extract crypto/sha512'. DO NOT EDIT.

// +build go1.14,!go1.15

package stdlib

import (
	"crypto/sha512"
	"go/constant"
	"go/token"
	"reflect"
)

func init() {
	Symbols["crypto/sha512"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"BlockSize":  reflect.ValueOf(constant.MakeFromLiteral("128", token.INT, 0)),
		"New":        reflect.ValueOf(sha512.New),
		"New384":     reflect.ValueOf(sha512.New384),
		"New512_224": reflect.ValueOf(sha512.New512_224),
		"New512_256": reflect.ValueOf(sha512.New512_256),
		"Size":       reflect.ValueOf(constant.MakeFromLiteral("64", token.INT, 0)),
		"Size224":    reflect.ValueOf(constant.MakeFromLiteral("28", token.INT, 0)),
		"Size256":    reflect.ValueOf(constant.MakeFromLiteral("32", token.INT, 0)),
		"Size384":    reflect.ValueOf(constant.MakeFromLiteral("48", token.INT, 0)),
		"Sum384":     reflect.ValueOf(sha512.Sum384),
		"Sum512":     reflect.ValueOf(sha512.Sum512),
		"Sum512_224": reflect.ValueOf(sha512.Sum512_224),
		"Sum512_256": reflect.ValueOf(sha512.Sum512_256),
	}
}
