package mapbox

type GeocodingRes struct {
	Type        string    `json:"type,omitempty"`
	Query       []string  `json:"query,omitempty"`
	Features    []Feature `json:"features,omitempty"`
	Attribution string    `json:"attribution,omitempty"`
}

type Feature struct {
	ID                string     `json:"id,omitempty"`
	Type              string     `json:"type,omitempty"`
	PlaceType         []string   `json:"place_type,omitempty"`
	Relevance         float64    `json:"relevance,omitempty"`
	Properties        Properties `json:"properties,omitempty"`
	TextZh            string     `json:"text_zh,omitempty"`
	LanguageZh        string     `json:"language_zh,omitempty"`
	PlaceNameZh       string     `json:"place_name_zh,omitempty"`
	Text              string     `json:"text,omitempty"`
	Language          string     `json:"language,omitempty"`
	PlaceName         string     `json:"place_name,omitempty"`
	MatchingText      string     `json:"matching_text,omitempty"`
	MatchingPlaceName string     `json:"matching_place_name,omitempty"`
	Center            [2]float64 `json:"center,omitempty"`
	Geometry          Geometry   `json:"geometry,omitempty"`
	Context           []Context  `json:"context,omitempty"`
}

type Properties struct {
	Foursquare string `json:"foursquare,omitempty"`
	Wikidata   string `json:"wikidata,omitempty"`
	Landmark   bool   `json:"landmark,omitempty"`
	Address    string `json:"address,omitempty"`
	Category   string `json:"category,omitempty"`
}

type Geometry struct {
	Coordinates [2]float64 `json:"coordinates,omitempty"`
	Type        string     `json:"type,omitempty"`
}

type Context struct {
	ID         string `json:"id,omitempty"`
	MapboxID   string `json:"mapbox_id,omitempty"`
	TextZh     string `json:"text_zh,omitempty"`
	Text       string `json:"text,omitempty"`
	Wikidata   string `json:"wikidata,omitempty"`
	ShortCode  string `json:"short_code,omitempty"`
	LanguageZh string `json:"language_zh,omitempty"`
	Language   string `json:"language,omitempty"`
}
