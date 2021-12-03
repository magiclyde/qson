/**
 * Created by GoLand.
 * @author: clyde
 * @date: 2021/12/3 下午4:35
 * @note:
 */

package qson

import (
	"bytes"
	"reflect"
	"unsafe"
)

// Concat join string
func Concat(values ...string) string {
	var buffer bytes.Buffer
	for _, s := range values {
		buffer.WriteString(s)
	}
	return buffer.String()
}

// S2B converts string to a byte slice without memory allocation
func S2B(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

// B2S converts a byte slice to string without memory allocation.
func B2S(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
