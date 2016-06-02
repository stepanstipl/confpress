package main

import (
  "encoding/json"
  "strings"
  "io/ioutil"
  "errors"

  yaml "gopkg.in/yaml.v2"
)


func loadJSON(filePath string) (ret map[string]interface{}, err error) {
  file, err := openStream(filePath)
  if err != nil {
    return
  }
  defer closeStream(file)

  err = json.NewDecoder(file).Decode(&ret)
  return
}


func loadYAML(filePath string) (ret map[string]interface{}, err error) {
  file, err := openStream(filePath)
  if err != nil {
    return
  }
  defer closeStream(file)

  yamlData, err := ioutil.ReadAll(file)
  if err != nil {
    return
  }

  ret = make(map[string]interface{})
  err = yaml.Unmarshal([]byte(yamlData), &ret)
  return
}


func loadData(filePath string) (ret map[string]interface{}, err error) {
  if strings.HasSuffix(filePath, ".yaml") || strings.HasSuffix(filePath, ".yml") {
    ret, err = loadYAML(filePath)
  } else if strings.HasSuffix(filePath, ".json") {
    ret, err = loadJSON(filePath)
  } else {
    err = errors.New("Unrecognized input variable file extension, currently supported are .yaml, .yml, .json.")
  }

  return
}
