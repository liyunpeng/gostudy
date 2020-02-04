package basic

import "fmt"

type Human interface {
	Say() string
}

type Man struct {
}

func (m *Man) Say() string {
	return "man"
}

func IsNil(h interface{}) bool {
	return h == nil
}

/*
func f2()  {
	var c Man
	var d Human
	//接口变量和实现这个接口类型的变量, 属于两个不同的类型, 不能比较
	fmt.Println( c == d)
}
*/

func interfaceNil() {
	/*
		接口有两种类型:
		内部有方法的, 就是eface类型, 其_type在eface->_type
		内部无方法的, 就是iface类型, 其_type在iface->tab->_type
		接口变量在比较是否相等时, 都是用的_type成员和data成员比较的.
	 */
	var a interface{}
	var d Human

	fmt.Println("not assign a d:", a==d)
	/*
		运行结果:
		not assign a d: true
		理解:
		虽然a是eface类型, d是iface类型,
		但a, d比较时, 使用他们的_type和data成员比较的, 都为nil, 所以相等
	*/

	var b *Man
	var c *Man
	var e interface{}
	/*
		超集包含子集,
		type interface child
		type interface super {
			child
		}
		var s super
		var c child
		则super为超级接口, child为子集接口
		只能是超集接口变量赋值给子集接口变量
		即:
		c = s
		反过来就错误, 即不能 s = c
	*/
	a = b
	e = a

	fmt.Println(a == nil)
	/*
		运行结果:
		false
		理解:
	    a的_type是*Man类型，data是nil，所以为false
	 */

	fmt.Println(a == c)
	/*
	运行结果:
		true
	理解:
		a本身是eface类型, 但eface类型不是用来比较的
		经过a = b,
		a的_type就变为b的类型, 即*Man类型，data是nil, 这个_type是拿来和其他类型比较的
	    c本身就是*Man类型, data是nil, 所以相等
	 */

	fmt.Println("a d:", a == d)
	/*
		运行结果:
		a d:false

		理解:
		a为eface类型，d为iface类型，但eface类型和iface类型都是用来计较的,
		只有里面的_type成员才是用来比较的
		a的_type为*Man类型, data是nil,
		d没有赋值, 所以d的_type是nil, data是nil. 注意d的_type在iface->tab->_type
		所以 a 和 d不相等
	*/

	fmt.Println(c == d)
	// (5) false
	// c和d其实是两种不同的数据类型

	fmt.Println(e == b)
	// (6) true
	// 分析见(4)

	fmt.Println(IsNil(c))
	// (7) false
	// c是*Man类型，以参数的形式传入IsNil方法
	// 虽然c指向的是nil，但是参数i的_type指向的是*Man，所以i不为nil

	fmt.Println(IsNil(d))
	// (8) true
	// d没有指定具体的类型，所以d的itab指向的是nil，data也是nil
}
