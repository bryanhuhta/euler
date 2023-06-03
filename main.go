package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/bryanhuhta/euler/internal/problem"
)

func main() {
	id, params := mustParseArgs(os.Args[1:])

	err := problem.Execute(id, params)
	if err != nil {
		exitWithMessage("error: %v", err)
	}
}

func mustParseArgs(args []string) (int, problem.Params) {
	if len(args) < 1 {
		exitWithMessage("error: need at least 1 argument\n  usage: euler id [key:value...]")
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		exitWithMessage("error: id is not an integer: %v", err)
	}

	params := make(problem.Params, len(args[1:]))
	for i, arg := range args[1:] {
		tokens := strings.SplitN(arg, ":", 2)
		if len(tokens) != 2 {
			exitWithMessage("error: malformed parameter %d: %s", i, arg)
		}

		key, val := tokens[0], tokens[1]
		params[key] = val
	}

	return id, params
}

func exitWithMessage(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}
