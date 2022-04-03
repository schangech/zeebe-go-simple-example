package handler

import (
	"context"
	"log"

	"github.com/camunda-cloud/zeebe/clients/go/pkg/entities"
	"github.com/camunda-cloud/zeebe/clients/go/pkg/worker"
	"github.com/schangech/zeebe-go-simple-example/handler/common"
)

func Step1HandleJob(client worker.JobClient, job entities.Job) {

	vars := job.GetVariables()
	log.Println("Step1 handler job get variables ", vars)

	jobKey := job.GetKey()

	headers, err := job.GetCustomHeadersAsMap()
	if err != nil {
		// failed to handle job as we require the custom job headers
		common.FailJob(client, job)
		return
	}

	variables, err := job.GetVariablesAsMap()
	if err != nil {
		// failed to handle job as we require the variables
		common.FailJob(client, job)
		return
	}

	variables["step1"] = "step1"
	request, err := client.NewCompleteJobCommand().JobKey(jobKey).VariablesFromMap(variables)
	if err != nil {
		// failed to set the updated variables
		common.FailJob(client, job)
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
	//close(readyClose)
	close(OnceReadyClose[registerStep1JobName])
}
