package main

import "fmt"

func modify(arr *[3]int) {
    arr[0] = 1000
}

func main() {
    // Declare and initialize
    arr := [3]int{10, 20, 30}

    // Access and update
    arr[1] = 200

    // Loop using range
    for i, val := range arr {
        fmt.Printf("Index %d: %d\n", i, val)
    }

    // Pass to function (copy)
    copyArr := arr
    copyArr[0] = 999
    fmt.Println("Original:", arr)
    fmt.Println("Copy:", copyArr)

    // Modify original using pointer
    modify(&arr)
    fmt.Println("Modified by pointer:", arr)

    // Multidimensional
    matrix := [2][2]int{{1, 2}, {3, 4}}
    fmt.Println("Matrix value:", matrix[1][1]) // Output: 4
}
