package main

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		val    string = "Amith B"
		valNum int    = 27
	)

	valNum += 1

	shortVal := "Short Value"
	var cvtFlt float32 = float32(valNum)

	str := strconv.Itoa(valNum)

	fmt.Printf("Hi, %v %.4f %s \n", val, cvtFlt, str)
	fmt.Println(shortVal)
	fmt.Println(shortVal, cvtFlt)

	fmt.Println(!true)

	if true {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

	switch '1' {
	case 1, '1':
		fmt.Println(1)
		fmt.Println("Line 2")
		fallthrough
	case 2:
		fmt.Println(2)
	}
	fmt.Println("___________")
	switch {
	case 2+2 == 4:
		fmt.Println(4)
	case 3+3 == 6:
		fmt.Println(6)
	}

	fmt.Println("___________")

	for i := 1; i < 5; i++ {
		fmt.Println(i)
	}

	i := 1

	fmt.Println("___________")

	for {
		fmt.Println(i)
		i++

		if i == 10 {
			break
		}
	}
	fmt.Println("___________")

	var strList [4]string = [4]string{"a", "b", "c", "d"}
	fmt.Println(strList, len(strList))

	for index, item := range strList {
		fmt.Println(index, item)
	}

	strListNew := strList[1:2]

	strListUpdated := append(strListNew, "e", "f", "g")

	fmt.Println((strList), strListNew, strListUpdated)
	fmt.Println(cap(strList), cap(strListNew), cap(strListUpdated))
	fmt.Println(len(strList), len(strListNew), len(strListUpdated))
	fmt.Println("___________")

	var mapItem map[string]string = map[string]string{"a": "apple", "b": "ball", "c": "cat"}

	mapValue, found := mapItem["a"]

	fmt.Println(mapValue, found)

	fmt.Println("___________")

	fmt.Println("add", add(2, 3))

	fmt.Println("___________")

	add, mul := operation(2, 3)
	fmt.Println("add and mul:", add, mul)
	fmt.Println("___________")

	anonymFunc := func(a int, b int) int {
		return a + b
	}

	fmt.Println("AnonymFunc", anonymFunc(5, 5))

	fmt.Println("___________")

	var pVal int = 20
	var ptr *int = &pVal

	*ptr = 30

	// *(&val) -> * is a dereferencing operator and don't confuse it with * in *int
	fmt.Println("Pointer", *(&pVal), (ptr))

	fmt.Println("___________")

	// strct:= new(Student)

	strct := Student{
		name:   "Amith",
		rollNo: 1,
	}

	strctNew := Student{
		"Amith",
		1,
	}

	fmt.Println("Struct", strct.name, strct.rollNo)
	fmt.Println("Struct", strctNew.name, strctNew.rollNo)
	fmt.Println("___________")

	crlAr := Circle{
		radius: 20,
	}

	printData(&crlAr)
}

// func  <receiver> <functionName> returnType
func (c *Circle) calcArea() int {
	c.area = 3.142 * c.radius * c.radius

	return int(c.area)
}

type CircleMethods interface {
	calcArea() int
}

func printData(c CircleMethods) {
	fmt.Println("Method with interface", c.calcArea())
}

type Circle struct {
	radius float64
	area   float64
}

type Student struct {
	name   string
	rollNo int
}

func add(v ...int) int {
	sum := 0

	for _, value := range v {
		sum += value
	}
	return sum
}

func operation(v ...int) (sum int, mul int) {
	sum = 0
	mul = 1

	for _, value := range v {
		sum += value
		mul *= value
	}
	return
}
