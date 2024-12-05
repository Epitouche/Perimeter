package schemas

type Response struct {
	Message string `json:"message"`
}

type ErrorRespose struct {
	Error string `json:"error"`
}
