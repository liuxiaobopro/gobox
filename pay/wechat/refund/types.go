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
