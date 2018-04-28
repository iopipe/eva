package templates

import (
	"net/http"
)

type RequestHandler func(request *http.Request) string
type ResponseHandler func(response []byte, w http.ResponseWriter)
