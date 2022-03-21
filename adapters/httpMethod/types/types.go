package types

type JSONResponse struct {
	Result string `json:"result"`
}

type BulkJSONResponse struct {
	CorrelationID string `json:"correlation_id"`
	ShortURL      string `json:"short_url"`
}
