// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at
// {{ .Timestamp }}
package slice

import "fmt"

func Includes(inputInterface interface{}, eInterface interface{}, indexes ...int) (bool, error) {
	switch e := eInterface.(type) {
	{{ range .Types }}
	case {{ . }}:
		input, ok := inputInterface.([]{{ . }})
		if !ok {
			e := fmt.Sprintf("%s must be from {{ . }}", inputInterface)
			return false, newError(e)
		}
		return {{ . }}Includes(input, e, indexes...), nil
	{{ end }}
	}

	return false, nil
}

{{ range .Types }}
// {{ Title . }}
func {{ Title . }}ForEach(input []{{ . }}, function interface{}) error {
	if !(isFunc(function)) {
		e := fmt.Sprintf("parameter (%s) is not a function", function)
		return newError(e)
	}

	switch function.(type) {
	case func({{ . }}):
		fn := function.(func({{ . }}))
		for _, element := range input {
			fn(element)
		}

	case func({{ . }}, int):
		fn := function.(func({{ . }}, int))
		for i, element := range input {
			fn(element, i)
		}

	default:
		e := fmt.Sprintf("function (%s) must have a firm: '%s' or '%s' and got: '%s'",
			getFuncName(function), "func({{ . }})", "func({{ . }}, int)", getFuncFirm(function))

		return newError(e)
	}

	return nil
}

func {{ . }}Includes(input []{{ . }}, e {{ . }}, indexes ...int) bool {
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
{{ end }}