// Code generated by 'github.com/containous/yaegi/extract archive/tar'. DO NOT EDIT.

// +build go1.13,!go1.14

package stdlib

import (
	"archive/tar"
	"go/constant"
	"go/token"
	"reflect"
)

func init() {
	Symbols["archive/tar"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"ErrFieldTooLong":    reflect.ValueOf(&tar.ErrFieldTooLong).Elem(),
		"ErrHeader":          reflect.ValueOf(&tar.ErrHeader).Elem(),
		"ErrWriteAfterClose": reflect.ValueOf(&tar.ErrWriteAfterClose).Elem(),
		"ErrWriteTooLong":    reflect.ValueOf(&tar.ErrWriteTooLong).Elem(),
		"FileInfoHeader":     reflect.ValueOf(tar.FileInfoHeader),
		"FormatGNU":          reflect.ValueOf(tar.FormatGNU),
		"FormatPAX":          reflect.ValueOf(tar.FormatPAX),
		"FormatUSTAR":        reflect.ValueOf(tar.FormatUSTAR),
		"FormatUnknown":      reflect.ValueOf(tar.FormatUnknown),
		"NewReader":          reflect.ValueOf(tar.NewReader),
		"NewWriter":          reflect.ValueOf(tar.NewWriter),
		"TypeBlock":          reflect.ValueOf(constant.MakeFromLiteral("52", token.INT, 0)),
		"TypeChar":           reflect.ValueOf(constant.MakeFromLiteral("51", token.INT, 0)),
		"TypeCont":           reflect.ValueOf(constant.MakeFromLiteral("55", token.INT, 0)),
		"TypeDir":            reflect.ValueOf(constant.MakeFromLiteral("53", token.INT, 0)),
		"TypeFifo":           reflect.ValueOf(constant.MakeFromLiteral("54", token.INT, 0)),
		"TypeGNULongLink":    reflect.ValueOf(constant.MakeFromLiteral("75", token.INT, 0)),
		"TypeGNULongName":    reflect.ValueOf(constant.MakeFromLiteral("76", token.INT, 0)),
		"TypeGNUSparse":      reflect.ValueOf(constant.MakeFromLiteral("83", token.INT, 0)),
		"TypeLink":           reflect.ValueOf(constant.MakeFromLiteral("49", token.INT, 0)),
		"TypeReg":            reflect.ValueOf(constant.MakeFromLiteral("48", token.INT, 0)),
		"TypeRegA":           reflect.ValueOf(constant.MakeFromLiteral("0", token.INT, 0)),
		"TypeSymlink":        reflect.ValueOf(constant.MakeFromLiteral("50", token.INT, 0)),
		"TypeXGlobalHeader":  reflect.ValueOf(constant.MakeFromLiteral("103", token.INT, 0)),
		"TypeXHeader":        reflect.ValueOf(constant.MakeFromLiteral("120", token.INT, 0)),

		// type definitions
		"Format": reflect.ValueOf((*tar.Format)(nil)),
		"Header": reflect.ValueOf((*tar.Header)(nil)),
		"Reader": reflect.ValueOf((*tar.Reader)(nil)),
		"Writer": reflect.ValueOf((*tar.Writer)(nil)),
	}
}
