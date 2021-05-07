package tournament

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"sort"
	"strings"
)

// Score stores stats for a team
type Score struct {
	wins   int
	draws  int
	losses int
}

// TeamScore associates scores with a team
type TeamScore struct {
	team  string
	score Score
}

// validOutcomes contains the possible outcomes for a game
var validOutcomes = map[string]bool{"win": true, "loss": true, "draw": true}

// Update the score for a team given an outcome
func (score *Score) Update(field string) {
	if field == "win" {
		score.wins++
	} else if field == "draw" {
		score.draws++
	} else if field == "loss" {
		score.losses++
	}
}

// MatchesPlayed calculates the number of games played by a team
func (score *Score) MatchesPlayed() int {
	return score.wins + score.draws + score.losses
}

// Points calculates the total score value for a team
func (score *Score) Points() int {
	return 3*score.wins + score.draws
}

// ParseSingleLine adds the data from a single input line to the overall score mapping
func ParseSingleLine(scores map[string]*Score, line string) error {
	lineSplit := strings.Split(line, ";")
	if len(lineSplit) < 3 {
		return errors.New("invalid number of separators")
	}
	team1, team2, outcome := lineSplit[0], lineSplit[1], lineSplit[2]
	if _, ok := scores[team1]; !ok {
		return errors.New("invalid team name")
	} else if _, ok := scores[team2]; !ok {
		return errors.New("invalid team name")
	}
	if _, ok := validOutcomes[outcome]; !ok {
		return errors.New("invalid outcome")
	}

	scores[team1].Update(outcome)
	if outcome == "win" {
		scores[team2].Update("loss")
	} else if outcome == "loss" {
		scores[team2].Update("win")
	} else {
		scores[team2].Update(outcome)
	}
	return nil
}

// Tally parses in team score input as string into a leaderboard
func Tally(r io.Reader, w io.Writer) error {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	allScores := map[string]*Score{"Allegoric Alaskians": {}, "Blithering Badgers": {}, "Courageous Californians": {}, "Devastating Donkeys": {}}

	for _, line := range strings.Split(string(data), "\n") {
		// Ignore comment lines
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		err = ParseSingleLine(allScores, line)
		if err != nil {
			return err
		}
	}

	// To sort the map by number of points, we must create an array
	var SortedScores []TeamScore
	for team, score := range allScores {
		SortedScores = append(SortedScores, TeamScore{team, *score})
	}
	// Sort by points, if equal points, sort by name ascending
	sort.Slice(SortedScores, func(i, j int) bool {
		if SortedScores[i].score.Points() == SortedScores[j].score.Points() {
			return SortedScores[i].team < SortedScores[j].team
		}
		return SortedScores[i].score.Points() > SortedScores[j].score.Points()
	})

	fmt.Fprintf(w, `%-31v| %v |  %v |  %v |  %v |  %v`, "Team", "MP", "W", "D", "L", "P")
	for _, ts := range SortedScores {
		fmt.Fprintf(w, "\n")
		fmt.Fprintf(w, `%-31v|  %v |  %v |  %v |  %v |  %v`, ts.team, ts.score.MatchesPlayed(), ts.score.wins, ts.score.draws, ts.score.losses, ts.score.Points())
	}
	fmt.Fprintf(w, "\n")
	return nil
}
