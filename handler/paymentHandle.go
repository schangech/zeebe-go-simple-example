package handler

import (
	"fmt"

	"github.com/camunda-cloud/zeebe/clients/go/pkg/entities"
	"github.com/camunda-cloud/zeebe/clients/go/pkg/worker"
)

func PaymentHandleJob(client worker.JobClient, job entities.Job) {

	ff := func() {
		fmt.Println("is a payment handle func")
	}

	wrapHandle(client, job, ff)
}
