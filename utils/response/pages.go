package response

import "math"

type Pages struct {
	Count     int64 `json:"count"`
	CurPage   int64 `json:"page"`
	TotalPage int64 `json:"totalPage"`
	PageSize  int64 `json:"pageSize"`
}

func MakePages(count int64, curPage int64, pageSize int64) (pages Pages) {
	pages.Count = count
	pages.PageSize = pageSize
	pages.CurPage = curPage
	var totalPageFloat float64
	totalPageFloat = float64(count) / float64(pageSize)
	pages.TotalPage = int64(math.Ceil(totalPageFloat))
	return
}
