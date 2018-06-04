package main

import (
	"fmt"
	"math"
	"math/rand"
	"math/cmplx"
)
//函数
func add(x, y int) int {
	return x+y
}

//多值返回
func swap(x,y string) (string, string)  {
	return y, x
}
//命名返回值 return直接返回
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
//变量
	var c, python, java bool
//变量的初始化
var o, j int = 1, 2

//基本类型
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

//常量
const Pi = 3.14

//数值常量
const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {

	fmt.Println("My favorite number is", rand.Intn(100))
//方法名首字母大写，为导出包
	fmt.Println(math.Pi)
//方法调用
	fmt.Println(add(2,5))
//多值返回
	a, b :=swap("hello", "world")
	fmt.Println(a, b)
//命名返回值return
	fmt.Println(split(10))
//变量
	var i int
	fmt.Println(i, c, python, java)
//变量的初始化
	var c, python, java = true, false, "no!"
	k := 4
	fmt.Println(i, j, c, python, java,k)
// 简洁赋值语句
	var ii, j int = 1, 2
	kk := 3
	cc, cpython, cjava := true, false, "no!"

	fmt.Println(ii, j, kk, cc, cpython, cjava)
// 基本类型操作
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)


	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)

	v := 42 // 修改这里！
	fmt.Printf("v is of type %T\n", v)

	//常量
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)

//数值常量
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}