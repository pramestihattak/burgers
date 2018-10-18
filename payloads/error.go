package payloads

type ErrorRequest struct {
	Code     int         `json:"code"`
	Response interface{} `json:"response"`
}
