package timePackage

import (
	"fmt"
	"time"
)

func MainTimePackage() {
	now := time.Now()
	fmt.Println("Now:", now)

	nowLocal := now.Local()
	fmt.Println("Now Local:", nowLocal)

	utc := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	fmt.Println("UTC:", utc)
	fmt.Println("UTC Local:", utc.Local())

	formatter := "2006-01-02 15:04:05"
	valueTime, err := time.Parse(formatter, "2021-01-01 15:00:00")
	if err != nil {
		fmt.Println(err.Error())
		return
	} else {
		fmt.Println(valueTime)
	}

	var duration time.Duration = time.Second * 100
	fmt.Println("Duration in Second:", duration)
	var durationHour time.Duration = time.Hour * 2
	fmt.Println("Duration in Hour:", durationHour)

	oneHourFromNow := now.Add(time.Hour)

	fmt.Println("One Hour from Now:", oneHourFromNow)

	fmt.Println("Time Difference between One Hour from Now and Now:", oneHourFromNow.Sub(now))

	oneMonthFromNow := now.AddDate(0, 1, 0)

	fmt.Println("Time Difference between One Month from now and now : ", oneMonthFromNow.Sub(now))

}