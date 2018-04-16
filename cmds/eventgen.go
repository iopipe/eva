package cmds

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)


var cmdFlagMakeEventHost string
var cmdFlagMakeEventUri string
var cmdFlagMakeEventAuthorization string

var cmdMakeEvent = &cobra.Command{
	Use:   "generate",
	Short: "Generate an event.",
	Run: func(cmd *cobra.Command, args []string) {
		createCloudfrontEvent(cmdFlagMakeEventHost, cmdFlagMakeEventUri, cmdFlagMakeEventAuthorization)
	},
}
func initMakeEvent() {
  cmdMakeEvent.Flags().StringVarP(&cmdFlagMakeEventHost, "host", "H", "", "HTTP(s) host for event data.")
  cmdMakeEvent.Flags().StringVarP(&cmdFlagMakeEventUri, "path", "p", "", "HTTP(s) path or uri.")
  cmdMakeEvent.Flags().StringVarP(&cmdFlagMakeEventAuthorization, "auth", "u", "", "Authorization header")
}

func createCloudfrontEvent(host, uri, authorization string) {
	data := map[string]interface{}{
		"Records": []interface{}{
			map[string]interface{}{
				"cf": map[string]interface{}{
          "config": map[string]interface{}{
            "distributionId": "EDFDVBD6EXAMPLE",
          },
					"request": map[string]interface{}{
						"clientIp": "2001:0db8:85a3:0:0:8a2e:0370:7334",
						"method":   "GET",
						"uri":      uri,
						"headers": map[string]interface{}{
              "authorization": []interface{}{
                map[string]interface{}{
                  "key": "Authorization",
                  "value": authorization,
                },
              },
							"host": []interface{}{
								map[string]interface{}{
									"key":   "Host",
									"value": host,
								},
							},
							"user-agent": []interface{}{
								map[string]interface{}{
									"key":   "User-Agent",
									"value": "curl/7.51.0",
								},
							},
						},
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
