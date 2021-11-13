/*
 * @Descripttion:
 * @version:
 * @Author: xiaoshuyui
 * @email: guchengxi1994@qq.com
 * @Date: 2021-11-13 10:57:48
 * @LastEditors: xiaoshuyui
 * @LastEditTime: 2021-11-13 10:57:54
 */
package pyLikeType

import "strconv"

func IsNum(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}
