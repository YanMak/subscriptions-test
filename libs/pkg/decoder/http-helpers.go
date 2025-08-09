package decoder

import (
	"encoding/json"
	"net/http"
)

func DecodeBody[T any](r *http.Request, out *T) error {
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(out)
}
