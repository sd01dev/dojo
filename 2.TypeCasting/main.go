package main

import "fmt"

// 1. Что такое Type Casting?
// 2. Написать пример, с использованием приведения типов.

/*
Type Casting - приведение типов, при возможности этого. Например, int к int64 иди float к byte
Синтаксис: x := type(value)
В случае невозможности, IDE подсветит ошибку, например:
Cannot convert an expression of the type 'int64' to the type 'bool'
*/

func main() {
	var valueInt64 int64 = 20
	fmt.Printf("%T\n", valueInt64) // int64
	valueInt := int(valueInt64)
	fmt.Printf("%T\n", valueInt) // int
}
