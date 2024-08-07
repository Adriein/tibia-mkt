package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Encode[T any](w http.ResponseWriter, status int, v T) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(v); err != nil {
		return fmt.Errorf("encode json: %w", err)
	}

	return nil
}

func Decode[T any](r io.Reader) (T, error) {
	var v T
	if err := json.NewDecoder(r).Decode(&v); err != nil {
		return v, fmt.Errorf("decode json: %w", err)
	}
	return v, nil
}

func JsonEncode[T any](v T) ([]byte, error) {
	encoded, err := json.Marshal(v)

	if err != nil {
		return nil, fmt.Errorf("encode json: %w", err)
	}

	return encoded, nil
}

func JsonDecode[T any](data []byte) (T, error) {
	var v T

	if err := json.Unmarshal(data, &v); err != nil {
		return v, fmt.Errorf("decode json: %w", err)
	}

	return v, nil
}
