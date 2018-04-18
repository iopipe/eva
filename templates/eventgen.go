package templates

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func CreateCloudfrontEvent(request *http.Request) {
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
						"uri":      request.URL,
						"body":     request.Body,
						/*"headers": map[string]interface{}{
							"authorization": []interface{}{
								map[string]interface{}{
									"key":   "Authorization",
									"value": authorization,
								},
							},
							"host": []interface{}{
								map[string]interface{}{
									"key":   "Host",
									"value": ,
								},
							},
							"user-agent": []interface{}{
								map[string]interface{}{
									"key":   "User-Agent",
									"value": "curl/7.51.0",
								},
							},
						},*/
					},
				},
			},
		},
	}
	json, err := json.MarshalIndent(data, " ", " ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(json))
}
