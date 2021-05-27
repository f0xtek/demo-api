package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type payload struct {
	Status  int
	Message string
}

func getJSONResponse() ([]byte, error) {
	responsePayload := payload{
		Status:  http.StatusOK,
		Message: "Hello",
	}
	return json.Marshal(responsePayload)
}

func mainHandler(w http.ResponseWriter, req *http.Request) {
	p, err := getJSONResponse()
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(p))
}

func main() {
	port := 8090
	http.HandleFunc("/", mainHandler)
	fmt.Printf("Listening on \":%d\"", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
