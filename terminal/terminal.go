package terminal

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/eminmuhammadi/pulsatio/config"
)

func Stdin() string {
	fmt.Printf("%s-%s# ", config.Name, config.Version)

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	return strings.TrimSpace(input)
}

func Exec(command string, timeout int) (string, error) {
	var cmd *exec.Cmd

	// Create a new context and add a timeout to it
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(timeout)*time.Millisecond,
	)
	defer cancel() // The cancel should be deferred so resources are cleaned up

	if runtime.GOOS == "windows" {
		cmd = exec.CommandContext(ctx, "cmd", "/C", command)
	} else {
		cmd = exec.CommandContext(ctx, "sh", "-c", command)
	}

	// Results
	stdout, stderr := cmd.Output()

	// We want to check the context error to see if the timeout was executed.
	// The error returned by cmd.Output() will be OS specific based on what
	// happens when a process is killed.
	if ctx.Err() == context.DeadlineExceeded {
		return "", errors.New("command timed out")
	}

	result := stdout

	if stderr != nil {
		result = append(result, []byte(stderr.Error())...)
	}

	return string(result), nil
}
