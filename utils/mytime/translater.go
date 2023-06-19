package mytime

import (
	"errors"
	"strings"
	"time"

	"github.com/spf13/cast"
)

func String2Int64(timeStr string) int64 {
	loc, _ := time.LoadLocation("Local")
	location, err := time.ParseInLocation("2006-01-02 15:04:05", timeStr, loc)
	if err != nil {
		return 0
	}
	return location.Unix()
}

func Unix2String(unix int64) string {
	timeTemplate := "2006-01-02 15:04:05"
	return time.Unix(unix, 0).Format(timeTemplate)
}

func TimeToSecond(timeStr string) (int64, error) {
	runAt := strings.Split(strings.TrimSpace(timeStr), ":")
	if len(runAt) != 2 {
		return 0, errors.New("invalid time")
	}
	hour := cast.ToInt64(strings.TrimPrefix(runAt[0], "0"))
	min := cast.ToInt64(strings.TrimPrefix(runAt[1], "0"))

	return hour*3600 + min*60, nil
}

func IsCurrentTimeInRange(start string, end string) bool {
	if start > end { //跨天
		start1 := `00:00`
		end1 := end
		if IsCurrentTimeInRangeMinAndMax(start1, end1) {
			return true
		}
		start2 := start
		end2 := `23:59`
		if IsCurrentTimeInRangeMinAndMax(start2, end2) {
			return true
		}
	}
	return IsCurrentTimeInRangeMinAndMax(start, end)
}

func IsCurrentTimeInRangeMinAndMax(min string, max string) bool {
	current := time.Now().Unix()
	min += `:00`
	max += `:59`
	today := time.Now().Format("2006-01-02")
	minTmp := today + ` ` + min
	maxTmp := today + ` ` + max

	loc, _ := time.LoadLocation("Local")
	minUnix, _ := time.ParseInLocation("2006-01-02 15:04:05", minTmp, loc)
	maxUnix, _ := time.ParseInLocation("2006-01-02 15:04:05", maxTmp, loc)
	if current >= minUnix.Unix() && current <= maxUnix.Unix() {
		return true
	}
	return false
}
