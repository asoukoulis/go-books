package main

import "fmt"

// z = x &^ y，表示如果 y 中的 bit 位为 1，则 z 对应 bit 位为 0
// 否则 z 对应 bit 位等于 x 中相应的 bit 位的值
// 表达式 z = x | y，如果 y 中的 bit 位为 1，则 z 对应 bit 位为 1
// 否则 z 对应 bit 位等于 x 中相应的 bit 位的值，与 &^ 完全相反

// 取反对于无符号整数来说就是按位取反

// 3
// 00000011 原
// 00000011 反
// 00000011 补

// 11111100 按位取反
// 10000011 发现符号位(即最高位)为1(表示负数)，将除符号位之外的其他数字取反 ＝ 10000001
// 10000100 末位加1取其补码 ＝ 10000010 转换回十进制 ＝ -2

// 4
// 00000100 原
// 00000100 反
// 00000100 补

// 11111011
// 10000100
// 10000101 -5

// 对于
// 将1(这里叫：原码)转二进制 ＝ 00000001
// 按位取反 ＝ 11111110
// 发现符号位(即最高位)为1(表示负数)，将除符号位之外的其他数字取反 ＝ 10000001
// 末位加1取其补码 ＝ 10000010
// 转换回十进制 ＝ -2

// 负数取反是按照补码进行取反的
// -6
// 10000110 原
// 11111001 反
// 11111010 补
// 00000101 取反

// -3
// 10000011 原
// 11111100 反
// 11111101 补
// 00000010 取反

func day006() {
	var a int8 = 3
	fmt.Printf("^%08b=%08b %d\n", a, ^a, ^a) // ^11=-100 -4
	var a4 int8 = 4
	fmt.Printf("^%08b=%08b %d\n", a4, ^a4, ^a4) // ^100=-101 -5
	var b uint8 = 3
	fmt.Printf("^%08b=%08b %d\n", b, ^b, ^b) // ^11=11111100 252
	var c int8 = -3
	fmt.Printf("^%08b=%08b %d\n", c, ^c, ^c) // ^-11=10 2
	var d int8 = -6
	fmt.Printf("^%08b=%08b %d\n", d, ^d, ^d) // ^-110=101 5
}
