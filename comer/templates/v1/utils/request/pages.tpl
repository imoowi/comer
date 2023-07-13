package request

type PageList struct {
	Page      int64  `json:"page" form:"page"`           //页码,默认为1
	PageSize  int64  `json:"pageSize" form:"pageSize" `  //页数,默认为20，最小为1，最大不超过1000
	SearchKey string `json:"searchKey" form:"searchKey"` //搜索关键字
}
