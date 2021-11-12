/*
 * @Descripttion:
 * @version:
 * @Author: xiaoshuyui
 * @email: guchengxi1994@qq.com
 * @Date: 2021-11-12 22:18:39
 * @LastEditors: xiaoshuyui
 * @LastEditTime: 2021-11-12 22:24:24
 */
package pyLikeType_test

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_interface_0(t *testing.T) {
	var value interface{}
	value = 3.1415
	fmt.Println(reflect.TypeOf(value).Name())

	var i float32
	var j float64

	i = 3.1514
	j = 3.1514

	fmt.Println(float64(i) == j)
	fmt.Println(float64(i))
	fmt.Println(float64(j))
}
