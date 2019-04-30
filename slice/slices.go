// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at
// 2019-04-30 18:12:06.194051 -0500 CDT m=+0.001776842
package slice

import "fmt"

func Includes(inputInterface interface{}, eInterface interface{}, indexes ...int) (bool, error) {
	switch e := eInterface.(type) {
	
	case bool:
		input, ok := inputInterface.([]bool)
		if !ok {
			e := fmt.Sprintf("%s must be from bool", inputInterface)
			return false, newError(e)
		}
		return boolIncludes(input, e, indexes...), nil
	
	case complex128:
		input, ok := inputInterface.([]complex128)
		if !ok {
			e := fmt.Sprintf("%s must be from complex128", inputInterface)
			return false, newError(e)
		}
		return complex128Includes(input, e, indexes...), nil
	
	case complex64:
		input, ok := inputInterface.([]complex64)
		if !ok {
			e := fmt.Sprintf("%s must be from complex64", inputInterface)
			return false, newError(e)
		}
		return complex64Includes(input, e, indexes...), nil
	
	case error:
		input, ok := inputInterface.([]error)
		if !ok {
			e := fmt.Sprintf("%s must be from error", inputInterface)
			return false, newError(e)
		}
		return errorIncludes(input, e, indexes...), nil
	
	case float32:
		input, ok := inputInterface.([]float32)
		if !ok {
			e := fmt.Sprintf("%s must be from float32", inputInterface)
			return false, newError(e)
		}
		return float32Includes(input, e, indexes...), nil
	
	case float64:
		input, ok := inputInterface.([]float64)
		if !ok {
			e := fmt.Sprintf("%s must be from float64", inputInterface)
			return false, newError(e)
		}
		return float64Includes(input, e, indexes...), nil
	
	case int:
		input, ok := inputInterface.([]int)
		if !ok {
			e := fmt.Sprintf("%s must be from int", inputInterface)
			return false, newError(e)
		}
		return intIncludes(input, e, indexes...), nil
	
	case int16:
		input, ok := inputInterface.([]int16)
		if !ok {
			e := fmt.Sprintf("%s must be from int16", inputInterface)
			return false, newError(e)
		}
		return int16Includes(input, e, indexes...), nil
	
	case int32:
		input, ok := inputInterface.([]int32)
		if !ok {
			e := fmt.Sprintf("%s must be from int32", inputInterface)
			return false, newError(e)
		}
		return int32Includes(input, e, indexes...), nil
	
	case int64:
		input, ok := inputInterface.([]int64)
		if !ok {
			e := fmt.Sprintf("%s must be from int64", inputInterface)
			return false, newError(e)
		}
		return int64Includes(input, e, indexes...), nil
	
	case int8:
		input, ok := inputInterface.([]int8)
		if !ok {
			e := fmt.Sprintf("%s must be from int8", inputInterface)
			return false, newError(e)
		}
		return int8Includes(input, e, indexes...), nil
	
	case string:
		input, ok := inputInterface.([]string)
		if !ok {
			e := fmt.Sprintf("%s must be from string", inputInterface)
			return false, newError(e)
		}
		return stringIncludes(input, e, indexes...), nil
	
	case uint:
		input, ok := inputInterface.([]uint)
		if !ok {
			e := fmt.Sprintf("%s must be from uint", inputInterface)
			return false, newError(e)
		}
		return uintIncludes(input, e, indexes...), nil
	
	case uint16:
		input, ok := inputInterface.([]uint16)
		if !ok {
			e := fmt.Sprintf("%s must be from uint16", inputInterface)
			return false, newError(e)
		}
		return uint16Includes(input, e, indexes...), nil
	
	case uint32:
		input, ok := inputInterface.([]uint32)
		if !ok {
			e := fmt.Sprintf("%s must be from uint32", inputInterface)
			return false, newError(e)
		}
		return uint32Includes(input, e, indexes...), nil
	
	case uint64:
		input, ok := inputInterface.([]uint64)
		if !ok {
			e := fmt.Sprintf("%s must be from uint64", inputInterface)
			return false, newError(e)
		}
		return uint64Includes(input, e, indexes...), nil
	
	case uint8:
		input, ok := inputInterface.([]uint8)
		if !ok {
			e := fmt.Sprintf("%s must be from uint8", inputInterface)
			return false, newError(e)
		}
		return uint8Includes(input, e, indexes...), nil
	
	case uintptr:
		input, ok := inputInterface.([]uintptr)
		if !ok {
			e := fmt.Sprintf("%s must be from uintptr", inputInterface)
			return false, newError(e)
		}
		return uintptrIncludes(input, e, indexes...), nil
	
	}

	return false, nil
}


// Bool
func BoolForEach(input []bool, function interface{}) error {
	if !(isFunc(function)) {
		e := fmt.Sprintf("parameter (%s) is not a function", function)
		return newError(e)
	}

	switch function.(type) {
	case func(bool):
		fn := function.(func(bool))
		for _, element := range input {
			fn(element)
		}

	case func(bool, int):
		fn := function.(func(bool, int))
		for i, element := range input {
			fn(element, i)
		}

	default:
		e := fmt.Sprintf("function (%s) must have a firm: '%s' or '%s' and got: '%s'",
			getFuncName(function), "func(bool)", "func(bool, int)", getFuncFirm(function))

		return newError(e)
	}

	return nil
}

func boolIncludes(input []bool, e bool, indexes ...int) bool {
	index := 0

	// We only care for the first index on indexes because the lack of optional
	// parameters in Go
	if len(indexes) > 0 {
		index = indexes[0]
	}

	switch {
	case index >= len(input):
		return false
	case index < 0 && len(input) + index <= -1 * len(input):
		index = 0
	case index < 0:
		index = len(input) + index
	}

	for i := index; i < len(input); i++ {
		if input[i] == e {
			return true
		}
	}

	return false
}

// Complex128
func Complex128ForEach(input []complex128, function interface{}) error {
	if !(isFunc(function)) {
		e := fmt.Sprintf("parameter (%s) is not a function", function)
		return newError(e)
	}

	switch function.(type) {
	case func(complex128):
		fn := function.(func(complex128))
		for _, element := range input {
			fn(element)
		}

	case func(complex128, int):
		fn := function.(func(complex128, int))
		for i, element := range input {
			fn(element, i)
		}

	default:
		e := fmt.Sprintf("function (%s) must have a firm: '%s' or '%s' and got: '%s'",
			getFuncName(function), "func(complex128)", "func(complex128, int)", getFuncFirm(function))

		return newError(e)
	}

	return nil
}

func complex128Includes(input []complex128, e complex128, indexes ...int) bool {
	index := 0

	// We only care for the first index on indexes because the lack of optional
	// parameters in Go
	if len(indexes) > 0 {
		index = indexes[0]
	}

	switch {
	case index >= len(input):
		return false
	case index < 0 && len(input) + index <= -1 * len(input):
		index = 0
	case index < 0:
		index = len(input) + index
	}

	for i := index; i < len(input); i++ {
		if input[i] == e {
			return true
		}
	}

	return false
}

// Complex64
func Complex64ForEach(input []complex64, function interface{}) error {
	if !(isFunc(function)) {
		e := fmt.Sprintf("parameter (%s) is not a function", function)
		return newError(e)
	}

	switch function.(type) {
	case func(complex64):
		fn := function.(func(complex64))
		for _, element := range input {
			fn(element)
		}

	case func(complex64, int):
		fn := function.(func(complex64, int))
		for i, element := range input {
			fn(element, i)
		}

	default:
		e := fmt.Sprintf("function (%s) must have a firm: '%s' or '%s' and got: '%s'",
			getFuncName(function), "func(complex64)", "func(complex64, int)", getFuncFirm(function))

		return newError(e)
	}

	return nil
}

func complex64Includes(input []complex64, e complex64, indexes ...int) bool {
	index := 0

	// We only care for the first index on indexes because the lack of optional
	// parameters in Go
	if len(indexes) > 0 {
		index = indexes[0]
	}

	switch {
	case index >= len(input):
		return false
	case index < 0 && len(input) + index <= -1 * len(input):
		index = 0
	case index < 0:
		index = len(input) + index
	}

	for i := index; i < len(input); i++ {
		if input[i] == e {
			return true
		}
	}

	return false
}

// Error
func ErrorForEach(input []error, function interface{}) error {
	if !(isFunc(function)) {
		e := fmt.Sprintf("parameter (%s) is not a function", function)
		return newError(e)
	}

	switch function.(type) {
	case func(error):
		fn := function.(func(error))
		for _, element := range input {
			fn(element)
		}

	case func(error, int):
		fn := function.(func(error, int))
		for i, element := range input {
			fn(element, i)
		}

	default:
		e := fmt.Sprintf("function (%s) must have a firm: '%s' or '%s' and got: '%s'",
			getFuncName(function), "func(error)", "func(error, int)", getFuncFirm(function))

		return newError(e)
	}

	return nil
}

func errorIncludes(input []error, e error, indexes ...int) bool {
	index := 0

	// We only care for the first index on indexes because the lack of optional
	// parameters in Go
	if len(indexes) > 0 {
		index = indexes[0]
	}

	switch {
	case index >= len(input):
		return false
	case index < 0 && len(input) + index <= -1 * len(input):
		index = 0
	case index < 0:
		index = len(input) + index
	}

	for i := index; i < len(input); i++ {
		if input[i] == e {
			return true
		}
	}

	return false
}

// Float32
func Float32ForEach(input []float32, function interface{}) error {
	if !(isFunc(function)) {
		e := fmt.Sprintf("parameter (%s) is not a function", function)
		return newError(e)
	}

	switch function.(type) {
	case func(float32):
		fn := function.(func(float32))
		for _, element := range input {
			fn(element)
		}

	case func(float32, int):
		fn := function.(func(float32, int))
		for i, element := range input {
			fn(element, i)
		}

	default:
		e := fmt.Sprintf("function (%s) must have a firm: '%s' or '%s' and got: '%s'",
			getFuncName(function), "func(float32)", "func(float32, int)", getFuncFirm(function))

		return newError(e)
	}

	return nil
}

func float32Includes(input []float32, e float32, indexes ...int) bool {
	index := 0

	// We only care for the first index on indexes because the lack of optional
	// parameters in Go
	if len(indexes) > 0 {
		index = indexes[0]
	}

	switch {
	case index >= len(input):
		return false
	case index < 0 && len(input) + index <= -1 * len(input):
		index = 0
	case index < 0:
		index = len(input) + index
	}

	for i := index; i < len(input); i++ {
		if input[i] == e {
			return true
		}
	}

	return false
}

// Float64
func Float64ForEach(input []float64, function interface{}) error {
	if !(isFunc(function)) {
		e := fmt.Sprintf("parameter (%s) is not a function", function)
		return newError(e)
	}

	switch function.(type) {
	case func(float64):
		fn := function.(func(float64))
		for _, element := range input {
			fn(element)
		}

	case func(float64, int):
		fn := function.(func(float64, int))
		for i, element := range input {
			fn(element, i)
		}

	default:
		e := fmt.Sprintf("function (%s) must have a firm: '%s' or '%s' and got: '%s'",
			getFuncName(function), "func(float64)", "func(float64, int)", getFuncFirm(function))

		return newError(e)
	}

	return nil
}

func float64Includes(input []float64, e float64, indexes ...int) bool {
	index := 0

	// We only care for the first index on indexes because the lack of optional
	// parameters in Go
	if len(indexes) > 0 {
		index = indexes[0]
	}

	switch {
	case index >= len(input):
		return false
	case index < 0 && len(input) + index <= -1 * len(input):
		index = 0
	case index < 0:
		index = len(input) + index
	}

	for i := index; i < len(input); i++ {
		if input[i] == e {
			return true
		}
	}

	return false
}

// Int
func IntForEach(input []int, function interface{}) error {
	if !(isFunc(function)) {
		e := fmt.Sprintf("parameter (%s) is not a function", function)
		return newError(e)
	}

	switch function.(type) {
	case func(int):
		fn := function.(func(int))
		for _, element := range input {
			fn(element)
		}

	case func(int, int):
		fn := function.(func(int, int))
		for i, element := range input {
			fn(element, i)
		}

	default:
		e := fmt.Sprintf("function (%s) must have a firm: '%s' or '%s' and got: '%s'",
			getFuncName(function), "func(int)", "func(int, int)", getFuncFirm(function))

		return newError(e)
	}

	return nil
}

func intIncludes(input []int, e int, indexes ...int) bool {
	index := 0

	// We only care for the first index on indexes because the lack of optional
	// parameters in Go
	if len(indexes) > 0 {
		index = indexes[0]
	}

	switch {
	case index >= len(input):
		return false
	case index < 0 && len(input) + index <= -1 * len(input):
		index = 0
	case index < 0:
		index = len(input) + index
	}

	for i := index; i < len(input); i++ {
		if input[i] == e {
			return true
		}
	}

	return false
}

// Int16
func Int16ForEach(input []int16, function interface{}) error {
	if !(isFunc(function)) {
		e := fmt.Sprintf("parameter (%s) is not a function", function)
		return newError(e)
	}

	switch function.(type) {
	case func(int16):
		fn := function.(func(int16))
		for _, element := range input {
			fn(element)
		}

	case func(int16, int):
		fn := function.(func(int16, int))
		for i, element := range input {
			fn(element, i)
		}

	default:
		e := fmt.Sprintf("function (%s) must have a firm: '%s' or '%s' and got: '%s'",
			getFuncName(function), "func(int16)", "func(int16, int)", getFuncFirm(function))

		return newError(e)
	}

	return nil
}

func int16Includes(input []int16, e int16, indexes ...int) bool {
	index := 0

	// We only care for the first index on indexes because the lack of optional
	// parameters in Go
	if len(indexes) > 0 {
		index = indexes[0]
	}

	switch {
	case index >= len(input):
		return false
	case index < 0 && len(input) + index <= -1 * len(input):
		index = 0
	case index < 0:
		index = len(input) + index
	}

	for i := index; i < len(input); i++ {
		if input[i] == e {
			return true
		}
	}

	return false
}

// Int32
func Int32ForEach(input []int32, function interface{}) error {
	if !(isFunc(function)) {
		e := fmt.Sprintf("parameter (%s) is not a function", function)
		return newError(e)
	}

	switch function.(type) {
	case func(int32):
		fn := function.(func(int32))
		for _, element := range input {
			fn(element)
		}

	case func(int32, int):
		fn := function.(func(int32, int))
		for i, element := range input {
			fn(element, i)
		}

	default:
		e := fmt.Sprintf("function (%s) must have a firm: '%s' or '%s' and got: '%s'",
			getFuncName(function), "func(int32)", "func(int32, int)", getFuncFirm(function))

		return newError(e)
	}

	return nil
}

func int32Includes(input []int32, e int32, indexes ...int) bool {
	index := 0

	// We only care for the first index on indexes because the lack of optional
	// parameters in Go
	if len(indexes) > 0 {
		index = indexes[0]
	}

	switch {
	case index >= len(input):
		return false
	case index < 0 && len(input) + index <= -1 * len(input):
		index = 0
	case index < 0:
		index = len(input) + index
	}

	for i := index; i < len(input); i++ {
		if input[i] == e {
			return true
		}
	}

	return false
}

// Int64
func Int64ForEach(input []int64, function interface{}) error {
	if !(isFunc(function)) {
		e := fmt.Sprintf("parameter (%s) is not a function", function)
		return newError(e)
	}

	switch function.(type) {
	case func(int64):
		fn := function.(func(int64))
		for _, element := range input {
			fn(element)
		}

	case func(int64, int):
		fn := function.(func(int64, int))
		for i, element := range input {
			fn(element, i)
		}

	default:
		e := fmt.Sprintf("function (%s) must have a firm: '%s' or '%s' and got: '%s'",
			getFuncName(function), "func(int64)", "func(int64, int)", getFuncFirm(function))

		return newError(e)
	}

	return nil
}

func int64Includes(input []int64, e int64, indexes ...int) bool {
	index := 0

	// We only care for the first index on indexes because the lack of optional
	// parameters in Go
	if len(indexes) > 0 {
		index = indexes[0]
	}

	switch {
	case index >= len(input):
		return false
	case index < 0 && len(input) + index <= -1 * len(input):
		index = 0
	case index < 0:
		index = len(input) + index
	}

	for i := index; i < len(input); i++ {
		if input[i] == e {
			return true
		}
	}

	return false
}

// Int8
func Int8ForEach(input []int8, function interface{}) error {
	if !(isFunc(function)) {
		e := fmt.Sprintf("parameter (%s) is not a function", function)
		return newError(e)
	}

	switch function.(type) {
	case func(int8):
		fn := function.(func(int8))
		for _, element := range input {
			fn(element)
		}

	case func(int8, int):
		fn := function.(func(int8, int))
		for i, element := range input {
			fn(element, i)
		}

	default:
		e := fmt.Sprintf("function (%s) must have a firm: '%s' or '%s' and got: '%s'",
			getFuncName(function), "func(int8)", "func(int8, int)", getFuncFirm(function))

		return newError(e)
	}

	return nil
}

func int8Includes(input []int8, e int8, indexes ...int) bool {
	index := 0

	// We only care for the first index on indexes because the lack of optional
	// parameters in Go
	if len(indexes) > 0 {
		index = indexes[0]
	}

	switch {
	case index >= len(input):
		return false
	case index < 0 && len(input) + index <= -1 * len(input):
		index = 0
	case index < 0:
		index = len(input) + index
	}

	for i := index; i < len(input); i++ {
		if input[i] == e {
			return true
		}
	}

	return false
}

// String
func StringForEach(input []string, function interface{}) error {
	if !(isFunc(function)) {
		e := fmt.Sprintf("parameter (%s) is not a function", function)
		return newError(e)
	}

	switch function.(type) {
	case func(string):
		fn := function.(func(string))
		for _, element := range input {
			fn(element)
		}

	case func(string, int):
		fn := function.(func(string, int))
		for i, element := range input {
			fn(element, i)
		}

	default:
		e := fmt.Sprintf("function (%s) must have a firm: '%s' or '%s' and got: '%s'",
			getFuncName(function), "func(string)", "func(string, int)", getFuncFirm(function))

		return newError(e)
	}

	return nil
}

func stringIncludes(input []string, e string, indexes ...int) bool {
	index := 0

	// We only care for the first index on indexes because the lack of optional
	// parameters in Go
	if len(indexes) > 0 {
		index = indexes[0]
	}

	switch {
	case index >= len(input):
		return false
	case index < 0 && len(input) + index <= -1 * len(input):
		index = 0
	case index < 0:
		index = len(input) + index
	}

	for i := index; i < len(input); i++ {
		if input[i] == e {
			return true
		}
	}

	return false
}

// Uint
func UintForEach(input []uint, function interface{}) error {
	if !(isFunc(function)) {
		e := fmt.Sprintf("parameter (%s) is not a function", function)
		return newError(e)
	}

	switch function.(type) {
	case func(uint):
		fn := function.(func(uint))
		for _, element := range input {
			fn(element)
		}

	case func(uint, int):
		fn := function.(func(uint, int))
		for i, element := range input {
			fn(element, i)
		}

	default:
		e := fmt.Sprintf("function (%s) must have a firm: '%s' or '%s' and got: '%s'",
			getFuncName(function), "func(uint)", "func(uint, int)", getFuncFirm(function))

		return newError(e)
	}

	return nil
}

func uintIncludes(input []uint, e uint, indexes ...int) bool {
	index := 0

	// We only care for the first index on indexes because the lack of optional
	// parameters in Go
	if len(indexes) > 0 {
		index = indexes[0]
	}

	switch {
	case index >= len(input):
		return false
	case index < 0 && len(input) + index <= -1 * len(input):
		index = 0
	case index < 0:
		index = len(input) + index
	}

	for i := index; i < len(input); i++ {
		if input[i] == e {
			return true
		}
	}

	return false
}

// Uint16
func Uint16ForEach(input []uint16, function interface{}) error {
	if !(isFunc(function)) {
		e := fmt.Sprintf("parameter (%s) is not a function", function)
		return newError(e)
	}

	switch function.(type) {
	case func(uint16):
		fn := function.(func(uint16))
		for _, element := range input {
			fn(element)
		}

	case func(uint16, int):
		fn := function.(func(uint16, int))
		for i, element := range input {
			fn(element, i)
		}

	default:
		e := fmt.Sprintf("function (%s) must have a firm: '%s' or '%s' and got: '%s'",
			getFuncName(function), "func(uint16)", "func(uint16, int)", getFuncFirm(function))

		return newError(e)
	}

	return nil
}

func uint16Includes(input []uint16, e uint16, indexes ...int) bool {
	index := 0

	// We only care for the first index on indexes because the lack of optional
	// parameters in Go
	if len(indexes) > 0 {
		index = indexes[0]
	}

	switch {
	case index >= len(input):
		return false
	case index < 0 && len(input) + index <= -1 * len(input):
		index = 0
	case index < 0:
		index = len(input) + index
	}

	for i := index; i < len(input); i++ {
		if input[i] == e {
			return true
		}
	}

	return false
}

// Uint32
func Uint32ForEach(input []uint32, function interface{}) error {
	if !(isFunc(function)) {
		e := fmt.Sprintf("parameter (%s) is not a function", function)
		return newError(e)
	}

	switch function.(type) {
	case func(uint32):
		fn := function.(func(uint32))
		for _, element := range input {
			fn(element)
		}

	case func(uint32, int):
		fn := function.(func(uint32, int))
		for i, element := range input {
			fn(element, i)
		}

	default:
		e := fmt.Sprintf("function (%s) must have a firm: '%s' or '%s' and got: '%s'",
			getFuncName(function), "func(uint32)", "func(uint32, int)", getFuncFirm(function))

		return newError(e)
	}

	return nil
}

func uint32Includes(input []uint32, e uint32, indexes ...int) bool {
	index := 0

	// We only care for the first index on indexes because the lack of optional
	// parameters in Go
	if len(indexes) > 0 {
		index = indexes[0]
	}

	switch {
	case index >= len(input):
		return false
	case index < 0 && len(input) + index <= -1 * len(input):
		index = 0
	case index < 0:
		index = len(input) + index
	}

	for i := index; i < len(input); i++ {
		if input[i] == e {
			return true
		}
	}

	return false
}

// Uint64
func Uint64ForEach(input []uint64, function interface{}) error {
	if !(isFunc(function)) {
		e := fmt.Sprintf("parameter (%s) is not a function", function)
		return newError(e)
	}

	switch function.(type) {
	case func(uint64):
		fn := function.(func(uint64))
		for _, element := range input {
			fn(element)
		}

	case func(uint64, int):
		fn := function.(func(uint64, int))
		for i, element := range input {
			fn(element, i)
		}

	default:
		e := fmt.Sprintf("function (%s) must have a firm: '%s' or '%s' and got: '%s'",
			getFuncName(function), "func(uint64)", "func(uint64, int)", getFuncFirm(function))

		return newError(e)
	}

	return nil
}

func uint64Includes(input []uint64, e uint64, indexes ...int) bool {
	index := 0

	// We only care for the first index on indexes because the lack of optional
	// parameters in Go
	if len(indexes) > 0 {
		index = indexes[0]
	}

	switch {
	case index >= len(input):
		return false
	case index < 0 && len(input) + index <= -1 * len(input):
		index = 0
	case index < 0:
		index = len(input) + index
	}

	for i := index; i < len(input); i++ {
		if input[i] == e {
			return true
		}
	}

	return false
}

// Uint8
func Uint8ForEach(input []uint8, function interface{}) error {
	if !(isFunc(function)) {
		e := fmt.Sprintf("parameter (%s) is not a function", function)
		return newError(e)
	}

	switch function.(type) {
	case func(uint8):
		fn := function.(func(uint8))
		for _, element := range input {
			fn(element)
		}

	case func(uint8, int):
		fn := function.(func(uint8, int))
		for i, element := range input {
			fn(element, i)
		}

	default:
		e := fmt.Sprintf("function (%s) must have a firm: '%s' or '%s' and got: '%s'",
			getFuncName(function), "func(uint8)", "func(uint8, int)", getFuncFirm(function))

		return newError(e)
	}

	return nil
}

func uint8Includes(input []uint8, e uint8, indexes ...int) bool {
	index := 0

	// We only care for the first index on indexes because the lack of optional
	// parameters in Go
	if len(indexes) > 0 {
		index = indexes[0]
	}

	switch {
	case index >= len(input):
		return false
	case index < 0 && len(input) + index <= -1 * len(input):
		index = 0
	case index < 0:
		index = len(input) + index
	}

	for i := index; i < len(input); i++ {
		if input[i] == e {
			return true
		}
	}

	return false
}

// Uintptr
func UintptrForEach(input []uintptr, function interface{}) error {
	if !(isFunc(function)) {
		e := fmt.Sprintf("parameter (%s) is not a function", function)
		return newError(e)
	}

	switch function.(type) {
	case func(uintptr):
		fn := function.(func(uintptr))
		for _, element := range input {
			fn(element)
		}

	case func(uintptr, int):
		fn := function.(func(uintptr, int))
		for i, element := range input {
			fn(element, i)
		}

	default:
		e := fmt.Sprintf("function (%s) must have a firm: '%s' or '%s' and got: '%s'",
			getFuncName(function), "func(uintptr)", "func(uintptr, int)", getFuncFirm(function))

		return newError(e)
	}

	return nil
}

func uintptrIncludes(input []uintptr, e uintptr, indexes ...int) bool {
	index := 0

	// We only care for the first index on indexes because the lack of optional
	// parameters in Go
	if len(indexes) > 0 {
		index = indexes[0]
	}

	switch {
	case index >= len(input):
		return false
	case index < 0 && len(input) + index <= -1 * len(input):
		index = 0
	case index < 0:
		index = len(input) + index
	}

	for i := index; i < len(input); i++ {
		if input[i] == e {
			return true
		}
	}

	return false
}

