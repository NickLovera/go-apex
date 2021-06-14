package Mgr

import (
	"time"
)

//GetStartingStats returns the starting stats for all 5 members
//Chnage these values to the total starting count for any stat
func GetStartingStats() [5]int {
	return [5]int{41500, 4000, 21000, 12000, 6000}
}

func GetEndTime() time.Time {
	//fmt.Println(time.Now().AddDate(0, 0, 7)) //Used for getting end time of new event
	location := time.FixedZone("UTC", -5*60*60)
	return time.Date(2021, 06, 21, 14, 23, 9, 0500, location)
}
