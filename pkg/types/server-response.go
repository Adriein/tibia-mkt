package types

type ServerResponse struct {
	Ok    bool   `json:"ok"`
	Data  any    `json:"data"`
	Error string `json:"error,omitempty"`
}
