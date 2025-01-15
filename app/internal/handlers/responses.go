package handlers

type PingResponse struct {
	Ok bool `json:"ok,omitempty"`
}

type ErrResponse struct {
	Errors []Error `json:"errors,omitempty"`
}

type Error struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message,omitempty"`
}
