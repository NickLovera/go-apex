package Mgr

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

//Result contains entire json
type Result struct {
	Legends []Legend `json:"data"`
}

//Legend Individual Fields in the json
type Legend struct {
	//Type string `json:"type"`
	//Attributes Attribute `json:"attributes"`
	Metas MetaData `json:"metadata"`
	//Dates string `json:"expiryDate"`
	Stats Stat `json:"stats"`
}

//Attribute contains legend id (Will use for searching legend in future)
type Attribute struct {
	LegendId string `json:"id"`
}

//MetaData name is the legends name ie. bloodhound
type MetaData struct {
	Name string `json:"name"`
}

//Stat Each field is a stat ie. kills, winningKills, .....
type Stat struct {
	KillNum   Kills    `json:"kills"`
	Damages   Damage   `json:"damage"`
	Headshots Headshot `json:"headshots"`
}

//Kills
type Kills struct {
	Rank         float32 `json:"rank"`
	DisplayValue string  `json:"displayValue"`
	Value        float32 `json:"value"`
}

type Damage struct {
	Rank         float32 `json:"rank"`
	DisplayValue string  `json:"displayValue"`
	Value        float32 `json:"value"`
}

type Headshot struct {
	Rank         float32 `json:"rank"`
	DisplayValue string  `json:"displayValue"`
	Value        float32 `json:"value"`
}

var squad = [5]string{"HK_Dingledorf", "Its_SkeetR", "MoneyManRex937", "SourMonkeyy", "Mr__Briteside"}

func GetStats() [5]Result {
	var squadStats [5]Result

	for i := 0; i < 5; i++ {
		req, err := http.NewRequest("GET", "https://public-api.tracker.gg/v2/apex/standard/profile/psn/"+squad[i]+"/segments/legend", nil)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		req.Header.Set("TRN-Api-Key", "cf97cbea-dfd7-46f0-aa43-5acc8da4e47c")
		client := &http.Client{}

		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error getting response: ", err)
			os.Exit(1)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading body: ", err)
		}

		var stats Result
		err = json.Unmarshal(body, &stats)
		if err != nil {
			fmt.Println("Error Unmarshalling: ", err)
		}
		//Store each bois stats
		squadStats[i] = stats
	}
	return squadStats
}

//GetEveryone Get all stats of everyone
func GetEveryone(squadStats [5]Result) {

	for playerId, squadMember := range squadStats {
		showStats(playerId, squadMember)
	}
}

func showStats(playerId int, squadMember Result) {
	file, err := os.Create("/workspace/go-apex/Data/" + squad[playerId])
	if err != nil {
		fmt.Println("Error opening write file: ", err)
	}

	file.WriteString("------------------------------------------------------\n")
	fmt.Println("------------------------------------------------------")
	fmt.Println(squad[playerId], " stat's")
	file.WriteString(squad[playerId] + " stat's" + "\n\n")
	for _, legend := range squadMember.Legends {
		fmt.Println("Legend: ", legend.Metas.Name)
		file.WriteString("Legend: " + legend.Metas.Name + "\n")

		fmt.Println("Kills: ", legend.Stats.KillNum.DisplayValue, " Rank: ", legend.Stats.KillNum.Rank)
		file.WriteString("Kills: " + legend.Stats.KillNum.DisplayValue + " Rank: " + fmt.Sprintf("%g", legend.Stats.KillNum.Rank) + "\n")

		fmt.Println("Damage: ", legend.Stats.Damages.DisplayValue, " Rank: ", legend.Stats.Damages.Rank)
		file.WriteString("Damage: " + legend.Stats.Damages.DisplayValue + " Rank: " + fmt.Sprintf("%g", legend.Stats.Damages.Rank) + "\n")

		fmt.Println("Headshots: ", legend.Stats.Headshots.DisplayValue, " Rank: "+fmt.Sprintf("%g", legend.Stats.Headshots.Rank))
		file.WriteString("Headshots: " + legend.Stats.Headshots.DisplayValue + " Rank: " + fmt.Sprintf("%g", legend.Stats.Headshots.Rank) + "\n\n")

		fmt.Println()
	}
	fmt.Println("------------------------------------------------------")
	file.WriteString("------------------------------------------------------\n")

	err = file.Close()
	if err != nil {
		fmt.Println(err)
	}
}

//Get one person's stats
func GetIndivdual(squadStats [5]Result, playerId int) {
	showStats(playerId-1, squadStats[playerId-1])
}

func GetTimeTillUpdate(lastUpdate time.Time) (int, int) {
	currentTime := time.Now()
	difference := lastUpdate.Sub(currentTime)

	total := int(difference.Seconds())
	minutes := int(total/60) % 60
	seconds := total % 60

	return minutes, seconds
}