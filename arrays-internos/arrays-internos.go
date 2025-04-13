package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("arrays e slices")

	var array1 [5]string
	array1[0] = "posiçao 1"
	//var array1 [5]int

	fmt.Println(array1)

	array2 := [5]string{"1", "2", "3", "4", "5"}
	fmt.Println(array2)

	slice := []int{10, 11, 12, 13, 14}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array2))

	slice = append(slice, 18)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "posicao alterada"
	fmt.Println(slice2)

	//arrays internos
	slice3 := make([]float32, 10, 15)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) //length
	fmt.Println(cap(slice3)) //capacidade

	slice3 = append(slice3, 5)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) //length
	fmt.Println(cap(slice3)) //capacidade

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	slice4 = append(slice4, 10)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

}
