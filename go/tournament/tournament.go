package tournament

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"sort"
	"strings"
)

type teamScore struct {
	matchesPlayed int
	wins          int
	draws         int
}

func getLosses(score *teamScore) int {
	return score.matchesPlayed - score.wins - score.draws
}

func getPoints(score *teamScore) int {
	return 3*score.wins + score.draws
}

// Check that a team name is valid
func isValidTeam(team string) bool {
	teamNames := []string{
		"Allegoric Alaskians",
		"Blithering Badgers",
		"Courageous Californians",
		"Devastating Donkeys",
	}
	for _, name := range teamNames {
		if team == name {
			return true
		}
	}
	return false
}

func parseSingleTeam(score map[string]*teamScore, line string) (map[string]*teamScore, error) {
	line_split := strings.Split(line, ";")
	if len(line_split) < 3 {
		return score, errors.New("Invalid number of separators")
	}
	team_1, team_2, outcome := line_split[0], line_split[1], line_split[2]
	fmt.Printf("Team 1: %v, Team 2: %v, Outcome: %v\n", team_1, team_2, outcome)
	if !isValidTeam(team_1) || !isValidTeam(team_2) {
		return score, errors.New("Invalid team name")
	}
	team_1_matches := score[team_1].matchesPlayed
	team_2_matches := score[team_2].matchesPlayed
	score[team_1].matchesPlayed = team_1_matches + 1
	score[team_2].matchesPlayed = team_2_matches + 1

	team_1_wins := score[team_1].wins
	team_2_wins := score[team_1].wins
	team_1_draws := score[team_1].draws
	team_2_draws := score[team_1].draws
	if outcome == "win" {
		score[team_1].wins = team_1_wins + 1
	} else if outcome == "loss" {
		score[team_2].wins = team_2_wins + 1
	} else if outcome == "draw" {
		score[team_1].draws = team_1_draws + 1
		score[team_2].draws = team_2_draws + 1
	} else {
		return score, errors.New("Invalid outcome")
	}
	return score, nil
}

func Tally(r io.Reader, w io.Writer) error {
	// initialise mapping of team -> score
	allScores := map[string]*teamScore{}
	allScores["Allegoric Alaskians"] = &teamScore{}
	allScores["Blithering Badgers"] = &teamScore{}
	allScores["Courageous Californians"] = &teamScore{}
	allScores["Devastating Donkeys"] = &teamScore{}

	data, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	for _, line := range strings.Split(string(data), "\n") {
		if len(line) < 10 {
			continue
		}
		fmt.Printf("line: %v, length: %v\n", line, len(line))
		newScores, err := parseSingleTeam(allScores, line)
		if err != nil {
			return err
		}
		allScores = newScores
	}

	// Sort the teams/scores by the number of points won
	type kv struct {
		Key   string
		Value teamScore
	}
	var ss []kv
	for k, v := range allScores {
		ss = append(ss, kv{k, *v})
	}
	sort.Slice(ss, func(i, j int) bool {
		return getPoints(&ss[i].Value) > getPoints(&ss[j].Value)
	})

	fmt.Fprintf(w, "Team\t\t\t       | MP |  W |  D |  L |  P\n")
	for _, kv := range ss {
		team := kv.Key
		score := kv.Value
		losses := getLosses(&score)
		points := getPoints(&score)
		fmt.Fprintf(w, "%v\t       |  %v |  %v |  %v |  %v |  %v\n", team, score.matchesPlayed, score.wins, score.draws, losses, points)
	}
	// fmt.Fprintf(w, "\n")
	return nil
}
