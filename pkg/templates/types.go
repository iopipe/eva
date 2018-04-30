package templates

import (
	"net/http"
)

type Event map[string]interface{}
type RequestHandler func(request *http.Request) Event
type ResponseHandler func(response []byte, w http.ResponseWriter)
