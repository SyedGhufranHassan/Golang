package main

import (
    "encoding/json"
    "fmt"
)

type Address struct {
    City  string `json:"city"`
    State string `json:"state"`
}

type Person struct {
    Name    string  `json:"name"`
    Age     int     `json:"age"`
    Address Address `json:"address"`
}

func (p *Person) Birthday() {
    p.Age++
}

func main() {
    p := Person{
        Name: "Ghufran",
        Age:  23,
        Address: Address{
            City:  "Lahore",
            State: "Punjab",
        },
    }

    p.Birthday()

    data, _ := json.MarshalIndent(p, "", "  ")
    fmt.Println(string(data))
}
