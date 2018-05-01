package templates

import (
	"io/ioutil"
	"log"

	"net/http"
	"strings"
)

func HandleCloudfrontResponse(request []byte, w http.ResponseWriter) {
	w.WriteHeader(200)
}

func HandleCloudfrontEvent(request *http.Request) Event {
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
	return data
}

func HandleCloudfrontRequestEvent(request *http.Request) Event {
	event := HandleCloudfrontEvent(request)
	return event
}

func HandleCloudfrontResponseEvent(request *http.Request) Event {
	event := HandleCloudfrontRequestEvent(request)
	event["response"] = map[string]interface{}{
		"status":            "200",
		"statusDescription": "OK",
		"headers": map[string]interface{}{
			"server": []map[string]interface{}{
				{
					"key":   "Server",
					"value": "MyCustomOrigin",
				},
			},
			"set-cookie": []map[string]interface{}{
				{
					"key":   "Set-Cookie",
					"value": "theme=light",
				},
				{
					"key":   "Set-Cookie",
					"value": "sessionToken=abc123; Expires=Wed, 09 Jun 2021 10:18:14 GMT",
				},
			},
		},
	}
	return event
}
