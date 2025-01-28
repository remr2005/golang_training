package main

import (
	"main/api"
	"net/http"
)

func main() {
	http.HandleFunc("/get", api.Get_tasks)
	http.ListenAndServe(":8080", nil)
}
