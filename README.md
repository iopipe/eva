# Eva your Event Ally

***EXPERIMENTAL***

Eva is a CLI application that enables developers
to work with events to store, replay, deliver,
and proxy. It is designed to work with event-driven
serverless systems.

Eva can generate events:
  `eva generate <event-type>`,

Consume and dispatch events as a daemon:
  `eva daemon <event-type>`,

Replay events and redispatch:
  `eva play <event-id>`,

Store invocation data for serverless functions:
  `eva invocations`,

and is your serverless event ally brought to you with <3

# Documentation:

Read complete documentation at https://iopipe.github.io/eva/

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
    apigw
    cloudfront
    requestbin
      -A, --auth string     Authorization header
      -d, --data string     Data for body, or '-' for stdin.
      -h, --help            help for apigw
      -H, --host string     HTTP(s) host for event data.
      -X, --method string   HTTP Method (default "GET")
      -p, --path string     HTTP(s) path or uri.
  help        Help about any command
  inspect     Inspect an event history record.
  invocations Invocations list
  list        List generated events
  play        Play event specified by id
      -e, --exec string       Pipe events into specified shell command.
      -q, --request string    Save request JSON into file.
  requestbin  Return a public URL to send HTTP(S) request to.
```

Licensed under the Apache 2.0 license
