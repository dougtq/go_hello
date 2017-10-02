package main

import (
	"fmt"
	"math"
	"time"
)

func main() {

	i, j := testeFunc("World")
	fmt.Println(i + j)
	matematica()
	valores()
	variaveis()
	loops()
	condicionais()
	switchCases()
	arraya()
	mapeado()
	vary(1, 2, 3)
	vary(5, 2, 6, 8, 9, 5, 7, 9)
	fmt.Println(recurs(70))
}

func testeFunc(palavra string) (string, string) {
	return "Hello " + palavra + " \n", "3"
}

func matematica() {
	const n = 500000000
	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
}

func valores() {
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}

func variaveis() {
	var b, c int = 1, 2
	fmt.Println(b, c)
	d := 3
	fmt.Println(d)
}

func loops() {

	i := 0
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	// A classic initial/condition/after `for` loop.
	for j := 0; j <= 4; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	// Use `continue` para ir para próxima iteração
	for n := 0; n <= 2; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}

func condicionais() {
	i := 0
	for i < 7 {
		if i%2 == 0 {
			fmt.Println("is even")
		} else {
			fmt.Println("is odd")
		}
		i++
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}

func switchCases() {
	t := time.Now()
	fmt.Println(time.Now())

	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")

	}
}

func arraya() {
	xy := [5]int{1, 2, 3, 4, 5}
	fmt.Println(xy)

	var _2d [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			_2d[i][j] = i + j
		}
	}
	fmt.Println("_2d: ", _2d)
}

func mapeado() {
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}

func vary(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func recurs(n int) int {
	if n == 0 {
		return 1
	}
	return n + recurs(n-1)
}
