'''
Descripttion: 
version: 
Author: xiaoshuyui
email: guchengxi1994@qq.com
Date: 2021-11-12 22:38:41
LastEditors: xiaoshuyui
LastEditTime: 2021-11-12 22:40:29
'''


"""
s[0]*31^(n-1) + s[1]*31^(n-2) + ... + s[n-1]
使用 int 算法，这里 s[i] 是字符串的第 i 个字符的 ASCII 码，n 是字符串的长度，^ 表示求幂。空字符串的哈希值为 0。

java


public class Test {
        public static void main(String args[]) {
                String Str = new String("www.runoob.com");
                System.out.println("字符串的哈希码为 :" + Str.hashCode() );
        }
}

"""
a = hash(5)