package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func romanToArabic(roman string) (int, error) {
	romanMap := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}

	return romanMap[roman], nil
}

func arabicToRoman(arabic int) string {
	romanMap := map[int]string{
		1:  "I",
		2:  "II",
		3:  "III",
		4:  "IV",
		5:  "V",
		6:  "VI",
		7:  "VII",
		8:  "VIII",
		9:  "IX",
		10: "X",
	}

	return romanMap[arabic]
}

func calculate(a, b int, operator string) int {
	switch operator {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		if b == 0 {
			panic("Деление на ноль")
		}
		return a / b
	default:
		panic("Неподдерживаемая операция: " + operator)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Введите выражение (например, 3 + 5):")
	input, _ := reader.ReadString('\n')

	parts := strings.Fields(input)
	if len(parts) != 3 {
		panic("Некорректный формат ввода")
	}

	isArabic := strings.ContainsAny(parts[0], "0123456789") && strings.ContainsAny(parts[2], "0123456789")
	isRoman := strings.ContainsAny(parts[0], "IVXLCDM") && strings.ContainsAny(parts[2], "IVXLCDM")
	if !(isArabic || isRoman) {
		panic("Калькулятор умеет работать только с арабскими или римскими цифрами одновременно")
	}

	a, err := strconv.Atoi(parts[0])
	if err != nil && isArabic {
		panic("Первый операнд должен быть арабским числом от 1 до 10")
	}
	if isRoman {
		a, err = romanToArabic(parts[0])
		if err != nil {
			panic("Первый операнд должен быть римской цифрой (I-X)")
		}
	}

	b, err := strconv.Atoi(parts[2])
	if err != nil && isArabic {
		panic("Второй операнд должен быть арабским числом от 1 до 10")
	}
	if isRoman {
		b, err = romanToArabic(parts[2])
		if err != nil {
			panic("Второй операнд должен быть римской цифрой (I-X)")
		}
	}

	if a >= 1 && a <= 10 && b >= 1 && b <= 10 {
		result := calculate(a, b, parts[1])

		if strings.ContainsAny(parts[0], "IVXLCDM") {
			fmt.Println(arabicToRoman(result))
		} else {
			fmt.Println(result)
		}
	} else {
		panic("Числа должны быть от 1 до 10 включительно")
	}
}
