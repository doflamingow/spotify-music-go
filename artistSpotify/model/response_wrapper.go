package model

type ResponseWrapper struct {
	Success bool        `from:"Success" json:"success"`
	Message string      `from:"Message" json:"message"`
	Data    interface{} `from:"Date" json:"data,omitempty"`
}
