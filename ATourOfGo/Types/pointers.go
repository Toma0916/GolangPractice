package main

import "fmt"

func main() {
	i, j := 42, 2701

	// 変数Tのポインタは*T型
	// ゼロ値はnil
	// &オペレータによりそのオペランドへのポインタを引き出す
	p := &i         // point to i
	fmt.Println(*p) // read i through pointer
	*p = 21         // set i through pointer
	fmt.Println(i)  // see the new value of i

	p = &j       // point to j
	*p = *p / 37 // divide j through the pointer
	fmt.Println(j)
}
