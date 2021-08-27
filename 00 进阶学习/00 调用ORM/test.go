

package main

import (
	"fmt"
)

/* 变量的赋值 自增 空指针案例
 */

func main() {
	// 普通值
	var num int // 默认值为0（不赋值）。 普通赋值方法 ps：var num int =1
	fmt.Println("打印值",num)
	fmt.Println("打印地址",&num)
	fmt.Println("引用函数",goValue(num))
	fmt.Println("引用函数，打印地址",goAddress(num))
	num++
	fmt.Println("增加值",num)
	fmt.Println("----------------------------")

	// 指针值
	//var i,num2 *int =  2,4
	var i *int
	fmt.Println("打印值",i)
	// 注意！空指针不能引用
	// 原因：go 初始化指针的时候会为指针 i 的值赋为 nil，，但 i 的值代表的是 *i 的地址，nil 的话系统还并没有给 *i 分配地址，所以这时给 *i 赋值肯定会出错
	// 处理方法：指针赋值前可以先创建一块内存分配给赋值对象即可
	i = new(int) //分配内存
	*i = 1 //分配内存值
	fmt.Println("打印*i 值",i)
	fmt.Println("打印地址",&i)
	fmt.Println("自增i值")

	*i++
	fmt.Println("打印*i 值",i)
	fmt.Println("打印地址",&i)
	fmt.Println("引用函数，打印值",goValue(*i))
	fmt.Println("引用函数，打印地址",goAddress(*i))
	//var x float32 =3.4
	//fmt.Println("type:", reflect.TypeOf(x))
	//
	fmt.Println("----------------------------")
}

func goValue(param int)  int{
	return param
}

func goAddress(param int)  *int{
	return &param
}

