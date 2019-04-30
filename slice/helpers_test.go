//go:generate go run gen.go
package slice

import (
	"errors"
	"math/cmplx"
	"strings"
)

func getInput(i string) interface{} {
	switch i {
	case "bool":
		return []bool{true, false}
	case "byte":
		return []byte{1, 2}
	case "complex128":
		return []complex128{complex(2, 1), complex(-1, 1)}
	case "complex64":
		return []complex64{complex64(2 + 1i), complex64(-1 + 1i)}
	case "error":
		return []error{errors.New("wrong!"), errors.New("bad!")}
	case "float32":
		return []float32{float32(2), float32(4)}
	case "float64":
		return []float64{float64(3), float64(9)}
	case "int":
		return []int{4, 16}
	case "int16":
		return []int16{1, 2}
	case "int32":
		return []int32{1, 2}
	case "int64":
		return []int64{3, 5}
	case "int8":
		return []int8{5, 3}
	case "rune":
		return []rune{'e', 'i'}
	case "string":
		return []string{"hello", "world!"}
	case "uint":
		return []uint{5, 8}
	case "uint16":
		return []uint16{5, 8}
	case "uint32":
		return []uint32{5, 8}
	case "uint64":
		return []uint64{5, 8}
	case "uint8":
		return []uint8{5, 8}
	case "uintptr":
		return []uintptr{1, 2}
	}

	return nil
}

func getExpect(i string) interface{} {
	switch i {
	case "bool":
		return []bool{false, true}
	case "byte":
		return []byte{2, 4}
	case "complex128":
		return []complex128{complex(2, -1), complex(-1, -1)}
	case "complex64":
		return []complex64{complex64(2 + 1i), complex64(-1 + 1i)}
	case "error":
		return []error{errors.New("too wrong!"), errors.New("too bad!")}
	case "float32":
		return []float32{float32(4), float32(8)}
	case "float64":
		return []float64{float64(9), float64(27)}
	case "int":
		return []int{2, 8}
	case "int16":
		return []int16{1, 2}
	case "int32":
		return []int32{1, 4}
	case "int64":
		return []int64{9, 25}
	case "int8":
		return []int8{7, 5}
	case "rune":
		return []rune{'a', 'a'}
	case "string":
		return []string{"Hello", "World!"}
	case "uint":
		return []uint{25, 64}
	case "uint16":
		return []uint16{25, 64}
	case "uint32":
		return []uint32{25, 64}
	case "uint64":
		return []uint64{25, 64}
	case "uint8":
		return []uint8{1, 0}
	case "uintptr":
		return []uintptr{1, 2}
	}

	return nil
}

// bool
func boolFunction(b bool) bool {
	return !b
}

// byte
func byteFunction(b byte) byte {
	return b * 2
}

// complex128
func complex128Function(c complex128) complex128 {
	return cmplx.Conj(c)
}

// complex64
func complex64Function(c complex64) complex64 {
	return c
}

// error
func errorFunction(e error) error {
	return errors.New("too " + e.Error())
}

// float32
func float32Function(f float32) float32 {
	return f * 2
}

// float64
func float64Function(f float64) float64 {
	return f * 3
}

// int
func intFunction(i int) int {
	return i / 2
}

// int16
func int16Function(i int16) int16 {
	return i * 1
}

// int32
func int32Function(i int32) int32 {
	return i * i
}

// int64
func int64Function(i int64) int64 {
	return i * i
}

// int8
func int8Function(i int8) int8 {
	return i + 2
}

// rune
func runeFunction(r rune) rune {
	return 'a'
}

// string
func stringFunction(i string) string {
	return strings.Title(i)
}

// uint
func uintFunction(u uint) uint {
	return u * u
}

// uint16
func uint16Function(u uint16) uint16 {
	return u * u
}

// uint32
func uint32Function(u uint32) uint32 {
	return u * u
}

// uint64
func uint64Function(u uint64) uint64 {
	return u * u
}

// uint8
func uint8Function(u uint8) uint8 {
	return u & 1
}

// uintptr
func uintptrFunction(u uintptr) uintptr {
	return u
}
