// Package exec runs external commands. It is similar to os/exec, but it
// provides a simpler interface and does not support streaming.

// The main difference is that the process does not wait for child processes to
// finish. This is useful for running long-running processes that are expected
// to run in the background.
package exec

import (
	"bytes"
	"errors"
	"io"
	"os"
	"sync"
)

// Cmd represents an external command being prepared or run.
//
// A Cmd cannot be reused after calling its Output or CombinedOutput
// methods.
type Cmd struct {
	// Path is the path of the command to run.
	//
	// This is the only field that must be set to a non-zero
	// value. If Path is relative, it is evaluated relative
	// to Dir.
	Path string

	// Args holds command line arguments, including the command as Args[0].
	// If the Args field is empty or nil, Run uses {Path}.
	//
	// In typical use, both Path and Args are set by calling Command.
	Args []string

	// Process is the underlying process, once started.
	Process *os.Process

	// ProcessState contains information about an exited process.
	// If the process was started successfully, Wait or Run will
	// populate its ProcessState when the command completes.
	ProcessState *os.ProcessState

	// mu locks the Cmd during Output, preventing race conditions
	mu sync.Mutex
}

// Command returns the Cmd struct to execute the named program with
// the given arguments.
//
// It sets only the Path and Args in the returned structure.
func Command(name string, arg ...string) *Cmd {
	return &Cmd{
		Path: name,
		Args: append([]string{name}, arg...),
		mu:   sync.Mutex{},
	}
}

// Output runs the command and returns its standard output.
func (c *Cmd) Output() ([]byte, error) {
	// Create a pipe for the child's stdout.
	w := bytes.NewBuffer([]byte{})
	pr, pw, err := os.Pipe()
	if err != nil {
		return nil, err
	}
	defer pw.Close()

	// Start copying from the pipe to the buffer.
	go func() {
		c.mu.Lock()
		io.Copy(w, pr)
		c.mu.Unlock()
		pr.Close() // in case io.Copy stopped due to write error
	}()

	proc, err := os.StartProcess(c.Path, c.Args, &os.ProcAttr{Files: []*os.File{os.Stdin, pw, os.Stderr}})
	if err != nil {
		return nil, err
	}

	c.Process = proc
	state, err := proc.Wait()
	if err != nil {
		return nil, err
	}
	c.ProcessState = state
	// Stop the copy and wait for it to finish.
	pr.Close()
	c.mu.Lock()
	defer c.mu.Unlock()
	if !state.Success() {
		return w.Bytes(), errors.New(state.String())
	}
	return w.Bytes(), nil
}

// CombinedOutput runs the command and returns its combined standard
// output and standard error.
func (c *Cmd) CombinedOutput() ([]byte, error) {
	// Create a pipe for the child's stdout.
	w := bytes.NewBuffer([]byte{})
	pr, pw, err := os.Pipe()
	if err != nil {
		return nil, err
	}
	defer pw.Close()

	// Start copying from the pipe to the buffer.
	go func() {
		c.mu.Lock()
		io.Copy(w, pr)
		c.mu.Unlock()
		pr.Close() // in case io.Copy stopped due to write error
	}()

	proc, err := os.StartProcess(c.Path, c.Args, &os.ProcAttr{Files: []*os.File{os.Stdin, pw, pw}})
	if err != nil {
		return nil, err
	}

	c.Process = proc
	state, err := proc.Wait()
	if err != nil {
		return nil, err
	}
	c.ProcessState = state
	// Stop the copy and wait for it to finish.
	pr.Close()
	c.mu.Lock()
	defer c.mu.Unlock()
	if !state.Success() {
		return w.Bytes(), errors.New(state.String())
	}
	return w.Bytes(), nil
}
