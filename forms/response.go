package forms

//Response is the universla data type for query responses
type Response struct {
	Key     string `json:"key"`
	Message string `json:"message"`
}
