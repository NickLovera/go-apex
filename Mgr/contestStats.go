package Mgr

import (
	"strconv"
	"time"
)

//GetStartingStats returns the starting stats for all 5 members
//Chnage these values to the total starting count for any stat
//"HK_Dingledorf", "Its_SkeetR", "MoneyManRex937", "SourMonkeyy", "Mr__Briteside"
func GetStartingStats() [5]int {
	return [5]int{9759658, 1508048, 6251766, 2673304, 1564391}
}

func GetEndTime() time.Time {
	/**/
	//fmt.Println(time.Now().AddDate(0, 0, 7)) //Used for getting end time of new event
	location := time.FixedZone("UTC", -5*60*60)
	return time.Date(2021, 06, 29, 18, 0, 0, 0, location)
}

func GetTimeTillContestEnd() (int, int, int) {
	endTime := GetEndTime()
	currentTime := time.Now()
	difference := endTime.Sub(currentTime)

	total := int(difference.Seconds())
	days := int(total / (60 * 60 * 24))
	hours := int(total / (60 * 60) % 24)
	minutes := int(total/60) % 60

	return days, hours, minutes
}

func GetContestLeaderboard(squadStats [5]Result) [5]string {

	var unsortedBoard [5][2]int

	for playerId, squadMember := range squadStats {
		unsortedBoard[playerId][0] = playerId
		for _, legend := range squadMember.Legends {
			unsortedBoard[playerId][1] += int(legend.Stats.Damages.Value)
		}
		/*Comment out when starting new contest*/
		unsortedBoard[playerId][1] -= GetStartingStats()[playerId]
	}
	sortedBoard := sortBoard(unsortedBoard)
	leaderBoard := createLeaderBoard(sortedBoard)
	/*Used to get starting stats Make sure to comment out above ^*/
	//fmt.Println(unsortedBoard)
	return leaderBoard
}

func sortBoard(unsortedBoard [5][2]int) [5][2]int {
	for i := 0; i < len(unsortedBoard)-1; i++ {
		for i := 0; i < len(unsortedBoard)-1; i++ {
			if unsortedBoard[i][1] < unsortedBoard[i+1][1] {
				temp := unsortedBoard[i]
				unsortedBoard[i] = unsortedBoard[i+1]
				unsortedBoard[i+1] = temp
			}
		}
	}
	return unsortedBoard
}

func createLeaderBoard(sortedBoard [5][2]int) [5]string {
	var leaderBoard [5]string
	for i := 0; i < 5; i++ {
		switch i {
		case 0:
			leaderBoard[i] = "1st: " + squad[sortedBoard[i][0]] + " Damage: " + strconv.Itoa(sortedBoard[i][1])
		case 1:
			leaderBoard[i] = "2nd: " + squad[sortedBoard[i][0]] + " Damage: " + strconv.Itoa(sortedBoard[i][1])
		case 2:
			leaderBoard[i] = "3rd: " + squad[sortedBoard[i][0]] + " Damage: " + strconv.Itoa(sortedBoard[i][1])
		case 3:
			leaderBoard[i] = "4th: " + squad[sortedBoard[i][0]] + " Damage: " + strconv.Itoa(sortedBoard[i][1])
		case 4:
			leaderBoard[i] = "5th: " + squad[sortedBoard[i][0]] + " Damage: " + strconv.Itoa(sortedBoard[i][1])
		}
	}
	return leaderBoard
}