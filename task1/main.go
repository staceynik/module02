package main

import "fmt"

func main() {
	// Приведение строки "104" к целому числу
	str := "104"
	num := 0
	fmt.Sscanf(str, "%d", &num)
	fmt.Printf("Значение строки '%s' преобразовано в целое число: %d\n", str, num)

	// Приведение целого числа 35 к строке
	num2 := 35
	str2 := fmt.Sprintf("%d", num2)
	fmt.Printf("Значение целого числа %d преобразовано в строку: '%s'\n", num2, str2)
}
