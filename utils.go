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

func Merge_compute_minrun(n int) int {
	r := 0
	for n >= 64 {
		r |= n & 1
		n >>= 1
	}
	return n + r
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
