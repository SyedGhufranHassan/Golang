package main

import (
	"fmt"
	"math"
	"runtime"
)

func Add(x int, y int) int {
	return x + y
}

func swap(x string, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum - 4*9
	y = sum / 10
	return
}

var z int

func sqrt(x float64) float64 {
	z := math.Sqrt(x)

	return z

}
func compute(fn func(int, int) int) int {
    return fn(10, 5)
}

func multiply(x, y int) int {
    return x * y
}
var c, python bool
var name string = "Ghufran"

func main() {

	fmt.Println("Ghufran ", math.Sqrt(10))
	//fmt.Println("It will gives error because it is not exported name", math.pi)
	fmt.Println("Now it will work good because it is Exported Name", math.Pi)

	defer fmt.Println("Addidtion of 5 + 5:", Add(5, 5))

	a, b := swap("Ghufran", "Syed")
	fmt.Println("Before Swaping: Ghufran Syed, After Swaping:", a, b)

	fmt.Println(split(50))

	var i int
	var age int = 23
	k := 20
	fmt.Println("My name is:", name, ", Age is:", age)
	fmt.Println(i, c, python)

	fmt.Println(k)

	for i := 0; i <= 5; i++ {
		fmt.Println(i)
		if i == 1 {
			fmt.Println("It continue here", i)
			continue
		}
		if i == 3 {
			fmt.Println("It breaks here", i)
			break
		}
	}

	j := 10
	for j < 100 {
		fmt.Println(j)
		j = j + 25
	}
	fmt.Println(sqrt(2))

	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
	day := "Wednesday"
	switch day {
	case "Monday":
		fmt.Println("Today is Monday.")
	case "Tuesday":
		fmt.Println("Today is Tuesday.")
	case "Wednesday":
		fmt.Println("Today is Wednesday.")
	case "Thursday":
		fmt.Println("Today is Thursday.")
	case "Friday":
		fmt.Println("Today is Friday.")
	case "Saturday":
		fmt.Println("Today is Saturday.")
	default:
		fmt.Println("Today is Sunday.")
	}

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")

	x := 10

	ptr := &x

	fmt.Println("Pointer of x: ",*ptr)

	m := map[string]int{"a": 1, "b": 2}

	fmt.Println("Maps")

	for k, v := range m {
		fmt.Printf("Key: %s, Value: %d\n", k, v)
	}
	result := compute(multiply)
	fmt.Println("Result:", result) // Output: 50

}