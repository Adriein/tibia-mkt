package types

import "net/http"

type TibiaMktHttpHandler func(w http.ResponseWriter, r *http.Request) error
