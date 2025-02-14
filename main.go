package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	minDist := 100.
	maxDist := 1600.
	minAngle := 0.
	maxAngle := 0.
	dist := 0.
	input := ""
	interactive := true

	scanner := bufio.NewScanner(os.Stdin)

	// Use first argument as team name if given
	if len(os.Args) < 2 {
		fmt.Print("Enter a team ([b]ritain, [g]ermany, [r]ussia, [u]sa): ")
		scanner.Scan()
		input = scanner.Text()
	} else {
		input = os.Args[1]
	}

	// Use remaining arguments as distances for which we should calculate angles
	// Passing distances as arguments disables interactive mode
	dists := []float64{}
	if len(os.Args) >= 3 {
		interactive = false
		for _, arg := range os.Args[2:] {
			dist, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Printf("Invalid angle: %s\n", arg)
				continue
			}
			dists = append(dists, dist)
		}
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
	M := (minAngle - maxAngle) / (maxDist - minDist)
	B := (minDist * ((maxAngle - minAngle) / (maxDist - minDist))) + maxAngle

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
			// "Pop" an element from the beginning of the distances list
			if len(dists) == 0 {
				break
			}
			dist = dists[0]
			dists = dists[1:]
		}

		if team == 'b' {
			// The repo https://github.com/pastuh/hllminicalculator does this
			// They commented "random values fix (faction-gb)"
			// TODO: Figure out why this is needed
			if (dist >= 200 && dist <= 800) || (dist >= 1100 && dist <= 1200) {
				dist -= 5
			}
		}

		angle := M*dist + B

		fmt.Printf("%.0f\n", angle)
	}
}
