// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at
// {{ .Timestamp }}
package slice

import (
	"reflect"
	"testing"
)
{{ range .Types }}
func Test{{ Title . }}ForEach(t *testing.T) {
	inputInterface := getInput({{ printf "%q" . }})
	input := inputInterface.([]{{ . }})

	expectInterface := getExpect({{ printf "%q" . }})
	expect := expectInterface.([]{{ . }})

	current := make([]{{ . }}, len(expect))

	if err := {{ Title . }}ForEach(input, func(e {{ . }}, i int) {
		current[i] = {{ . }}Function(e)
	}); err != nil {
		t.Errorf("running {{ Title . }}ForEach got the error: %s", err.Error())
	}

	if got := reflect.DeepEqual(current, expect); !got {
		t.Errorf("{{ Title . }}ForEach returns %v; expected %v", current, expect)
	}
}
{{ end }}

func TestIncludes(t *testing.T) {
	t.Run("Testing Includes", func(t *testing.T) {
	{{ range .Types }}
		t.Run("Testing {{ . }}Includes", func (t *testing.T) {
			input := getInput({{ printf "%q" . }})

			elements := input.([]{{ . }})
			expect := true

			// Testing without third parameter
			t.Logf("Testing {{ . }}Includes(%v, %v)", input, elements[0])
			got, err := Includes(input, elements[0])
			if err != nil {
				t.Errorf("running Includes got the error: %s", err.Error())
			}

			if got != expect {
				t.Errorf("Includes returns %v; expected %v", got, expect)
			}

			// Testing with third negative parameter but out the range of the
			// elements of input, expected true
			t.Logf("Testing {{ . }}Includes(%v, %v, %d)", input, elements[0], -10)
			got, err = Includes(input, elements[0], -10)
			if err != nil {
				t.Errorf("running Includes got the error: %s", err.Error())
			}

			if got != expect {
				t.Errorf("Includes returns %v; expected %v", got, expect)
			}

			// Testing with third negative parameter but in the range of the
			// elements of input, expected false
			expect = false

			t.Logf("Testing {{ . }}Includes(%v, %v, %d)", input, elements[0], -1)
			got, err = Includes(input, elements[0], -1)
			if err != nil {
				t.Errorf("running Includes got the error: %s", err.Error())
			}

			if got != expect {
				t.Errorf("Includes returns %v; expected %v", got, expect)
			}

			// Testing with third positive parameter but in the range of the
			// elements of input, expected false
			t.Logf("Testing {{ . }}Includes(%v, %v, %d)", input, elements[0], 1)
			got, err = Includes(input, elements[0], 1)
			if err != nil {
				t.Errorf("running Includes got the error: %s", err.Error())
			}

			if got != expect {
				t.Errorf("Includes returns %v; expected %v", got, expect)
			}

			// Testing with third positive parameter but out the range of the
			// elements of input, expected false
			t.Logf("Testing {{ . }}Includes(%v, %v, %d)", input, elements[0], 10)
			got, err = Includes(input, elements[0], 10)
			if err != nil {
				t.Errorf("running Includes got the error: %s", err.Error())
			}

			if got != expect {
				t.Errorf("Includes returns %v; expected %v", got, expect)
			}
		})
	{{ end }}
	})
}
