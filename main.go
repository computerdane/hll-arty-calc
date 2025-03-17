package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var Version string

const minDist = 100.
const maxDist = 1600.

type Team struct {
	Name     string
	MinAngle float64
	MaxAngle float64
}

var Britain = Team{
	Name:     "Britain",
	MinAngle: 267,
	MaxAngle: 533,
}

var Germany = Team{
	Name:     "Germany",
	MinAngle: 622,
	MaxAngle: 978,
}

var Russia = Team{
	Name:     "Russia",
	MinAngle: 800,
	MaxAngle: 1120,
}

var Usa = Team{
	Name:     "Usa",
	MinAngle: 622,
	MaxAngle: 978,
}

func parseTeam(input string) (team *Team, err error) {
	c := input[0]
	switch c {
	case 'b':
		return &Britain, nil
	case 'g':
		return &Germany, nil
	case 'r':
		return &Russia, nil
	case 'u':
		return &Usa, nil
	default:
		return nil, fmt.Errorf("Invalid team: %s\n", input)
	}
}

func getEquationConstants(team *Team) (m float64, b float64) {
	//  maxAngle - angle      dist - minDist
	// ------------------- = -----------------, simplify to the form: angle = M * dist + B
	// maxAngle - minAngle   maxDist - minDist
	m = (team.MinAngle - team.MaxAngle) / (maxDist - minDist)
	b = (minDist * ((team.MaxAngle - team.MinAngle) / (maxDist - minDist))) + team.MaxAngle
	return m, b
}

func main() {
	var err error

	var team *Team

	dist := 0.
	input := ""
	interactive := true
	argsIndex := 0

	scanner := bufio.NewScanner(os.Stdin)

	if len(os.Args) < 2 {
		// Prompt for team name if none provided
		for team == nil || err != nil {
			fmt.Print("Enter a team ([b]ritain, [g]ermany, [r]ussia, [u]sa): ")
			scanner.Scan()
			input = scanner.Text()
			team, err = parseTeam(input)
			if err != nil {
				fmt.Println(err)
			}
		}
	} else {
		// Use first argument as team name if provided
		input = os.Args[1]
		if input == "-v" || input == "--version" {
			fmt.Println(Version)
			os.Exit(0)
		}
		team, err = parseTeam(input)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Passing distances as arguments disables interactive mode
	if len(os.Args) >= 3 {
		interactive = false
		argsIndex = 2
	}

	m, b := getEquationConstants(team)

	for {
		if interactive {
			fmt.Printf("[%s] Enter a distance in meters: ", team.Name)

			scanner.Scan()
			input = scanner.Text()

			var err error
			dist, err = strconv.ParseFloat(input, 64)
			if err != nil {
				_team, err := parseTeam(input)
				if err != nil {
					fmt.Printf("Invalid distance: %s\n", input)
				} else {
					team = _team
					m, b = getEquationConstants(team)
				}

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
		if team.Name == "Britain" {
			if (dist >= 200 && dist <= 800) || (dist >= 1100 && dist <= 1200) {
				dist -= 5
			}
		}

		angle := m*dist + b

		fmt.Printf("%.0f\n", angle)
	}
}
