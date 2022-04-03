package actions

import (
	"context"
	"fmt"

	"github.com/camunda-cloud/zeebe/clients/go/pkg/zbc"
)

// DeployProcess deploy a zeeBe process bpmn
func DeployProcess(ctx context.Context, zbClient zbc.Client, resourcePath string) {
	// deploy workflow

	response, err := zbClient.NewDeployProcessCommand().AddResourceFile(resourcePath).Send(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.String())
}
