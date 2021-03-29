package tournament

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

type teamScore struct {
	matchesPlayed int
	wins          int
	draws         int
}

type allScores map[string]teamScore

func getLosses(score teamScore) int {
	return score.matchesPlayed - score.wins - score.draws
}

func getScore(score teamScore) int {
	return 3*score.wins + score.draws
}

// Check that a team name is valid
func checkValidTeam(team string) bool {
	teamNames := []string{"Devastating Donkeys",
		"Allegoric Alaskans",
		"Blithering Badgers",
		"Courageous Californians"}
	for _, name := range teamNames {
		if team == name {
			return true
		}
	}
	return false
}

func Tally(r io.Reader, w io.Writer) error {
	scores := make(allScores)
	line := fmt.Sprintln(r)

	line_split := strings.Split(line, ";")
	team_1 := line_split[0]
	team_2 := line_split[1]
	scores[team_1] = teamScore.matchesPlayed + 1
	scores[team_2] = teamScore.matchesPlayed + 1
	if checkValidTeam(team_1) || checkValidTeam(team_2) {
		return errors.New("Invalid team name")
	}
	outcome := line_split[2]
	if outcome == "win" {
		scores[team_1] = teamScore.wins + 1
	} else if outcome == "loss" {
		scores[team_2] = teamScore.wins + 1		
	} else if outcome == "draw" {
		scores[team_1] = teamScore.draws + 1		
		scores[team_2] = teamScore.draws + 1		
	} else {
		return errors.New("Invalid outcome %v", outcome)
	}
	fmt.Printf("Scores:\n%v", scores)
	return nil
}
