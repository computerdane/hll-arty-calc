package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var Version string

func main() {
	minDist := 100.
	maxDist := 1600.
	minAngle := 0.
	maxAngle := 0.
	dist := 0.
	input := ""
	interactive := true
	argsIndex := 0

	scanner := bufio.NewScanner(os.Stdin)

	// Use first argument as team name if given
	if len(os.Args) < 2 {
		fmt.Print("Enter a team ([b]ritain, [g]ermany, [r]ussia, [u]sa): ")
		scanner.Scan()
		input = scanner.Text()
	} else {
		input = os.Args[1]
		if input == "-v" || input == "--version" {
			fmt.Println(Version)
			os.Exit(0)
		}
	}

	// Passing distances as arguments disables interactive mode
	if len(os.Args) >= 3 {
		interactive = false
		argsIndex = 2
	}

	team := input[0]

	switch team {
	case 'b':
		minAngle = 267
		maxAngle = 533
	case 'g':
		minAngle = 622
		maxAngle = 978
	case 'r':
		minAngle = 800
		maxAngle = 1120
	case 'u':
		minAngle = 622
		maxAngle = 978
	default:
		fmt.Printf("Invalid team: %s\n", input)
		os.Exit(1)
	}

	//  maxAngle - angle      dist - minDist
	// ------------------- = -----------------, simplify to the form: angle = M * dist + B
	// maxAngle - minAngle   maxDist - minDist
	m := (minAngle - maxAngle) / (maxDist - minDist)
	b := (minDist * ((maxAngle - minAngle) / (maxDist - minDist))) + maxAngle

	for {
		if interactive {
			fmt.Print("Enter a distance in meters: ")

			scanner.Scan()
			input = scanner.Text()

			var err error
			dist, err = strconv.ParseFloat(input, 64)
			if err != nil {
				fmt.Printf("Invalid distance: %s\n", input)
				continue
			}
		} else {
			// Use arguments as distances for which we should calculate angles
			if argsIndex >= len(os.Args) {
				break
			}
			arg := os.Args[argsIndex]
			argsIndex++

			var err error
			dist, err = strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Printf("Invalid angle: %s\n", arg)
				break
			}
		}

		// The repo https://github.com/pastuh/hllminicalculator does this
		// They commented "random values fix (faction-gb)"
		// TODO: Figure out why this is needed
		if team == 'b' {
			if (dist >= 200 && dist <= 800) || (dist >= 1100 && dist <= 1200) {
				dist -= 5
			}
		}

		angle := m*dist + b

		fmt.Printf("%.0f\n", angle)
	}
}
