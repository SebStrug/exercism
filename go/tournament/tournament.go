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
	matchesPlayed int
	wins          int
	draws         int
	losses        int
	points        int
}

// TeamScore associates scores with a team
type TeamScore struct {
	team  string
	score Score
}

// AllTeamScores contains scores for each team
var AllTeamScores = map[string]*Score{
	"Allegoric Alaskians":     {},
	"Blithering Badgers":      {},
	"Courageous Californians": {},
	"Devastating Donkeys":     {},
}

// GetPoints calculates the total points for a team
func GetPoints(score *Score) int {
	return 3*score.wins + score.draws
}

// Increment updates one of the stats in a score
func (score *Score) Increment(field string) {
	if field == "matches" {
		score.matchesPlayed++
	} else if field == "wins" {
		score.wins++
	} else if field == "draws" {
		score.draws++
	} else if field == "losses" {
		score.losses++
	}
}

// ParseSingleLine adds the data from a single input line to the overall score mapping
func ParseSingleLine(scores map[string]*Score, line string) (map[string]*Score, error) {
	lineSplit := strings.Split(line, ";")
	if len(lineSplit) < 3 {
		return scores, errors.New("invalid number of separators")
	}
	team1, team2, outcome := lineSplit[0], lineSplit[1], lineSplit[2]
	// fmt.Printf("Team 1: %v, Team 2: %v, Outcome: %v\n", team1, team2, outcome)

	if _, ok := AllTeamScores[team1]; !ok {
		return scores, errors.New("invalid team name")
	} else if _, ok := AllTeamScores[team2]; !ok {
		return scores, errors.New("invalid team name")
	}

	scores[team1].Increment("matches")
	scores[team2].Increment("matches")

	if outcome == "win" {
		scores[team1].Increment("wins")
		scores[team2].Increment("losses")
	} else if outcome == "loss" {
		scores[team2].Increment("wins")
		scores[team1].Increment("losses")
	} else if outcome == "draw" {
		scores[team1].Increment("draws")
		scores[team2].Increment("draws")
	} else {
		return scores, errors.New("invalid outcome")
	}

	return scores, nil
}

// Tally parses in team score input as string into a leaderboard
func Tally(r io.Reader, w io.Writer) error {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	for _, line := range strings.Split(string(data), "\n") {
		// Ignore invalid lines, comment lines
		if len(line) < 10 || string(line[0]) == "#" {
			continue
		}

		// fmt.Printf("line: %v, length: %v\n", line, len(line))
		newScores, err := ParseSingleLine(AllTeamScores, line)
		if err != nil {
			return err
		}
		AllTeamScores = newScores
	}
	for _, score := range AllTeamScores {
		score.points = GetPoints(score)
	}

	// To sort the map by number of points, we must create an array
	var SortedScores []TeamScore
	for team, score := range AllTeamScores {
		SortedScores = append(SortedScores, TeamScore{team, *score})
	}
	sort.Slice(SortedScores, func(i, j int) bool {
		return SortedScores[i].score.points > SortedScores[j].score.points
	})

	fmt.Fprintf(w, `Team                           | MP |  W |  D |  L |  P
`)
	for ind, ts := range SortedScores {
		if ind == len(SortedScores) -1 {
			fmt.Fprintf(w, `%-31v|  %v |  %v |  %v |  %v |  %v
`, ts.team, ts.score.matchesPlayed, ts.score.wins, ts.score.draws, ts.score.losses, ts.score.points)
		} else {
			fmt.Fprintf(w, `%-31v|  %v |  %v |  %v |  %v |  %v
`, ts.team, ts.score.matchesPlayed, ts.score.wins, ts.score.draws, ts.score.losses, ts.score.points)
		}
	}
	return nil
}
