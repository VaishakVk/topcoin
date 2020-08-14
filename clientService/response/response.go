package response

import (
	"encoding/json"
	"net/http"
)

// APIResponse struct
type APIResponse struct {
	Status   bool        `json:"status"`
	Response interface{} `json:"response"`
	Error    string      `json:"error"`
}

// SendResponse Endpoint
func SendResponse(response http.ResponseWriter, statusCode int, Status bool, Response interface{}, Error string) {
	response.WriteHeader(statusCode)
	json.NewEncoder(response).Encode(APIResponse{Status: Status, Response: Response, Error: Error})
}
