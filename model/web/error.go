package web

type WebError struct {
	Message string  `json:"message"`
	Errors  []Error `json:"errors"`
}

type Error struct {
	Resource string `json:"resource"`
	Field    string `json:"field"`
	Code     string `json:"code"`
}
