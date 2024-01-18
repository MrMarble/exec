# exec Package

The `exec` package in Go is designed to run external commands. It is similar to the `os/exec` package, but provides a simpler interface and does not support streaming.

## Main Features

- The main difference between `exec` and `os/exec` is that the process does not wait for child processes to finish. This is useful for running long-running processes that are expected to run in the background.

## Example

```go
cmd := &Cmd{
    Path: "/usr/bin/ls",
    Args: []string{"ls", "-l"},
}

err := cmd.Run()
if err != nil {
    log.Fatal(err)
}
```

This example runs the `ls -l` command and logs any errors that occur.