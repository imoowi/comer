/*
Copyright © 2023 jun<simpleyuan@gmail.com>
*/
package response

import "math"

type Pages struct {
	Count     int64 `json:"count"`     //总数
	CurPage   int64 `json:"page"`      //当前页码
	TotalPage int64 `json:"totalPage"` //总页数
	PageSize  int64 `json:"pageSize"`  //分页数
}

type PageList struct {
	Pages Pages `json:"pages"` //分页数据
	List  any   `json:"data"`  //返回数据
}

func MakePages(count int64, curPage int64, pageSize int64) (pages Pages) {
	pages.Count = count
	pages.PageSize = pageSize
	pages.CurPage = curPage
	totalPageFloat := float64(count) / float64(pageSize)
	pages.TotalPage = int64(math.Ceil(totalPageFloat))
	return
}
