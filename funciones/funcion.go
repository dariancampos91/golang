package main

import "fmt"

func add(x int, y int) int {
    return x + y
}
func swap(w, z string) (string, string) {
    return w, z
}

func main() {
	a, b := swap("La ", "Suma : ")
	fmt.Print(a,b)
    fmt.Println(add(42, 13))
}
