package main

import (
	"fmt"
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

func main() {
	FormatFloat()
}
