// this file was generated by gomacro command: import _b "encoding"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"encoding"
)

// reflection: allow interpreted code to import "encoding"
func init() {
	Packages["encoding"] = Package{
	Types: map[string]Type{
		"BinaryMarshaler":	TypeOf((*encoding.BinaryMarshaler)(nil)).Elem(),
		"BinaryUnmarshaler":	TypeOf((*encoding.BinaryUnmarshaler)(nil)).Elem(),
		"TextMarshaler":	TypeOf((*encoding.TextMarshaler)(nil)).Elem(),
		"TextUnmarshaler":	TypeOf((*encoding.TextUnmarshaler)(nil)).Elem(),
	}, Proxies: map[string]Type{
		"BinaryMarshaler":	TypeOf((*BinaryMarshaler_encoding)(nil)).Elem(),
		"BinaryUnmarshaler":	TypeOf((*BinaryUnmarshaler_encoding)(nil)).Elem(),
		"TextMarshaler":	TypeOf((*TextMarshaler_encoding)(nil)).Elem(),
		"TextUnmarshaler":	TypeOf((*TextUnmarshaler_encoding)(nil)).Elem(),
	}, 
	}
}

// --------------- proxy for encoding.BinaryMarshaler ---------------
type BinaryMarshaler_encoding struct {
	Object	interface{}
	MarshalBinary_	func(interface{}) (data []byte, err error)
}
func (Proxy *BinaryMarshaler_encoding) MarshalBinary() (data []byte, err error) {
	return Proxy.MarshalBinary_(Proxy.Object)
}

// --------------- proxy for encoding.BinaryUnmarshaler ---------------
type BinaryUnmarshaler_encoding struct {
	Object	interface{}
	UnmarshalBinary_	func(_proxy_obj_ interface{}, data []byte) error
}
func (Proxy *BinaryUnmarshaler_encoding) UnmarshalBinary(data []byte) error {
	return Proxy.UnmarshalBinary_(Proxy.Object, data)
}

// --------------- proxy for encoding.TextMarshaler ---------------
type TextMarshaler_encoding struct {
	Object	interface{}
	MarshalText_	func(interface{}) (text []byte, err error)
}
func (Proxy *TextMarshaler_encoding) MarshalText() (text []byte, err error) {
	return Proxy.MarshalText_(Proxy.Object)
}

// --------------- proxy for encoding.TextUnmarshaler ---------------
type TextUnmarshaler_encoding struct {
	Object	interface{}
	UnmarshalText_	func(_proxy_obj_ interface{}, text []byte) error
}
func (Proxy *TextUnmarshaler_encoding) UnmarshalText(text []byte) error {
	return Proxy.UnmarshalText_(Proxy.Object, text)
}
