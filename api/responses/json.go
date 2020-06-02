package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Responses struct{
	message		string
	data		interface{}
	code		int
}

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {

	w.WriteHeader(statusCode)
	elements := map[string]interface{}{
		"data" : data,
		"msg" : "success",
		"code" : statusCode,
	}

	err := json.NewEncoder(w).Encode(elements)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())	
	}

}

func ERROR(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		JSON(w, statusCode, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}
	JSON(w, http.StatusBadRequest, nil)
}