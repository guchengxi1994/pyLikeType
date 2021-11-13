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

type ExtendList struct {
	List []interface{}
}

func (el *ExtendList) Contains(value interface{}) bool {
	if len(el.List) == 0 {
		return false
	}

	return false
}
