package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		result, err := interpret(scanner.Text())
		if err == nil {
			fmt.Print(result)
		} else {
			fmt.Println("An error occurred while interpreting the commands:")
			fmt.Println(err)
		}
	} else {
		fmt.Println("Unable to read the commands")
	}
}

func interpret(input string) (string, error) {
	var (
		err      error
		memory   [30000]byte
		cell     int
		output   string
		brackets int
	)

	commands := strings.Split(input, "")
	length := len(commands)

	for i := 0; i < length; i++ {
		switch commands[i] {
		case ">":
			cell++
		case "<":
			cell--
		case "+":
			memory[cell]++
		case "-":
			memory[cell]--
		case ".":
			output += string(memory[cell])
		case ",":
			_, err := fmt.Scanf("%d", memory[cell])
			if err != nil {
				return "", err
			}
		case "[":
			if memory[cell] == 0 {
				brackets++
				for brackets != 0 {
					i++
					if i >= length {
						return "", errors.New("square brackets mismatch")
					}
					if commands[i] == "[" {
						brackets++
					}
					if commands[i] == "]" {
						brackets--
					}
				}
			}
		case "]":
			if memory[cell] != 0 {
				brackets++
				for brackets != 0 {
					i--
					if commands[i] == "[" {
						brackets--
					}
					if commands[i] == "]" {
						brackets++
					}
				}
				i--
			}
		default:
			return "", fmt.Errorf("unsupported command detected %s, symbol #%d", commands[i], i)
		}
	}

	return output, err
}
