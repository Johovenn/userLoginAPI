package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(r *http.Request, user interface{}) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(user)	

	PanicIfError(err)
}


func WriteResponse(writer http.ResponseWriter, response interface{}){
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(response)

	PanicIfError(err)
}