package log

type LogDataRequest struct {
	Message string
	Nested  Nested
}

type Nested struct {
	Values any
}

type LogDataResponse struct {
	Message string
}
