package list

type Default struct {
	PageNum  int `json:"page_num" form:"page_num" uri:"page_num"`
	PageSize int `json:"page_size" form:"page_size" uri:"page_size"`
}
