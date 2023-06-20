/*
Copyright © 2023 yuanjun<simpleyuan@gmail.com>
*/
package mytime

import (
	"fmt"
	"strconv"
	"time"
)

func CurrentMondayDateTime() string {
	now := time.Now()
	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}
	monday := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	mondayStr := monday.Format(`2006-01-02`)
	return mondayStr + ` 00:00:00`
}

func NextMondayDateTime() string {
	now := time.Now()
	nextMonday := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, 7)
	nextMondayStr := nextMonday.Format(`2006-01-02`)
	return nextMondayStr + ` :00:00:00`
}

func CurrentWeekDays() (days []string) {
	now := time.Now()
	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}
	for i := 0; i < 7; i++ {
		weekday := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, i+offset)
		weekdayStr := weekday.Format(`2006-01-02`)
		days = append(days, weekdayStr)
	}
	return
}
func NextWeekDays() (days []string) {
	now := time.Now()
	for i := 7; i < 14; i++ {
		weekday := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, i)
		weekdayStr := weekday.Format(`2006-01-02`)
		days = append(days, weekdayStr)
	}
	return
}

func GetDateFromRange(start, end time.Time) (days []string) {
	// startTime, _ := time.Parse(`2006-01-02 15:04:05`, start)
	// endTime, _ := time.Parse(`2006-01-02 15:04:05`, end)
	startUnix := start.Unix()
	endUnix := end.Unix()
	if endUnix < startUnix {
		return
	}
	count := int((endUnix-startUnix)/86400) + 1
	for i := 0; i < count; i++ {
		_day := time.Date(start.Year(), start.Month(), start.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, i)
		days = append(days, _day.Format(`2006-01-02`))
	}
	return
}

func GetWeekDayFromDate(days []string) (wday []string) {
	if len(days) <= 0 {
		return
	}
	weekDayMap := map[string]int{
		`Monday`:    1,
		`Tuesday`:   2,
		`Wednesday`: 3,
		`Thursday`:  4,
		`Friday`:    5,
		`Saturday`:  6,
		`Sunday`:    7,
	}
	for _, v := range days {
		t, err := time.Parse(`2006-01-02`, v)
		if err != nil {
			return
		}
		wday = append(wday, fmt.Sprintf(`%d`, weekDayMap[t.Weekday().String()]))
	}
	return
}

func GetWeekDay(dayTime time.Time) int {
	weekDay := dayTime.Weekday()
	if weekDay == 0 {
		return 6
	} else if weekDay >= 1 && weekDay <= 6 {
		return int(weekDay) - 1
	}
	return int(weekDay)
}

func GetMondayOffsetSecond(week int, delay int) (res int) {

	now := time.Now()
	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}
	monday := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	res = monday.Second() + (week-1)*24*3600 + delay
	return res
}

func GenVarRunTimeStr(start string, offset string) string {
	_start, err := time.Parse(`2006-01-02 15:04:05`, start)
	_offset, err := strconv.Atoi(offset)
	startUnix := _start.Unix() + int64(_offset)
	startTime := time.Unix(startUnix, 0)
	startStr := startTime.Format(`2006-01-02 15:04:05`)
	if err != nil {
		return ``
	}
	return startStr
}

func GetCurrentDateTimeStr() string {
	now := time.Now()
	return now.Format(`2006-01-02 15:04:05`)
}

func GetDateTimeStrFromUnix(unix int64) string {
	timeUnix := time.Unix(unix, 0)
	timeStr := timeUnix.Format(`2006-01-02 15:04:05`)
	return timeStr
}
