package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

// Score stores stats for a team
type Score struct {
	name   string
	wins   int
	draws  int
	losses int
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
func ParseSingleLine(scores map[string]Score, line string) error {
	lineSplit := strings.Split(line, ";")
	if len(lineSplit) < 3 {
		return errors.New("invalid number of separators")
	}
	team1, team2, outcome := lineSplit[0], lineSplit[1], lineSplit[2]

	a, b := scores[team1], scores[team2]
	a.name, b.name = team1, team2
	switch outcome {
	case "win":
		a.wins++
		b.losses++
	case "loss":
		a.losses++
		b.wins++
	case "draw":
		a.draws++
		b.draws++
	default:
		return errors.New("invalid outcome")
	}
	scores[team1], scores[team2] = a, b
	return nil
}

// Tally parses in team score input as string into a leaderboard
func Tally(r io.Reader, w io.Writer) error {
	allScores := make(map[string]Score)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()

		// Ignore comment lines
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		err := ParseSingleLine(allScores, line)
		if err != nil {
			return err
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	// To sort the map by number of points, we must create an array
	var SortedScores []Score
	for _, score := range allScores {
		SortedScores = append(SortedScores, score)
	}
	// Sort by points, if equal points, sort by name ascending
	sort.Slice(SortedScores, func(i, j int) bool {
		if SortedScores[i].Points() == SortedScores[j].Points() {
			return SortedScores[i].name < SortedScores[j].name
		}
		return SortedScores[i].Points() > SortedScores[j].Points()
	})

	fmt.Fprintf(w, `%-31v| %v |  %v |  %v |  %v |  %v`, "Team", "MP", "W", "D", "L", "P")
	for _, ts := range SortedScores {
		fmt.Fprintf(w, "\n")
		fmt.Fprintf(w, `%-31v|  %v |  %v |  %v |  %v |  %v`, ts.name, ts.MatchesPlayed(), ts.wins, ts.draws, ts.losses, ts.Points())
	}
	fmt.Fprintf(w, "\n")
	return nil
}
