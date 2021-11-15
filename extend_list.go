/*
 * @Descripttion:
 * @version:
 * @Author: xiaoshuyui
 * @email: guchengxi1994@qq.com
 * @Date: 2021-11-12 22:08:30
 * @LastEditors: xiaoshuyui
 * @LastEditTime: 2021-11-13 09:58:15
 */
package pyLikeType

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

type ExtendList struct {
	List
}

type List []interface{}

func (l List) Len() int {
	return len(l)
}

func (el *ExtendList) Contains(value interface{}) bool {
	for _, v := range el.List {
		if v == value {
			return true
		}
	}
	return false
}

func (el *ExtendList) Find(value interface{}) int {
	for i, v := range el.List {
		if v == value {
			return i
		}
	}
	return -1

}

func (el *ExtendList) Len() int {
	return el.List.Len()
}

func (l List) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l List) Less(i, j int) bool {
	if reflect.TypeOf(l[i]) != reflect.TypeOf(l[j]) {
		return true
	}

	if !reflect.TypeOf(l[i]).Comparable() || !reflect.TypeOf(l[j]).Comparable() {
		return true
	} else {
		switch v := l[i].(type) {
		case string:
			return strings.Compare(v, l[j].(string)) < 0
		case int:
			return v < l[j].(int)
		case float32:
			return v < l[j].(float32)
		case float64:
			return v < l[j].(float64)
			// ...
		}
		return true
	}

}

func (el *ExtendList) Append(value interface{}) {
	el.List = append(el.List, value)
}

func (el *ExtendList) Remove(value interface{}) {
	for i, v := range el.List {
		if v == value {
			el.RemoveByIndex(i)
		}
	}

}

func (el *ExtendList) RemoveByIndex(index int) error {
	if index > el.Len()-1 {
		return fmt.Errorf("out of boundary, List length is %v", el.Len())
	}

	el.List = append(el.List[:index], el.List[index+1:]...)
	return nil
}

func (el *ExtendList) Reverse() {
	for i := 0; i < el.Len()/2; i++ {
		el.List[el.Len()-1-i], el.List[i] = el.List[i], el.List[el.Len()-1-i]
	}
}

// 1 for desc 0 for asc
func (el *ExtendList) Sort(i int) {
	if i == 0 {
		sort.Sort(el.List)
	} else {
		sort.Sort(sort.Reverse(el.List))
	}
}
