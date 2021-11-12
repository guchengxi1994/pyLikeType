/*
 * @Descripttion:
 * @version:
 * @Author: xiaoshuyui
 * @email: guchengxi1994@qq.com
 * @Date: 2021-11-12 21:29:57
 * @LastEditors: xiaoshuyui
 * @LastEditTime: 2021-11-12 22:21:45
 */
package pyLikeType

type Iterable interface {
	Next() (int, error)
	Hash() int
	Contains() bool
}

type Any struct {
	Value interface{}
}

func Equals(self interface{}, other interface{}) bool {
	return false
}
