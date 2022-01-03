package main

import "fmt"

func main() {
	/* INTEGER */
	myIntList := List[int]{}
	myIntList.insert(1, 2, 3)
	myIntList.insert(4)
	myIntList.insert(5, 6)
	myIntList.insert(7, 8, 9, 10)

	fmt.Println("forward", myIntList.returnAsSlice())
	fmt.Println("backward", myIntList.returnReverseAsSlice())

	fmt.Println("Deleteing integers")

	myIntList.deleteByValue(10)
	myIntList.deleteByValue(3)
	myIntList.deleteByValue(2)
	myIntList.deleteByValue(1)

	fmt.Println("forward", myIntList.returnAsSlice())
	fmt.Println("backward", myIntList.returnReverseAsSlice())

	/* FLOAT64 */

	myFloat64List := List[float64]{}
	myFloat64List.insert(1.3, 2.4, 3.0)
	myFloat64List.insert(4.5)
	myFloat64List.insert(5.987, 6.7)
	myFloat64List.insert(7.2, 8.9, 9.34, 10.23)

	fmt.Println("forward", myFloat64List.returnAsSlice())
	fmt.Println("backward", myFloat64List.returnReverseAsSlice())

	fmt.Println("Deleteing floats")

	myFloat64List.deleteByValue(10.23)
	myFloat64List.deleteByValue(3.0)
	myFloat64List.deleteByValue(2.4)
	myFloat64List.deleteByValue(1.3)

	fmt.Println("forward", myFloat64List.returnAsSlice())
	fmt.Println("backward", myFloat64List.returnReverseAsSlice())

}
