/*
 * @Descripttion:
 * @version:
 * @Author: xiaoshuyui
 * @email: guchengxi1994@qq.com
 * @Date: 2021-11-12 22:08:20
 * @LastEditors: xiaoshuyui
 * @LastEditTime: 2021-11-13 11:06:31
 */
package pyLikeType

import (
	"errors"
	"unicode"
)

type ExtendString struct {
	Value string
}

func (s *ExtendString) Len() int {
	var length int = 0
	for range s.Value {
		length += 1
	}
	return length
}

func (s *ExtendString) Reverse() string {
	strArr := []rune(s.Value)
	for i := 0; i < len(strArr)/2; i++ {
		strArr[len(strArr)-1-i], strArr[i] = strArr[i], strArr[len(strArr)-1-i]
	}
	return string(strArr)
}

func (s *ExtendString) Index(id int) (string, error) {
	if id > s.Len()-1 {
		return "", errors.New("out of boundray")
	}

	_value := []rune(s.Value)

	if float32(id) > float32(s.Len())*0.5 {
		_value = []rune(s.Reverse())
		for i, v := range _value {
			if i == s.Len()-id-1 {
				return string(v), nil
			}
		}
	}

	for i, v := range _value {
		if i == id {
			return string(v), nil
		}
	}
	return "", errors.New("can not find")
}

// number type in go is `int` or `float` + number or complex
func (s *ExtendString) IsNumberType() bool {
	return s.IsComplex() || s.IsFloat() || s.IsInt()
}

func (s *ExtendString) IsComplex() bool {
	return s.Value == "complex64" || s.Value == "complex128"
}

func (s *ExtendString) IsFloat() bool {
	return s.Value == "float32" || s.Value == "float64"
}

func (s *ExtendString) IsInt() bool {
	return s.Value == "int" || s.Value == "int8" || s.Value == "int16" || s.Value == "int32" || s.Value == "int64"
}

func (s *ExtendString) IsNumber() bool {
	return false
}

func (s *ExtendString) SplitNumberType() (string, string) {

	for i, v := range s.Value {
		if i == s.Len()-1 {
			return s.Value, ""
		} else {
			_v, err := s.Index(i + 1)
			if err == nil && (unicode.IsLetter(v) && IsNum(_v)) {
				return string([]rune(s.Value)[:i+1]), string([]rune(s.Value)[i+1:])
			}
		}
	}

	return "", ""

}

func (s *ExtendString) SubString(start, end int) string {
	if start > end {
		return ""
	}
	if end > s.Len() {
		return ""
	}
	var substring string = ""

	for i := start; i < end; i++ {
		v, _ := s.Index(i)
		substring += v
	}
	return substring

}

func (s *ExtendString) StartsWith(pattern string) bool {
	_e := ExtendString{
		Value: pattern,
	}

	_length := _e.Len()
	return s.SubString(0, _length) == pattern

}

func (s *ExtendString) EndsWith(pattern string) bool {
	_e := ExtendString{
		Value: pattern,
	}

	_length := _e.Len()
	return s.SubString(s.Len()-_length, s.Len()) == pattern
}
