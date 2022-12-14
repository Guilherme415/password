package models

// Model para responses HTTP
type Response struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
