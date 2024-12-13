package main

import "fmt"

func sumPoints(a, x1, y1, x2, y2, ord int) (x3, y3 int) {
	switch {
	case x1 != x2:
		lamb := mod((y2-y1)*inverseElement(x2-x1, ord), ord)
		x3 = mod((lamb*lamb)-x1-x2, ord)
		y3 = mod(lamb*(x1-x3)-y1, ord)
	case x1 == x2 && y1 == y2 && y1 != 0 && y2 != 0:
		lamb := mod((3*(x1*x1)+a)*inverseElement(2*y1, ord), ord)
		x3 = mod((lamb*lamb)-x1-x2, ord)
		y3 = mod(lamb*(x1-x3)-y1, ord)
	}
	return x3, y3
}

// ExtendedGCD реализует расширенный алгоритм Евклида.
func ExtendedGCD(a, ord int) (gcd, x, y int) {
	if ord == 0 {
		return a, 1, 0
	}

	gcd, x1, y1 := ExtendedGCD(ord, a%ord)
	x = y1
	y = x1 - (a/ord)*y1

	return gcd, x, y
}

// ModInverse вычисляет обратный элемент числа a по модулю m.
func inverseElement(a, ord int) int {
	gcd, x, _ := ExtendedGCD(a, ord)
	if gcd != 1 {
		return -1 // Обратного элемента не существует
	}
	inverse := (x%ord + ord) % ord
	return inverse
}

func mod(num, ord int) int {
	ans := num % ord
	if ans < 0 {
		ans += ord
	}
	return ans
}

func main() {
	a := 2
	//b := 1
	x1, y1 := 2, 2
	x2, y2 := 2, 2
	//k := 4
	ord := 13

	x, y := sumPoints(a, x1, y1, x2, y2, ord)

	fmt.Println(x, y)
}
