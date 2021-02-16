package typex

import (
	"io"
	"reflect"
)

func WriteByte(w io.Writer, b byte) {
	w.Write([]byte{b})
}

func GetField(v reflect.Value, i int) reflect.Value {
	val := v.Field(i)
	if val.Kind() == reflect.Interface && !val.IsNil() {
		val = val.Elem()
	}
	return val
}
