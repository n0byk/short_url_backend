package entities

//Структура будущей таблицы
type URLCatalog struct {
	ShortURL string `json:"short_url"`
	FullURL  string `json:"original_url"`
}
