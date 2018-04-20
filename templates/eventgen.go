package templates

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func CreateCloudfrontEvent(request *http.Request) string {
	headersMap := make(map[string]interface{})

	for headerName, headerValues := range request.Header {
		headerNameLC := strings.ToLower(headerName)
		headersMap[headerNameLC] = make([]map[string]interface{}, 1)
		for headerValue := range headerValues {
			headersMap[headerNameLC] = append(headersMap[headerNameLC].([]map[string]interface{}),
				map[string]interface{}{
					"key":   headerName,
					"value": headerValue,
				},
			)
		}
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
						"body":     request.Body,
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
