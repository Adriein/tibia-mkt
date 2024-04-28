package types

import "net/http"

type ExoriVisTradeHttpHandler func(w http.ResponseWriter, r *http.Request) error
