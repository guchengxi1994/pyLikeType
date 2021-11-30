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
	CompareFunction func(i, j int) int // <0 means List[i]<List[j],==0 means List[i]==List[j],>0 means List[i]>List[j]
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

func (el *ExtendList) Swap(i, j int) {
	el.List.Swap(i, j)
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

func (el ExtendList) Less(i, j int) bool {
	if el.CompareFunction == nil {
		return el.List.Less(i, j)
	} else {
		return el.CompareFunction(i, j) < 0
	}
}

func (el ExtendList) More(i, j int) bool {
	if el.CompareFunction == nil {
		return true
	} else {
		return el.CompareFunction(i, j) > 0
	}
}

func (el ExtendList) Equals(i, j int) bool {
	if el.CompareFunction == nil {
		return false
	} else {
		return el.CompareFunction(i, j) == 0
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
	if el.CompareFunction == nil {
		if i == 0 {
			sort.Sort(el.List)
		} else {
			sort.Sort(sort.Reverse(el.List))
		}
	} else {
		if i == 0 {
			el.insertionSort(0, el.Len())
		} else {
			el.insertionSort(0, el.Len())
			el.Reverse()
		}
	}
}

// copied from sort.Sort
func (el *ExtendList) insertionSort(a, b int) {
	for i := a + 1; i < b; i++ {
		for j := i; j > a && el.Less(j, j-1); j-- {
			el.Swap(j, j-1)
		}
	}
}
