package main

import (
	"strings"
	"unicode"
)

var (
	operations = map[rune]int{
		'(': 1,
		')': 1,
		'+': 2,
		'-': 2,
		'*': 3,
		'/': 3,
		'^': 4,
	}
)

func isOperation(sym rune) bool {
	_, ok := operations[sym]
	if ok {
		return true
	}
	return false
}

func alg(str string) {
	var stack []rune
	var output []string
	var number strings.Builder

	// проблема в том, что я прохожусь по символам: `12` = '1', '2'
	for _, s := range str {

		if isOperation(s) {
			if number.Len() > 0 {
				output = append(output, number.String())
				number.Reset()
			}

			if len(stack) == 0 {
				stack = append(stack, s)
			} else {
				//substr := str[i : len(str)-1]
				//runesSubstr := []rune(substr)
				for _, r := range popBiggerPriority(&stack, s) {
					output = append(output, string(r))
				}

				if s == '(' {
					stack = append(stack, s)
				} else if s == ')' {
					if len(stack) > 0 {
						for len(stack) > 0 {
							if stack[len(stack)-1] != '(' {
								output = append(output, string(pop(&stack)))
							} else {
								output = output[0 : len(output)-1]
							}
						}
					}
				} else {
					stack = append(stack, s)
				}
			}
		} else if unicode.IsDigit(s) {
			number.WriteRune(s)
		} else {
			continue
		}
	}

}

func pop(stack *[]rune) rune {
	sym := (*stack)[len(*stack)-2 : len(*stack)-1]
	*stack = (*stack)[0 : len(*stack)-2]

	return sym[0]
}

func popBiggerPriority(stack *[]rune, sym rune) []rune {
	var popArray []rune
	if sym == '(' || sym == ')' {
		return []rune{}
	}
	for i := len(*stack) - 1; i > 0; i-- {
		if operations[(*stack)[i]] > operations[sym] {
			popArray = append(popArray, pop(stack))
		}
	}

	return popArray
}

func main() {
	s := "(731+8)*(12+5)-19"
	alg(s)
}
