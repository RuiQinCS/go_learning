package main

import (
	"fmt"
	"math"
	"strconv"
)

func FormatFloat() {
	var v interface{}
	v = float64(2147480000)

	fmt.Println(fmt.Sprintf("%v", v))   //2.14748e+09
	fmt.Println(fmt.Sprintf("%+v", v))  //2.14748e+09
	fmt.Println(fmt.Sprintf("%.0f", v)) //2147480000
	fmt.Println(fmt.Sprintf("%f", v))   //2147480000.000000
	fmt.Println(fmt.Sprint(v))          //2.14748e+09
	fmt.Println(fmt.Sprintf("%v", "-")) //-
	fmt.Println(fmt.Sprintf("%v", "0")) //0
}

func FormatFloatStr() {
	var v interface{}
	v = float64(2147480000.9999)

	fmt.Println(strconv.FormatFloat(v.(float64), 'f', -1, 64))
}

// prec 代表小数位数
func TrunFloat(f float64, prec int) float64 {
	x := math.Pow10(prec)
	return math.Trunc(f*x) / x
}

func FormatFloatTrun() {
	f := 2.0 / 3

	f2 := TrunFloat(f, 4)

	fmt.Printf("%v\n", f2)

	// -1 参数表示保持原小数位数，千万要注意，如果你指定了位数就会四舍五入了
	d2 := strconv.FormatFloat(f2, 'f', -1, 64)

	fmt.Printf("%s\n", d2)
}

func FloatTrun() {
	fmt.Println(math.Trunc(0.999))
	fmt.Println(math.Trunc(2.199912121))
	fmt.Println(math.Trunc(1.999212121))
	fmt.Println(math.Trunc(0.9991212))
	fmt.Println(math.Trunc(3.999))
}

func main() {
	FormatFloat() // 测试fmt.Sprintf

	FormatFloatStr() // 测试strconv.FormatFloat

	FormatFloatTrun() // 截断浮点数格式化

	FloatTrun() // 测试math.Trunc
}
