package main

import (
	"fmt"
	"log"
	"time"
)

const debug = false

	
func main() {
	const challenge = "Day4.1"

	const startPassword = 264793
	const endPassword = 803935
	
	validPasswords := 0

	if debug {
		fmt.Println(challenge)
	}
	
	start := time.Now()
	
	for i := startPassword; i <= endPassword; i++ {
		if debug {
	    		fmt.Printf("Current Password: %d\n",i)
		}

		passwordSlice := intToSlice(i, []int{})

		if !twoAdjacentDigits(passwordSlice) {
			continue
		}
		
		if !alwaysIncrease(passwordSlice) {
			continue
		}
		validPasswords+=1
	}
	
	fmt.Println(validPasswords)


	elapsed := time.Since(start)
	if debug {
		log.Printf("%s took %s", challenge, elapsed)
	}
}

func twoAdjacentDigits(passwordSlice []int) bool{
	for i := 0; i < len(passwordSlice) - 1; i++ {
		if passwordSlice[i] == passwordSlice[i+1] {
			if debug {
				fmt.Printf("2AD - first: %d second %d\n",passwordSlice[i],passwordSlice[i+1])
			}
			return true
		}
	}
	return false
}

func alwaysIncrease(passwordSlice []int) bool{
	for i := 0; i < len(passwordSlice) - 1; i++ {
		if passwordSlice[i] > passwordSlice[i+1] {
			if debug {
				fmt.Printf("AI - first: %d second %d\n",passwordSlice[i],passwordSlice[i+1])
			}
			return false
		}
	}
	return true
}


func intToSlice(n int, passwordSlice []int) []int {
	if n != 0 {
		i := n % 10
		passwordSlice = append([]int{i}, passwordSlice...)
		return intToSlice(n/10, passwordSlice)
	}
	return passwordSlice
}
