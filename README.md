# Eva your Event Ally

Experimental cli for API and client debugging.

This package is NOT production ready!

```
Usage:
  eva [command]

Available Commands:
  daemon      Run HTTP daemon for listening to events.
    apigw
    cloudfront
    invocations
      -a, --addr string       HTTP(s) address to listen on. (default ":8080")
          --config string     config file (default is $HOME/.eva.yaml)
      -e, --exec string       Pipe events into specified shell command.
      -q, --request string    Save request JSON into file.
      -s, --response string   Save response JSON into file.
  generate    generate events for serverless functions.
    cloudfront
    apigw
      -A, --auth string     Authorization header
      -d, --data string     Data for body, or '-' for stdin.
      -h, --help            help for apigw
      -H, --host string     HTTP(s) host for event data.
      -X, --method string   HTTP Method (default "GET")
      -p, --path string     HTTP(s) path or uri.
  geturl      Return a URL to send HTTP(S) request to.
  help        Help about any command
  inspect     Inspect an event history record.
  invocations Invocations list
  list        List generated events
  play        Play event specified by id
      -e, --exec string       Pipe events into specified shell command.
      -q, --request string    Save request JSON into file.
Flags:
      --config string   config file (default is $HOME/.eva.yaml)
  -h, --help            help for eva
  -t, --toggle          Help message for toggle

Use "eva [command] --help" for more information about a command.
```

Licensed under the Apache 2.0 license
