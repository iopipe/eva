docs/eva_daemon.md## eva daemon

Run HTTP daemon for listening to events.

### Synopsis

Run HTTP daemon for listening to events.

### Options

```
  -a, --addr string                 HTTP(s) address to listen on. (default ":8080")
  -c, --command string              Pipe event(s) into specified shell command
  -h, --help                        help for daemon
  -l, --lambda string               Process event(s) with specified AWS Lambda ARN
  -e, --log-event string            Log process event(s) into file, or - for stdout
  -E, --log-event-response string   Log response event(s) into file, or - for stdout
  -q, --quiet                       Do not print to stdout/stderr unless -e or -E is specified.
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.eva.yaml)
```

### SEE ALSO

* [eva](eva.md)	 - Event Ally, a cli for managing serverless events
* [eva daemon apigw](eva_daemon_apigw.md)	 - Generate apigw requests from HTTP listener
* [eva daemon cloudfront](eva_daemon_cloudfront.md)	 - Generate cloudfront requests from HTTP listener

###### Auto generated by spf13/cobra on 4-May-2018
