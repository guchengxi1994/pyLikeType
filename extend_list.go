/*
 * @Descripttion:
 * @version:
 * @Author: xiaoshuyui
 * @email: guchengxi1994@qq.com
 * @Date: 2021-11-12 22:08:30
 * @LastEditors: xiaoshuyui
 * @LastEditTime: 2021-11-12 22:20:52
 */
package pyLikeType

type ExtendList struct {
	List []interface{}
}

func (el *ExtendList) Contains() bool {
	if len(el.List) == 0 {
		return false
	}

	return false
}
