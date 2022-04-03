package common

import (
	"context"
	"log"
	"reflect"
	"runtime"

	"github.com/camunda-cloud/zeebe/clients/go/pkg/entities"
	"github.com/camunda-cloud/zeebe/clients/go/pkg/worker"
)

func WrapHandle(client worker.JobClient, job entities.Job, f func()) {
	jobKey := job.GetKey()

	headers, err := job.GetCustomHeadersAsMap()
	if err != nil {
		// failed to handle job as we require the custom job headers
		FailJob(client, job)
		return
	}

	variables, err := job.GetVariablesAsMap()
	if err != nil {
		// failed to handle job as we require the variables
		FailJob(client, job)
		return
	}

	res := func() string {
		return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	}

	f()

	variables[res()] = "done"
	request, err := client.NewCompleteJobCommand().JobKey(jobKey).VariablesFromMap(variables)
	if err != nil {
		// failed to set the updated variables
		FailJob(client, job)
		return
	}

	log.Println("Complete job", jobKey, "of type", job.Type)
	log.Println("Processing order:", variables["orderId"])
	log.Println("Collect money using payment method:", headers["method"])

	ctx := context.Background()
	_, err = request.Send(ctx)
	if err != nil {
		panic(err)
	}

	log.Println("Successfully completed job")
}
