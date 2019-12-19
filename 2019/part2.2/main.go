package main

import (
	"fmt"
	"time"
	"log"
)

const debug = false

func main() {
	const challenge = "Day2.2"
	const outputGoal = 19690720

	if debug {
		fmt.Println(challenge)
	}
	
	initProgram := []int{1,12,2,3,1,1,2,3,1,3,4,3,1,5,0,3,2,1,10,19,2,9,19,23,2,23,10,27,1,6,27,31,1,31,6,35,2,35,10,39,1,39,5,43,2,6,43,47,2,47,10,51,1,51,6,55,1,55,6,59,1,9,59,63,1,63,9,67,1,67,6,71,2,71,13,75,1,75,5,79,1,79,9,83,2,6,83,87,1,87,5,91,2,6,91,95,1,95,9,99,2,6,99,103,1,5,103,107,1,6,107,111,1,111,10,115,2,115,13,119,1,119,6,123,1,123,2,127,1,127,5,0,99,2,14,0,0}
	start := time.Now()
	
	
	for noun := 0; noun <= 99; noun += 1 {
	
		for verb := 0; verb <= 99; verb += 1 {
	
			var program []int
			program = make([]int, 500, 500)
			program = append(initProgram[:0:0], initProgram...) 
			program[1] = noun
			program[2] = verb
			if debug {
				fmt.Printf("reset with noun:%d verb:%d\n", noun, verb)
			}
	
			for i := 0; i < len(program); i += 1 {
	
				switch program[i] {
    					case 1:
						executeOne(&i, program)
    					case 2:
						executeTwo(&i, program)
    					case 99:
						if debug {
							fmt.Println("99")
							fmt.Println(program[0])
						}
						if program[0] == outputGoal {
							fmt.Printf("noun: %d, verb: %d \n", noun, verb)
							goto done
						}
						break
					default:
						break

    				}
			}
		}
	}

done:
	
	elapsed := time.Since(start)
	if debug {
		log.Printf("%s took %s", challenge, elapsed)
	}
}


func executeOne(i *int, program []int) {
	if debug {
        	fmt.Println("1")
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
	if debug {
        	fmt.Println("2")
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
