package main

import (
	"fmt"
	"strings"
)

/*
Tally the results of a small football competition.

Tally the results of a small football competition. Based on an input file containing which team played against which and what the outcome was create a file with a table like this:

Team                           | MP |  W |  D |  L |  P
Devastating Donkeys            |  3 |  2 |  1 |  0 |  7
Allegoric Alaskans             |  3 |  2 |  0 |  1 |  6
Blithering Badgers             |  3 |  1 |  0 |  2 |  3
Courageous Californians        |  3 |  0 |  1 |  2 |  1
What do those abbreviations mean?

MP: Matches Played
W: Matches Won
D: Matches Drawn (Tied)
L: Matches Lost
P: Points
A win earns a team 3 points. A draw earns 1. A loss earns 0.

The outcome should be ordered by points, descending. In case of a tie, teams are ordered alphabetically.

Input

Your tallying program will receive input that looks like:

Allegoric Alaskans;Blithering Badgers;win
Devastating Donkeys;Courageous Californians;draw
Devastating Donkeys;Allegoric Alaskans;win
Courageous Californians;Blithering Badgers;loss
Blithering Badgers;Devastating Donkeys;loss
Allegoric Alaskans;Courageous Californians;win
The result of the match refers to the first team listed. So this line

Allegoric Alaskans;Blithering Badgers;win
Means that the Allegoric Alaskans beat the Blithering Badgers.

This line:

Courageous Californians;Blithering Badgers;loss
Means that the Blithering Badgers beat the Courageous Californians.

And this line:

Devastating Donkeys;Courageous Californians;draw
Means that the Devastating Donkeys and Courageous Californians tied.

Your program should only process input lines that follow this format. All other lines should be ignored: If an input contains both valid and invalid input lines, output a table that contains just the results from the valid lines.
*/
type Team struct {
	Name          string
	MatchesPlayed int
	MatchesWon    int
	MatchesDrawn  int
	MatchesLost   int
	Points        int
}

func TallyTheResults(gameHistory string) {
	teams, pWithTeams, maxPoint := make(map[string]*Team), make(map[int][]*Team), 0
	matches := strings.Split(gameHistory, "\n")
	for _, onceMatch := range matches {
		params := strings.Split(onceMatch, ";")
		if len(params) == 3 {
			team1, team2, result := params[0], params[1], params[2]
			if _, has := teams[team1]; !has {
				teams[team1] = &Team{Name: team1}
			}
			if _, has := teams[team2]; !has {
				teams[team2] = &Team{Name: team2}
			}
			switch result {
			case "win":
				teams[team1].MatchesWon++
				teams[team1].Points += 3
				teams[team2].MatchesLost++
			case "draw":
				teams[team1].MatchesDrawn++
				teams[team2].MatchesDrawn++
				teams[team1].Points++
				teams[team2].Points++
			case "loss":
				teams[team2].MatchesWon++
				teams[team2].Points += 3
				teams[team1].MatchesLost++
			}
			teams[team1].MatchesPlayed++
			teams[team2].MatchesPlayed++
		}
	}
	for _, team := range teams {
		if team.Points > maxPoint {
			maxPoint = team.Points
		}
		if pWithTeams[team.Points] == nil {
			pWithTeams[team.Points] = []*Team{team}
		} else {
			pWithTeams[team.Points] = insertTeam(team, pWithTeams[team.Points])
		}
	}
	fmt.Println("Team|MP|W|D|L|P")
	for i := maxPoint; i >= 0; i-- {
		if pWithTeams[i] != nil {
			for _, team := range pWithTeams[i] {
				fmt.Print(team.Name + "|")
				fmt.Print(team.MatchesPlayed)
				fmt.Print("|")
				fmt.Print(team.MatchesWon)
				fmt.Print("|")
				fmt.Print(team.MatchesDrawn)
				fmt.Print("|")
				fmt.Print(team.MatchesLost)
				fmt.Print("|")
				fmt.Printf("%d|\n", team.Points)
			}
		}
	}
}
func insertTeam(tam *Team, teams []*Team) (result []*Team) {
	result = append(teams, tam)
	return result
}
func main() {
	TallyTheResults(`
Allegoric Alaskans;Blithering Badgers;win
Devastating Donkeys;Courageous Californians;draw
Devastating Donkeys;Allegoric Alaskans;win
Courageous Californians;Blithering Badgers;loss
Blithering Badgers;Devastating Donkeys;loss
Allegoric Alaskans;Courageous Californians;win
`)
}
