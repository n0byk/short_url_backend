package entities

// Data Base struct
type URLCatalog struct {
	ShortURL string `json:"short_url"`
	FullURL  string `json:"original_url"`
}

// URL data struct
type URLdb struct {
	FullURL string
	UserID  string
}

type StatInfo struct {
	URLs  int `json:"urls"`
	Users int `json:"users"`
}
