package lib

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

func getEquationConstants(team *Team) (m float64, b float64) {
	//  maxAngle - angle      dist - minDist
	// ------------------- = -----------------, simplify to the form: angle = M * dist + B
	// maxAngle - minAngle   maxDist - minDist
	m = (team.MinAngle - team.MaxAngle) / (maxDist - minDist)
	b = (minDist * ((team.MaxAngle - team.MinAngle) / (maxDist - minDist))) + team.MaxAngle
	return m, b
}

func GetAngle(team *Team, dist float64) float64 {
	m, b := getEquationConstants(team)

	// The repo https://github.com/pastuh/hllminicalculator does this
	// They commented "random values fix (faction-gb)"
	// TODO: Figure out why this is needed
	if team.Name == "Britain" {
		if (dist >= 200 && dist <= 800) || (dist >= 1100 && dist <= 1200) {
			dist -= 5
		}
	}

	return m*dist + b
}
