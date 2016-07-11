package main

import (
	"io/ioutil"
	"os"
	"strings"
	"text/template"
  "bytes"

  "github.com/jessevdk/go-flags"
  "github.com/imdario/mergo"
  "github.com/op/go-logging"
)

type Options struct {
  Version bool `short:"v" long:"version" description:"Show version"`
  Debug bool `short:"d" long:"debug" description:"Log debug messages"`
  TemplatePath string `short:"t" long:"template" description:"Input template file (- for stdin)" default:"-"`
  OutputPath string `short:"o" long:"output" description:"Output file (- for stdout)" default:"-"`

  InputPaths []string `short:"i" long:"input" description:"Input variable file(s)"`

  EnvPrefix string `short:"e" long:"env_prefix" description:"Environment variables prefix" default:"CONF_"`
}

const version=`0.0.1`
var opts Options
var parser = flags.NewParser(&opts, flags.Default)
var log = logging.MustGetLogger("default")
var logFormatter = logging.MustStringFormatter(
    `%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.5s} %{id:03x}%{color:reset} %{message}`,
)

func main() {
  logBackend := logging.NewLogBackend(os.Stderr, "", 0)
  logging.SetFormatter(logFormatter)
  logBackendLev := logging.AddModuleLevel(logBackend)
  logBackendLev.SetLevel(logging.ERROR, "default")
  logging.SetBackend(logBackendLev)

  _, err := parser.Parse()

  if err != nil {
    os.Exit(1)
  }

  // enable debug logging if debug flag is on
  if opts.Debug {
    logging.SetLevel(logging.DEBUG, "default")
  }
  log.Debug("debug logging enabled")

  // print version and exit if version flag is on
  if opts.Version {
		println(version)
		return
	}
  log.Infof("version %s", version)

  // this map holds global merged config variables
  config := make(map[string]interface{})

  err = readAndMergeFiles(opts.InputPaths, &config)
  if err != nil {
    log.Errorf(err.Error())
    os.Exit(1)
  }

  // Get all env variables with given prefix
  envConfig, err := getAllEnvVariables(opts.EnvPrefix)

  err = mergo.MergeWithOverwrite(&config, envConfig)
  if err != nil {
    log.Errorf(err.Error())
    os.Exit(1)
  }
  log.Debugf("Merged data: %v", config)

  log.Infof("Loading template: '%s'", opts.TemplatePath)
	temp, err := loadTemplate(opts.TemplatePath)

	if err != nil {
    log.Errorf(err.Error())
    os.Exit(1)
	}

  // err on missing keys
  temp.Option("missingkey=error")

  log.Infof("Opening output file: '%s'", opts.OutputPath)
	outFile, err := createStream(opts.OutputPath)
	if err != nil {
    log.Errorf(err.Error())
    os.Exit(1)
	}
	defer closeStream(outFile)


  log.Infof("Processing template")
  if opts.Debug {
    var renderedTemplate bytes.Buffer
    temp.Execute(&renderedTemplate, config)
    log.Debugf("Rendered template: \n--- TEMPLATE BEGIN ---\n%v--- TEMPLATE END ---", renderedTemplate.String())
  }
  err = temp.Execute(outFile, config)
	if err != nil {
    log.Errorf(err.Error())
    os.Exit(1)
	}

  log.Infof("All done")
}

func loadTemplate(path string) (temp *template.Template, err error) {
  file, err := openStream(path)
  if err != nil {
    return
  }
  defer closeStream(file)

  templateData, err := ioutil.ReadAll(file)
  if err != nil {
    return
  }

  temp, err = template.New("default").Parse(string(templateData))
  log.Debugf("Loaded template file: \n--- TEMPLATE BEGIN ---\n%v--- TEMPLATE END ---", string(templateData))
  return
}

func getAllEnvVariables(prefix string) (ret map[string]interface{}, err error) {
  log.Infof("Loading env variables with prefix '%s'", opts.EnvPrefix)
  ret = make(map[string]interface{})

  for _, e := range os.Environ() {
    if strings.HasPrefix(e, prefix) {
      eName := strings.Split(strings.TrimPrefix(e, prefix), "=")[0]
      eValue := strings.SplitN(strings.TrimPrefix(e, prefix), "=", 2)[1]
      ret[eName] = eValue
    }
  }
  log.Debugf("Loaded env variables: %v", ret)

  return
}


func readAndMergeFiles(inputPaths []string, configMap *map[string]interface{}) (err error) {
  log.Info("Starting to merge files")

  // go through intput files and merge them into one
  for _, fileName := range inputPaths {
  	var data interface{} = nil

    log.Infof("Loading '%s'", fileName)
		data, err = loadData(fileName)

    if err != nil {
		  return
	  }
    log.Debugf("Loaded data: %v", data)

    // Merge to config map
    err = mergo.MergeWithOverwrite(configMap, data)
    if err != nil {
      return
    }

    log.Debugf("Merged data: %v", configMap)
	}

  return
}
