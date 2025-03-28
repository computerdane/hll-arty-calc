package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/computerdane/hll-arty-calc/lib"
)

var Version string

func parseTeam(input string) (team *lib.Team, err error) {
	c := input[0]
	switch c {
	case 'b':
		return &lib.Britain, nil
	case 'g':
		return &lib.Germany, nil
	case 'r':
		return &lib.Russia, nil
	case 'u':
		return &lib.Usa, nil
	default:
		return nil, fmt.Errorf("Invalid team: %s\n", input)
	}
}

func main() {
	var err error

	var team *lib.Team

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

		angle := lib.GetAngle(team, dist)

		fmt.Printf("%.0f\n", angle)
	}
}
