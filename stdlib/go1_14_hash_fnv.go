// Code generated by 'github.com/containous/yaegi/extract hash/fnv'. DO NOT EDIT.

// +build go1.14,!go1.15

package stdlib

import (
	"hash/fnv"
	"reflect"
)

func init() {
	Symbols["hash/fnv"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"New128":  reflect.ValueOf(fnv.New128),
		"New128a": reflect.ValueOf(fnv.New128a),
		"New32":   reflect.ValueOf(fnv.New32),
		"New32a":  reflect.ValueOf(fnv.New32a),
		"New64":   reflect.ValueOf(fnv.New64),
		"New64a":  reflect.ValueOf(fnv.New64a),
	}
}
