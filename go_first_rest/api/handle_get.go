package api

import (
	"encoding/json"
	"main/domain"
	"net/http"
)

func Get_tasks(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		data, _ := json.Marshal(domain.Get_tasks())
		w.Write(data)
	}
}
