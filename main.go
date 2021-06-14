package main

import (
	"fmt"
	menu "github.com/NickLovera/go-apex/Menu"
	logic "github.com/NickLovera/go-apex/Mgr"
	"time"
)

var squadStats = logic.GetStats()
var lastUpdate = time.Now()

func main() {

	for {
		menu.PrintMenu(lastUpdate)
		choice := menu.GetChoice()

		if time.Since(lastUpdate).Minutes() > 5 {
			fmt.Println("Retrieving Stats......")
			squadStats = logic.GetStats()
			lastUpdate = time.Now()
		}

		switch choice {
		case 1:
			getEveryone()
		case 2:
			menu.PrintNames()
			playerId := menu.GetChoice()
			getIndiv(playerId)
		case 3:
			leaderBoard := getLeaderBord()
			menu.PrintContest(leaderBoard)
		}
	}
}

func getEveryone() {
	logic.GetEveryone(squadStats)
}

func getIndiv(playerId int) {
	logic.GetIndivdual(squadStats, playerId)
}

func getLeaderBord() [5]string {
	return logic.GetContestLeaderboard(squadStats)
}
