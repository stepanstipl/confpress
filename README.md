[![wercker status](https://app.wercker.com/status/cf9e7754fd7c23ab91908d5b5299ea09/s/master "wercker status")](https://app.wercker.com/project/bykey/cf9e7754fd7c23ab91908d5b5299ea09)


**Confpress** is simple configuration templating tool to ease your life!

Given set of input files (either json or yaml) it will merge these and also merge any env variables matching given prefix and process template file with go templating engine.

```
Usage:
  confpress [OPTIONS]

Application Options:
  -v, --version     Show version
  -d, --debug       Log debug messages
  -t, --template=   Input template file (- for stdin) (default: -)
  -o, --output=     Output file (- for stdout) (default: -)
  -i, --input=      Input variable file(s)
  -e, --env_prefix= Environment variables prefix (default: CONF_)

Help Options:
  -h, --help        Show this help message
```
