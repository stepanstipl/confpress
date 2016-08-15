[![wercker status](https://app.wercker.com/status/cf9e7754fd7c23ab91908d5b5299ea09/s/master "wercker status")](https://app.wercker.com/project/bykey/cf9e7754fd7c23ab91908d5b5299ea09)


**Confpress** is simple configuration templating tool to ease your life!

Given set of input files (either json or yaml) it will merge these and also merge any env variables matching given prefix and process template file with go templating engine.

```
Usage:
  confpress [OPTIONS]

Application Options:
  -d, --debug       Log debug messages
  -e, --env_prefix= Environment variables prefix (default: CONF_)
  -i, --input=      Input variable file(s)
  -m, --missing     Allow missing keys
  -o, --output=     Output file (- for stdout) (default: -)
  -t, --template=   Input template file (- for stdin) (default: -)
  -v, --version     Show version

Help Options:
  -h, --help        Show this help message
```

### Examples
- **To allow empty variable**
  Use `-m` switch to allow empty values (else confpress throws an error). Golang templates will render it as `<no value>`. To keep the output really empty use construct `{{ or .my_value "" }}` in your template.

- **Reading a variable from yaml/json as shell env variable**

  `MY_VAR=$(echo "{{ .my_variable }}" | confpress -i variables.yaml)`
