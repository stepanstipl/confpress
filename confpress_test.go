/*
test:
- yaml
- json
- multiple flags
- merging
- non-existing variable
*/

package main

import (
  "testing"
  "os/exec"
  "io/ioutil"

  "github.com/stretchr/testify/assert"
)


func helperCommand(t *testing.T, bin string, args ...string) *exec.Cmd {
  cmd := exec.Command(bin, args...)
  return cmd
}

func TestEcho(t *testing.T) {

  output, err := helperCommand(t, "./confpress", "-t", "resources/1-template.yaml", "-i", "resources/1-input-1.yml", "-i", "resources/1-input-2.yaml", "-i", "resources/1-input-3.json").Output()

  expected, err := ioutil.ReadFile("resources/1-output.yaml")

  assert.Equal(t, string(output), string(expected), "Given inputs, output should match expected one.")
  assert.Nil(t, err, "No errors should be raised")
}
