package calculation

import (
	"fmt"
	"regexp"
	"strconv"
	"unicode"
)

var mapOp = map[string]int{
	"+": 1,
	"-": 1,
	"*": 2,
	"/": 2,
}

func topolish(a string) ([]string, error) {
	if len(a) == 0 {
		return []string{}, ErrEmptyString
	}
	if string(a[0]) == "*" || string(a[0]) == "/" || string(a[0]) == "-" || string(a[0]) == "+" || string(a[0]) == ")" {
		return []string{}, ErrStartWithOperator
	}
	if string(a[len(a)-1]) == "*" || string(a[len(a)-1]) == "/" || string(a[len(a)-1]) == "-" || string(a[len(a)-1]) == "+" || string(a[len(a)-1]) == "(" {
		return []string{}, ErrEndWithOperator
	}
	for i := range len(a) - 1 {
		if (string(a[i]) == "*" || string(a[i]) == "/" || string(a[i]) == "-" || string(a[i]) == "+") && (string(a[i+1]) == "*" || string(a[i+1]) == "/" || string(a[i+1]) == "-" || string(a[i+1]) == "+") {
			return []string{}, ErrTwoConsecOperation
		} else if (string(a[i]) == "(") && (string(a[i+1]) == "*" || string(a[i+1]) == "/" || string(a[i+1]) == "-" || string(a[i+1]) == "+") {
			return []string{}, ErrTwoConsecOperation
		} else if (string(a[i]) == "*" || string(a[i]) == "/" || string(a[i]) == "-" || string(a[i]) == "+") && (string(a[i]) == ")") {
			return []string{}, ErrTwoConsecOperation
		} else if string(a[i]) == "(" && string(a[i+1]) == ")" {
			return []string{}, ErrTwoConsecOperation
		}
	}

	stack := make([]string, 0)
	output := make([]string, 0)
	for i := range len(a) {
		if unicode.IsNumber(rune(a[i])) || string(a[i]) == "." {
			if i == 0 {
				output = append(output, string(a[i]))
			} else if string(a[i-1]) == "*" || string(a[i-1]) == "/" || string(a[i-1]) == "-" || string(a[i-1]) == "+" || string(a[i-1]) == "(" || string(a[i-1]) == ")" || len(output) == 0 {
				output = append(output, string(a[i]))
			} else {
				output[len(output)-1] += string(a[i])
			}
		} else if string(a[i]) == "*" || string(a[i]) == "+" || string(a[i]) == "-" || string(a[i]) == "/" || string(a[i]) == "(" || string(a[i]) == ")" {
			if len(stack) == 0 {
				stack = append(stack, string(a[i]))
			} else if mapOp[string(a[i])] > mapOp[string(stack[len(stack)-1])] || string(a[i]) == "(" {
				stack = append(stack, string(a[i]))
			} else if mapOp[string(a[i])] <= mapOp[string(stack[len(stack)-1])] && string(a[i]) != ")" && string(a[i]) != "(" {
				for j := len(stack) - 1; j >= 0; j-- {
					if string(stack[j]) == "(" {
						break
					} else if mapOp[string(stack[j])] < mapOp[string(a[i])] {
						break
					}
					output = append(output, string(stack[j]))
					stack[len(stack)-1] = ""
					stack = stack[:len(stack)-1]
				}
				stack = append(stack, string(a[i]))
			} else if string(a[i]) == ")" {
				for j := len(stack) - 1; j >= 0; j-- {
					if string(stack[j]) == "(" {
						break
					}
					output = append(output, string(stack[j]))
					stack[len(stack)-1] = ""
					stack = stack[:len(stack)-1]
				}
				stack[len(stack)-1] = ""
				stack = stack[:len(stack)-1]
			}

		}
	}
	for i := range len(stack) {
		if string(stack[i][0]) == "(" {
			return []string{}, ErrNoClosingParenthesis
		}
	}
	for j := len(stack) - 1; j >= 0; j-- {
		output = append(output, string(stack[j]))
		stack[len(stack)-1] = ""
		stack = stack[:len(stack)-1]
	}
	return output, nil
}
func Calc(expression string) (float64, error) {
	ans, err := regexp.Match(`[^0-9+\-*/]`, []byte(expression))
	if err != nil {
		return 0, err
	}
	if ans {
		return 0, ErrInvInputs
	}
	polish, err := topolish(expression)
	if err != nil {
		return 0, ErrInvInputs
	}
	for len(polish) > 1 {
		saveop := 0
		for i := 0; i < len(polish); i++ {
			if (string(polish[i][0]) == "*" || string(polish[i][0]) == "+" || string(polish[i][0]) == "-" || string(polish[i][0]) == "/") && len(polish[i]) == 1 {
				saveop = i
				break
			}
		}
		a, _ := strconv.ParseFloat(polish[saveop-2], 64)
		b, _ := strconv.ParseFloat(polish[saveop-1], 64)
		switch string(polish[saveop]) {
		case "*":
			polish[saveop-2] = fmt.Sprint(a * b)
		case "/":
			if b == 0 {
				return 0, ErrInvInputs
			}
			polish[saveop-2] = fmt.Sprint(a / b)
		case "+":
			polish[saveop-2] = fmt.Sprint(a + b)
		case "-":
			polish[saveop-2] = fmt.Sprint(a - b)
		}
		copy(polish[saveop:], polish[saveop+1:])
		polish[len(polish)-1] = ""
		polish = polish[:len(polish)-1]
		copy(polish[saveop-1:], polish[saveop:])
		polish[len(polish)-1] = ""
		polish = polish[:len(polish)-1]
	}
	bebra, err := strconv.ParseFloat(polish[0], 64)
	if err != nil {
		return 0, err
	}
	return bebra, nil
}
