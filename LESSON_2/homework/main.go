package main

import "fmt"

func FibonacciIterative(n int) int {
	// Функція вираховує і повертає n-не число фібоначчі
	// Імплементація без використання рекурсії
	if n <= 1 {
		return n
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

func FibonacciRecursive(n int) int {
	// Функція вираховує і повертає n-не число фібоначчі
	// Імплементація з використанням рекурсії
	if n <= 1 {
		return n
	}
	return FibonacciRecursive(n-1) + FibonacciRecursive(n-2)
}

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

func IsBinaryPalindrome(n int) bool {
	if n == 0 {
		return true
	}

	// Convert number to binary string
	binary := ""
	for n > 0 {
		binary = string('0'+n%2) + binary
		n /= 2
	}

	// Check if binary string is palindrome
	left, right := 0, len(binary)-1
	for left < right {
		if binary[left] != binary[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func ValidParentheses(s string) bool {
	stack := make([]rune, 0)
	mapping := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, char)
		} else if char == ')' || char == '}' || char == ']' {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if mapping[char] != top {
				return false
			}
		}
	}

	return len(stack) == 0
}

func Increment(num string) int {
	// Convert binary string to integer
	n := 0
	for _, ch := range num {
		n = n*2 + int(ch-'0')
	}
	// Increment by 1
	return n + 1
}

func main() {

	// TEST № 1
	if IsPrime(2) == true {
		println("PASS")
	}
	if IsPrime(2) == false {
		panic("FAILURE")
	}

	// TEST № 2
	if IsPrime(8) == false {
		println("PASS")
	}
	if IsPrime(8) == true {
		panic("FAILURE")
	}

	// TEST № 3
	if IsPrime(38) == false {
		println("PASS")
	}
	if IsPrime(38) == true {
		panic("FAILURE")
	}
	// TEST № 4
	if Increment("1101") == 14 {
		println("PASS")
	} else {
		panic("FAILURE")
	}
	// TEST № 5
	if Increment("11011") == 28 {
		println("PASS")
	} else {
		panic("FAILURE")
	}
	// TEST № 6
	if Increment("11101011") == 236 {
		println("PASS")
	} else {
		panic("FAILURE")
	}
	// Тести для IsBinaryPalindrome
	fmt.Println(IsBinaryPalindrome(7)) // true (111)
	fmt.Println(IsBinaryPalindrome(5)) // true (101)
	fmt.Println(IsBinaryPalindrome(6)) // false (110)

	// Тести для ValidParentheses
	fmt.Println(ValidParentheses("[{}]"))   // true
	fmt.Println(ValidParentheses("[{]}"))   // false
	fmt.Println(ValidParentheses("([]){}")) // true
	fmt.Println(ValidParentheses("(([]"))   // false

}
