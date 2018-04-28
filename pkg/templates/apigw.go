package templates

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	db "github.com/iopipe/eva/data"
	"github.com/satori/go.uuid"
)

func HandleApiGwResponse(response []byte, w http.ResponseWriter) {
	var object map[string]interface{}
	err := json.Unmarshal(response, &object)
	if err != nil {
		log.Fatalf("JSON unmarshalling error.\nError: %s\nraw bytes: %s", err, response)
	}
	w.WriteHeader(int(object["statusCode"].(float64)))
}

func HandleApiGwEvent(request *http.Request) string {
	headersMap := make(map[string]interface{})

	/* API Gateway doesn't handle duplicate headers...
	   TODO: lookup behavior of API Gateway (which overrides, first or last?) */
	for headerName, headerValues := range request.Header {
		for headerIndex := range headerValues {
			headersMap[headerName] = headerValues[headerIndex]
		}
	}

	requestId, err := uuid.NewV4()
	if err != nil {
		log.Fatal("UUID generation error.")
	}

	contextMap := map[string]interface{}{
		"accountId":  "123456789012",
		"resourceId": "us4z18",
		"stage":      "eva-generated",
		"requestId":  requestId.String(),
		"identity": map[string]interface{}{
			"cognitoIdentityPoolId":         "",
			"accountId":                     "",
			"cognitoIdentityId":             "",
			"caller":                        "",
			"apiKey":                        "",
			"sourceIp":                      "127.0.0.1",
			"cognitoAuthenticationType":     "",
			"cognitoAuthenticationProvider": "",
			"userArn":                       "",
			"userAgent":                     request.Header.Get("User-Agent"),
			"user":                          "",
		},
		"resourcePath": "/",
		"httpMethod":   request.Method,
		"apiId":        "eva-fake-api",
	}

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Fatal(err)
	}

	data := map[string]interface{}{
		"path":           request.URL.Path,
		"httpMethod":     request.Method,
		"body":           string(body),
		"headers":        headersMap,
		"requestContext": contextMap,
		"resource":       "UNKNOWN",
	}
	json, err := json.MarshalIndent(data, " ", " ")
	if err != nil {
		log.Fatal(err)
	}

	db.PutEvent(data)
	return string(json)
}
