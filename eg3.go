package main

import (
	"fmt"
	"strconv"
)

var Value int = 100 //global variavble-starts with caps - accesible within only this package

var value2 int = 200 //package level variable - starts with small case - accesible within all packages

func main() {

	var a = 12 // function level variablr
	fmt.Println(a)

	var b int = 14
	fmt.Println(b)

	c := 15
	fmt.Println(c)

	var val = 24
	fmt.Printf("%v,%T", val, val)

	var d int32 = 100
	var e int64 = int64(d)
	fmt.Println(e)

	//input
	fmt.Println("Enter your first name")
	var first string
	fmt.Scanln(&first)
	fmt.Println("Enter your second name")
	var second string
	fmt.Scanln(&second)

	fmt.Println("Your entered name is:", first, second)

	//conversion
	var f int = 70
	var s string = strconv.Itoa(f)
	fmt.Println(s)

	var g int = 70
	var h string = string(g) //ascii value
	fmt.Println(h)

	const i = 900
	fmt.Println(i)
	//iota
	const (
		x = iota
		y = iota
	)
	fmt.Println(x, y)

	const (
		j = iota
		k
	)
	fmt.Println(j, k)

	const (
		l = iota
		_
		m
		n
	)
	fmt.Println(l, m, n)
	//array
	var arr [4]int = [4]int{1, 2, 3, 4}
	fmt.Println(arr)

	var array = [...]int{1, 2, 3, 3, 4, 5, 6, 7}
	fmt.Println(array)
	fmt.Println(len(array)) //length of array
	array[1] = 20
	fmt.Println(array)
	fmt.Println(array[3])

	var array2 = array
	fmt.Println(array2)

	array[3] = 56
	fmt.Println(array2)
	fmt.Println(array)

	fmt.Print(array[1:4])

	array3 := array
	fmt.Println(array3)

	var matrix [2][2]int = [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(matrix)
	fmt.Println(matrix[1][0])
}
