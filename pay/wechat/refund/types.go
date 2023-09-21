package wechat

type Resource struct {
	OriginalType   string `json:"original_type,omitempty"`
	Algorithm      string `json:"algorithm,omitempty"`
	Ciphertext     string `json:"ciphertext,omitempty"`
	AssociatedData string `json:"associated_data,omitempty"`
	Nonce          string `json:"nonce,omitempty"`
}

type Notify struct {
	ID           string   `json:"id,omitempty"`
	CreateTime   string   `json:"create_time,omitempty"`
	ResourceType string   `json:"resource_type,omitempty"`
	EventType    string   `json:"event_type,omitempty"`
	Summary      string   `json:"summary,omitempty"`
	Resource     Resource `json:"resource,omitempty"`
}

type Amount struct {
	Total       int `json:"total,omitempty"`
	Refund      int `json:"refund,omitempty"`
	PayerTotal  int `json:"payer_total,omitempty"`
	PayerRefund int `json:"payer_refund,omitempty"`
}

type DecryptCiphertext struct {
	MchID               string `json:"mchid,omitempty"`
	TransactionID       string `json:"transaction_id,omitempty"`
	OutTradeNo          string `json:"out_trade_no,omitempty"`
	RefundID            string `json:"refund_id,omitempty"`
	OutRefundNo         string `json:"out_refund_no,omitempty"`
	RefundStatus        string `json:"refund_status,omitempty"`
	SuccessTime         string `json:"success_time,omitempty"`
	UserReceivedAccount string `json:"user_received_account,omitempty"`
	Amount              Amount `json:"amount,omitempty"`
}
