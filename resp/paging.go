package resp

type List struct {
	Count uint64      `json:"count"`
	List  interface{} `json:"list"`
}
