package Menu

import (
	"bufio"
	"fmt"
	logic "github.com/NickLovera/go-apex/Mgr"
	"log"
	"os"
	"strconv"
	"time"
)

func PrintMenu(lastUpdate time.Time) {
	var MENU = "---- Welcome to YMAH stat tracker ----"
	var OPTIONS = "What would you like to do\n" + "1. Get Everyone's Stats\n" + "2. Get individual squad member stat\n" + "3. Check current contest"

	min, sec := logic.GetTimeTillUpdate(lastUpdate.Add(time.Minute * 5))

	fmt.Println(MENU)
	fmt.Println("Time till next update: ", min, " Min ", sec, " Sec\n") //Implement timer
	fmt.Println(OPTIONS)
}

func GetChoice() int {
	scanner := bufio.NewScanner(os.Stdin)
	scanned := scanner.Scan()
	if !scanned {
		log.Fatalln("Unable to scan")
	}
	choice, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatalln(err)
	}

	return choice
}

func PrintNames() {
	fmt.Println("Who would you like to view?\n" +
		"1. Hk_Dingledorf\n2. Its_SkeetR\n3. MoneyManRex937\n4. SourMonkeyy\n5. Mr__Brightside")
}

func PrintContest(leaderBoard [5]string) {
	fmt.Println("----Current Contest is most Kills in a week----")
	days, hours, minutes := logic.GetTimeTillContestEnd()
	fmt.Println("Time remaining: ", days, " Days ", hours, " Hours ", minutes, " Minutes")
	for _, line := range leaderBoard {
		fmt.Println(line)
	}
	fmt.Println()
}
