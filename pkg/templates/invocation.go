package templates

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	db "github.com/iopipe/eva/data"
)

func HandleInvocationEvent(request *http.Request) string {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Fatal(err)
	}
	var object map[string]interface{}
	err = json.Unmarshal(body, &object)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Stderr.WriteString("\n")
	}
	db.PutInvocation(object)
	return "{}"
}

func HandleInvocationResponse(response []byte, w http.ResponseWriter) {
	/* NoOp */
}
