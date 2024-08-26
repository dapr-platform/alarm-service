package api

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

var (
	_ = time.Second // import time.Second for unknown usage in api

)

func readInt(r *http.Request, param string, v int64) (int64, error) {
	p := r.FormValue(param)
	if p == "" {
		return v, nil
	}

	return strconv.ParseInt(p, 10, 64)
}

func writeJSON(ctx context.Context, w http.ResponseWriter, v interface{}) {
	data, _ := json.Marshal(v)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache")
	w.Write(data)
}

func readJSON(r *http.Request, v interface{}) error {
	buf, err := ioutil.ReadAll(r.Body)

	if err != nil {
		return err
	}

	return json.Unmarshal(buf, v)
}
