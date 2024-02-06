package log

type LogDataRequest struct {
	Message string `json:"message"`
	Nested  Nested `json:"nested"`
}

type Nested struct {
	Values any `json:"values"`
}

type LogDataResponse struct {
	Message string
}
