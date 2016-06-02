package main

import (
  "os"
)

// open file or stdin if -
func openStream(path string) (file *os.File, err error) {
  if path == "-" {
    file = os.Stdin
  } else {
    file, err = os.Open(path)
  }
  return
}

// create file or stdout if -
func createStream(path string) (file *os.File, err error) {
  if path == "-" {
    file = os.Stdout
  } else {
    file, err = os.Create(path)
  }
  return
}

// close file if real file
func closeStream(file *os.File) (err error) {
  if file == os.Stdout || file == os.Stdin {
    return
  }
  return file.Close()
}
