package tournament

import (
	"errors"
	"strings"
)

// Score stores stats for a team
type Score struct {
	name          string
	matchesPlayed int
	wins          int
	draws         int
	losses        int
	points        int
}

// Increment updates one of the stats in a score
func (score *Score) Update(field string) {
	if field == "win" {
		score.wins++
	} else if field == "draw" {
		score.draws++
	} else if field == "loss" {
		score.losses++
	}
}

func (score *Score) MatchesPlayed() int {
	return score.wins + score.draws + score.losses
}

func (score *Score) Points() int {
	return 3*score.wins + score.draws
}

func UpdateScores(scores []*Score, team1 string, team2 string, outcome string) {
	team1Ind := -1
	team2Ind := -2
	for ind, s := range scores {
		if s.name == team1 {
			team1Ind = ind
		} else if s.name == team2 {
			team2Ind = ind
		}
	}

	if team1Ind == -1 {
		scores = append(scores, &Score{name: team1})
	} else {
		scores[team1Ind].Update(outcome)
	}
	if team2Ind == -1 {
		scores = append(scores, &Score{name: team2})
	} else {
		scores[team2Ind].Update(outcome)
	}
}

func ParseSingleLine(scores []*Score, line string) error {
	lineSplit := strings.Split(line, ";")
	if len(lineSplit) < 3 {
		return errors.New("invalid number of separators")
	}
	team1, team2, outcome := lineSplit[0], lineSplit[1], lineSplit[2]

}
