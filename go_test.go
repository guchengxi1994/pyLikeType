/*
 * @Descripttion:
 * @version:
 * @Author: xiaoshuyui
 * @email: guchengxi1994@qq.com
 * @Date: 2021-11-12 22:18:39
 * @LastEditors: xiaoshuyui
 * @LastEditTime: 2021-11-13 09:29:19
 */
package pyLikeType_test

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_interface_0(t *testing.T) {
	var value interface{} = 3.1415
	fmt.Println(reflect.TypeOf(value).Name())

	var i float32
	var j float64

	i = 3.1514
	j = 3.1514

	fmt.Println(float64(i) == j)
	fmt.Println(float64(i))
	fmt.Println(float64(j))

	var ii int16
	var jj int32
	ii = 1234
	jj = 1234

	fmt.Println(int32(ii) == jj)

}

func Test_len_0(t *testing.T) {
	a := "中国"
	fmt.Println(len(a))
	for i := 0; i < len(a); i++ {
		fmt.Printf("%x ", a[i])
	}
	for i := 0; i < len(a); i++ {
		fmt.Printf("%c ", a[i])
	}

}
