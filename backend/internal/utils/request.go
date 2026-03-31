package utils

import (
	"encoding/json"
	"net/http"
	"path"
	"strconv"
	"strings"
)

func Parse[T any](data []byte) (T, error) {
	var v T
	err := json.Unmarshal(data, &v)
	return v, err
}

func WriteJSON(w http.ResponseWriter, v interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func ParseIDFromPath(p, prefix string) (int64, error) {
	idStr := strings.TrimPrefix(p, prefix)
	idStr = path.Clean("/" + idStr)
	idStr = strings.TrimPrefix(idStr, "/")
	return strconv.ParseInt(idStr, 10, 64)
}
