package main

import "fmt"

func recoverExec() {
	if r := recover(); r != nil {
		fmt.Println("Try to recover from panic")
	}
}

func studentApprove(n1, n2 float64) bool {

	defer recoverExec()

	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("Media it 6!")
}

func main() {
	fmt.Println(studentApprove(8, 6))
	fmt.Println("After exec")
}
