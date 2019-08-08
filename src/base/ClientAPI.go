package base

import (
	"fmt"
)

const (
	API_COMMAND_CLIENT_START     = 1
	API_COMMAND_CLIENT_SUSPEND   = 2
	API_COMMAND_CLIENT_RESUME    = 5
	API_COMMAND_CLIENT_SETTING   = 3
	API_COMMAND_REFRESH_SETTINGS = 4
)

type ClientAPIResult struct {
	command    int
	resultText string
}

func (r ClientAPIResult) getResultText() string {
	return r.resultText
}

func (r ClientAPIResult) String() string {
	return fmt.Sprintf("{ClientAPIResult: command=%v, resultText=%v}", r.command, r.resultText)
}

type ClientAPI struct {
	client *HentaiAtHomeClient
}
