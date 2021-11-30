/*
	python difflib in go
*/

package pyLikeType

import (
	"reflect"
)

func (s *ExtendString) FindLongestMatch(str interface{}) string {
	var result string
	switch reflect.TypeOf(str).String() {
	case "pyLikeType.ExtendString":
		result = find_longest_match(*s, str.(ExtendString))
	case "string":
		result = find_longest_match(*s, ExtendString{Value: str.(string)})
	default:
		result = ""
	}

	return result
}

func find_longest_match(a, b ExtendString) string {
	var result ExtendString
	var longer, shorter ExtendString
	if a.Len() >= b.Len() {
		longer = a
		shorter = b
	} else {
		longer = b
		shorter = a
	}

	var bestMatch = 0

	var _tmp ExtendString

	for i := bestMatch; i < shorter.Len(); i++ {
		if float64(result.Len()) > 0.5*float64(shorter.Len()) {
			break
		}

		for j := 0; j < shorter.Len(); j++ {
			for k := 0; k < longer.Len(); k++ {

				if shorter.SubString(j, j+i+1) == longer.SubString(k, k+i+1) {
					_tmp.Value = shorter.SubString(j, j+i+1)
					bestMatch += 1
				}

				if _tmp.Len() > result.Len() {
					result.Value = _tmp.Value
				}
			}
		}

	}

	return result.Value
}
