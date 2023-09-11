package gaode

type GeocodeRes struct {
	Status   string        `json:"status,omitempty"`
	Info     string        `json:"info,omitempty"`
	InfoCode string        `json:"infocode,omitempty"`
	Count    string        `json:"count,omitempty"`
	Geocodes []GeocodeItem `json:"geocodes,omitempty"`
}

type GeocodeItem struct {
	FormattedAddress string      `json:"formatted_address,omitempty"`
	Country          string      `json:"country,omitempty"`
	Province         string      `json:"province,omitempty"`
	CityCode         string      `json:"citycode,omitempty"`
	City             interface{} `json:"city,omitempty"`
	District         interface{} `json:"district,omitempty"`
	Township         []string    `json:"township,omitempty"`
	Neighborhood     struct {
		Name []string `json:"name,omitempty"`
		Type []string `json:"type,omitempty"`
	} `json:"neighborhood,omitempty"`
	Building struct {
		Name []string `json:"name,omitempty"`
		Type []string `json:"type,omitempty"`
	} `json:"building,omitempty"`
	Adcode   string   `json:"adcode,omitempty"`
	Street   []string `json:"street,omitempty"`
	Number   []string `json:"number,omitempty"`
	Location string   `json:"location,omitempty"`
	Level    string   `json:"level,omitempty"`
}

type Tip struct {
	ID       interface{} `json:"id,omitempty"`
	Name     string      `json:"name,omitempty"`
	District string      `json:"district,omitempty"`
	Adcode   string      `json:"adcode,omitempty"`
	Location interface{} `json:"location,omitempty"`
	Address  interface{} `json:"address,omitempty"`
	Typecode string      `json:"typecode,omitempty"`
	City     interface{} `json:"city,omitempty"`
}

type TipsRes struct {
	Tips     []Tip  `json:"tips,omitempty"`
	Status   string `json:"status,omitempty"`
	Info     string `json:"info,omitempty"`
	Infocode string `json:"infocode,omitempty"`
	Count    string `json:"count,omitempty"`
}
