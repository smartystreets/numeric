package httpx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
)

func WriteResponse(response http.ResponseWriter, err error) {
	if err != nil {
		writeErrorMessage(response, err.Error(), http.StatusInternalServerError)
	} else {
		response.Header().Set(ContentTypeHeader, MIMEApplicationJSON)
	}
}
func WriteErrorMessage(response http.ResponseWriter, message string, statusCode int) {
	response.Header().Set(ContentTypeHeader, MIMETextPlain)
	http.Error(response, message, statusCode)
}

func WriteRequest(response http.ResponseWriter, request *http.Request, message string, status int) {
	dump, _ := httputil.DumpRequest(request, false)
	response.Header().Set(ContentTypeHeader, MIMETextPlain)
	http.Error(response, fmt.Sprintf("%d %s\n\nRaw Request:\n\n%s\n\n%s", status, message, string(dump)), status)
}
func WriteJSON(contents interface{}, response http.ResponseWriter) {
	response.Header().Set(ContentTypeHeader, MIMEApplicationJSON)
	json.NewEncoder(response).Encode(contents)
}
func WritePrettyJSON(contents interface{}, response http.ResponseWriter) {
	response.Header().Set(ContentTypeHeader, MIMEApplicationJSON)

	payload, _ := json.Marshal(envelope.Results)
	var buffer bytes.Buffer
	json.Indent(&buffer, payload, "", "\t")
	buffer.WriteTo(response)
}

const ContentTypeHeader = "Content-Type"
const MIMEApplicationJSON = "application/json; charset=utf-8"
const MIMETextPlain = "text/plain; charset=utf-8"