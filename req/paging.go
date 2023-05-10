package req

type List struct {
	Page int `json:"page" form:"page" uri:"page"`
	Size int `json:"size" form:"size" uri:"size"`
}
