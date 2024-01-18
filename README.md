# exec Package

[![Go Reference](https://pkg.go.dev/badge/github.com/mrmarble/exec.svg)](https://pkg.go.dev/github.com/mrmarble/exec)
[![Go Report Card](https://goreportcard.com/badge/github.com/mrmarble/exec)](https://goreportcard.com/report/github.com/mrmarble/exec)

The `exec` package in Go is designed to run external commands. It is similar to the `os/exec` package, but provides a simpler interface and does not support streaming.

## Main Features

- The main difference between `exec` and `os/exec` is that the process does not wait for child processes to finish. This is useful for running long-running processes that are expected to run in the background.

## Example

```go
out, err := exec.Command("date").Output()
if err != nil {
  log.Fatal(err)
}
fmt.Printf("The date is %s\n", out)
```

This example runs the `date` command and logs any errors that occur.
