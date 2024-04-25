package req

type List struct {
	Page int `json:"page" form:"page" uri:"page" binding:"required"`
	Size int `json:"size" form:"size" uri:"size" binding:"required"`
}

type ListNoRequiredSize struct {
	Page int `json:"page" form:"page" uri:"page" binding:"required"`
	Size int `json:"size" form:"size" uri:"size"`
}

type ListNoRequired struct {
	Page int `json:"page" form:"page" uri:"page"`
	Size int `json:"size" form:"size" uri:"size"`
}

type ListDefault struct {
	PN int `json:"pn" form:"pn" uri:"pn"`
	PL int `json:"pl" form:"pl" uri:"pl"`
}
