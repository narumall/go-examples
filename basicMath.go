package main

import "fmt"

func isEven(n int) bool {
	if n%2 == 0 {
		return true

	} else {
		return false
	}
}

func isPrime(n int) bool {
	isPrimeValue := true
	for i := 2; i < n/2; i++ {
		if n%i == 0 {
			isPrimeValue = false
			break
		}
	}
	return isPrimeValue
}

func main() {
	fmt.Println("Enter the value: ")
	var n int
	fmt.Scanln(&n)
	fmt.Println("is Even ==> ", isEven(n))
	fmt.Println("is Prime => ", isPrime(n))
}
