package main

import (
	"fmt"
	"os"
	"time"
	"log"
)

const debug = true

func main() {
	const challenge = "Day2"

	if debug {
		fmt.Println(challenge)
	}
	
	program := []int{1,12,2,3,1,1,2,3,1,3,4,3,1,5,0,3,2,1,10,19,2,9,19,23,2,23,10,27,1,6,27,31,1,31,6,35,2,35,10,39,1,39,5,43,2,6,43,47,2,47,10,51,1,51,6,55,1,55,6,59,1,9,59,63,1,63,9,67,1,67,6,71,2,71,13,75,1,75,5,79,1,79,9,83,2,6,83,87,1,87,5,91,2,6,91,95,1,95,9,99,2,6,99,103,1,5,103,107,1,6,107,111,1,111,10,115,2,115,13,119,1,119,6,123,1,123,2,127,1,127,5,0,99,2,14,0,0}
	start := time.Now()
	
	for i := 0; i < len(program); i += 1 {
	
		switch program[i] {
    			case 1:
				executeOne(&i, program)
    			case 2:
				executeTwo(&i, program)
    			case 99:
				i = len(program) + 1
				break
			default:
				os.Exit(1)

    		}
	}
	
	elapsed := time.Since(start)
	if debug {
		log.Printf("%s took %s", challenge, elapsed)
	}
}


func executeOne(i *int, program []int) {
        fmt.Println("1")
	if debug {
		log.Printf("%v", program)
	}
	var (
		firstNumber = *i+1
		secondNumber = *i+2
		result = *i+3
	) 
	program[program[result]] = program[program[firstNumber]] + program[program[secondNumber]]
	if debug {
		log.Printf("%v", program)
	}
	(*i)+=3
}

func executeTwo(i *int, program []int) {
        fmt.Println("2")
	if debug {
		log.Printf("%v", program)
	}
	var (
		firstNumber = *i+1
		secondNumber = *i+2
		result = *i+3
	) 
	program[program[result]] = program[program[firstNumber]] * program[program[secondNumber]]
	if debug {
		log.Printf("%v", program)
	}
	(*i)+=3
}
