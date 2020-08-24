package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

type result struct {
	team   string
	mp     int
	win    int
	draw   int
	loss   int
	points int
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
		if strings.HasPrefix(line, "#") || line == "" {
			continue
		}
		splittedLine := strings.Split(line, ";")
		if len(splittedLine) != 3 {
			return nil, errors.New("wrong input, should include the teams and the match result")
		}
		team1, team2, score := splittedLine[0], splittedLine[1], splittedLine[2]
		switch score {
		case "win":
			results[team1] = result{
				mp:     +1,
				win:    +1,
				points: +3,
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
				mp:     +1,
				win:    +1,
				points: +3,
			}
		case "draw":
			results[team1] = result{
				mp:     +1,
				draw:   +1,
				points: +1,
			}
			results[team2] = result{
				mp:     +1,
				draw:   +1,
				points: +1,
			}
		default:
			return nil, errors.New("The game result is invalid, should be: win, loss or draw")
		}

	}
	return results, nil
}

func createTable(teamsScores map[string]result, output io.Writer) error {

	table := make([]result, len(teamsScores))
	for team, score := range teamsScores {
		score.team = team
		table = append(table, score)
	}
	sort.Slice(table, func(i, j int) bool {
		return table[i].points > table[j].points
	})

	fmt.Println(table)

	output.Write([]byte("Team                           | MP |  W |  D |  L |  P\n"))
	for _, team := range table {
		line := fmt.Sprintf("%s                           | %d |  %d |  %d |  %d |  %d\n", team.team, team.mp, team.win, team.draw, team.loss, team.points)
		output.Write([]byte(line))
	}

	//for each team calculate the points and create a map[points]string (string is tha name of the team,
	//can be multiple teams on the same points)

	// write the wiriter in the right form

	return nil

}
