package templates

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/satori/go.uuid"
)

type RequestHandler func(request *http.Request) string
type ResponseHandler func(response []byte, w http.ResponseWriter)

func HandleCloudfrontEvent(request *http.Request) string {
	headersMap := make(map[string]interface{})

	for headerName, headerValues := range request.Header {
		headerNameLC := strings.ToLower(headerName)
		headersMap[headerNameLC] = make([]map[string]interface{}, 0)
		for headerValue := range headerValues {
			headersMap[headerNameLC] = append(headersMap[headerNameLC].([]map[string]interface{}),
				map[string]interface{}{
					"key":   headerName,
					"value": headerValues[headerValue],
				},
			)
		}
	}

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Fatal(err)
	}

	data := map[string]interface{}{
		"Records": []interface{}{
			map[string]interface{}{
				"cf": map[string]interface{}{
					"config": map[string]interface{}{
						"distributionId": "EDFDVBD6EXAMPLE",
					},
					"request": map[string]interface{}{
						"clientIp": "2001:0db8:85a3:0:0:8a2e:0370:7334",
						"method":   request.Method,
						"uri":      request.URL.Path,
						"body":     string(body),
						"headers":  headersMap,
					},
				},
			},
		},
	}
	json, err := json.MarshalIndent(data, " ", " ")
	if err != nil {
		log.Fatal(err)
	}

	return string(json)
}

func HandleApiGwResponse(response []byte, w http.ResponseWriter) {
	var object map[string]interface{}
	err := json.Unmarshal(response, &object)
	if err != nil {
		log.Fatalf("JSON unmarshalling error.\nError: %s\nraw bytes: %s", err, response)
	}
	w.WriteHeader(int(object["statusCode"].(float64)))
}

func HandleCloudfrontResponse(request []byte, w http.ResponseWriter) {
	w.WriteHeader(200)
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

	return string(json)
}
