func compute(fn func(int, int) int) int {
    return fn(10, 5)
}

func multiply(x, y int) int {
    return x * y
}

func main() {
    result := compute(multiply)
    fmt.Println("Result:", result) // Output: 50
}
