package types

import "encoding/json"

type ServerResponse struct {
	Ok    bool            `json:"ok"`
	Data  json.RawMessage `json:"data"`
	Error string          `json:"error,omitempty"`
}
