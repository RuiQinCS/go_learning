package main

import (
	"fmt"
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

func main() {
	FormatFloat()

	FormatFloatStr()
}
