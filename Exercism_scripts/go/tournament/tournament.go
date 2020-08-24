package tournament

import (
	"bufio"
	"errors"
	"io"
	"strings"
)

type result struct {
	mp   int
	win  int
	draw int
	loss int
}

//Tally reads the played matched and returns the tournament table
func Tally(input io.Reader, output io.Writer) error {

	teamsScores, err := read(input)
	if err != nil {
		return err
	}
	err = createTable(teamsScores, output)
	if err != nil {
		return err
	}

	return nil
}

func read(input io.Reader) (map[string]result, error) {
	scanner := bufio.NewScanner(input)
	results := make(map[string]result)

	//parse each line and populate the map/struct
	for scanner.Scan() {
		line := scanner.Text()
		//ignore comments and newlines
		if strings.HasPrefix(line, "#") || line == "\n" {
			continue
		}
		splittedLine := strings.Split(line, ";")
		team1, team2, score := splittedLine[0], splittedLine[1], splittedLine[2]
		switch score {
		case "win":
			results[team1] = result{
				mp:  +1,
				win: +1,
			}
			results[team2] = result{
				mp:   +1,
				loss: +1,
			}
		case "loss":
			results[team1] = result{
				mp:   +1,
				loss: +1,
			}
			results[team2] = result{
				mp:  +1,
				win: +1,
			}
		case "draw":
			results[team1] = result{
				mp:   +1,
				draw: +1,
			}
			results[team2] = result{
				mp:   +1,
				draw: +1,
			}
		default:
			return nil, errors.New("The game result is invalid, should be: win, loss or draw")
		}

	}
	return results, nil
}

func createTable(teamsScores map[string]result, output io.Writer) error {

	table := make(map[int][]string)
	for team, score := range teamsScores {
		points := 3*score.win + 1*score.draw
		table[points] = append(table[points], team)
	}
	//output.Write()

	//for each team calculate the points and create a map[points]string (string is tha name of the team,
	//can be multiple teams on the same points)

	// write the wiriter in the right form

	return nil

}
