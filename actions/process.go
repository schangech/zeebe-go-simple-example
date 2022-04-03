package actions

import (
	"context"
	"fmt"

	"github.com/camunda-cloud/zeebe/clients/go/pkg/zbc"
)

// CreateProcess launch a new zeebe process instance
func CreateProcess(ctx context.Context, zbClient zbc.Client, processName string) {
	// create a new workflow instance
	variables := make(map[string]interface{})
	variables["step"] = "start"

	request, err := zbClient.NewCreateInstanceCommand().BPMNProcessId(processName).LatestVersion().VariablesFromMap(variables)
	if err != nil {
		panic(err)
	}

	result, err := request.Send(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println(result.String())
}
