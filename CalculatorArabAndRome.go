package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func calc(input string) {
	var num1, num2 int
	parts := strings.Split(input, " ")
	if len(parts) != 3 {
		fmt.Print("некорректный ввод. Пример: '3 + 4'")
		os.Exit(0)
	}
	num1, err := strconv.Atoi(parts[0])
	if err == nil {
		if num1 < 0 {
			fmt.Print("некорректный ввод: цифры не могут быть отрицательными")
			os.Exit(0)
		} else if num1 == 0 {
			fmt.Print("некорректный ввод: цифры не могут быть равны нулю")
			os.Exit(0)
		} else if num1 > 10 {
			fmt.Print("некорректный ввод: введите целые числа от 1 до 10 или от I до X")
			os.Exit(0)
		}
	} else if err != nil {
		num1 = convertRomeToArab(parts[0])
		if num1 < 0 {
			fmt.Print("некорректный ввод: римские цифры не могут быть отрицательными")
			os.Exit(0)
		} else if num1 == 0 {
			fmt.Print("некорректный ввод: римские цифры не могут быть равны нулю")
			os.Exit(0)
		} else if num1 == 11 {
			fmt.Print("некорректный ввод: введите целые числа от 1 до 10 или от I до X")
			os.Exit(0)
		}
	}
	num2, err1 := strconv.Atoi(parts[2])
	if err1 == nil {
		if num2 < 0 {
			fmt.Print("некорректный ввод: цифры не могут быть отрицательными")
			os.Exit(0)
		} else if num2 == 0 {
			fmt.Print("некорректный ввод: цифры не могут быть равны нулю")
			os.Exit(0)
		} else if num2 > 10 {
			fmt.Print("некорректный ввод: введите целие числа от 1 до 10 или от I до X")
			os.Exit(0)
		}
	} else if err1 != nil {
		num2 = convertRomeToArab(parts[2])
		if num2 < 0 {
			fmt.Print("некорректный ввод: римские цифры не могут быть отрицательными")
			os.Exit(0)
		} else if num2 == 0 {
			fmt.Print("некорректный ввод: римские цифры не могут быть равны нулю")
			os.Exit(0)
		} else if num2 == 11 {
			fmt.Print("некорректный ввод: введите целые числа от 1 до 10 или от I до X")
			os.Exit(0)
		}
	}
	result := 0
	switch parts[1] {
	case "+":
		result = num1 + num2
		break
	case "-":
		result = num1 - num2
		break
	case "*":
		result = num1 * num2
		break
	case "/":
		result = num1 / num2
	default:
		fmt.Print("неизвестная операция")
		os.Exit(0)
	}

	if isArabic(parts[0], parts[2]) {
		fmt.Println(result)
	} else if isRome(parts[0], parts[2]) {
		if result <= 0 {
			fmt.Println("ошибка: результат вычисления римских чисел не может быть отрицательным или равен нулю")
			os.Exit(0)
		}
		fmt.Println(convertArabToRome(result))
	} else {
		fmt.Println("некорректный ввод: введите только арабские или только римские цифры")
	}
}
func convertRomeToArab(rome string) int {
	arabicValues := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	romanSymbols := [10]string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	romeArabPairs := make(map[string]int)
	for i := 0; i < 10; i++ {
		romeArabPairs[romanSymbols[i]] = arabicValues[i]
	}
	if romeArabPairs[rome] == 0 {
		return 11
	} else {
		return romeArabPairs[rome]
	}
}
func convertArabToRome(arab int) string {
	arabicValues := [9]int{100, 90, 50, 40, 10, 9, 5, 4, 1}
	romanSymbols := [9]string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	var res string
	for i := 0; i < len(arabicValues); i++ {
		for arab >= arabicValues[i] {
			res += romanSymbols[i]
			arab -= arabicValues[i]
		}
	}
	return res
}
func isArabic(numb1, numb2 string) bool {
	re := regexp.MustCompile(`\d+`)
	isArabicNum1 := re.MatchString(numb1)
	isArabicNum2 := re.MatchString(numb2)
	return isArabicNum1 && isArabicNum2
}
func isRome(numb1, numb2 string) bool {
	re := regexp.MustCompile(`\d+`)
	isArabicNum1 := re.MatchString(numb1)
	isArabicNum2 := re.MatchString(numb2)
	return !isArabicNum1 && !isArabicNum2
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	calc(input)
}
