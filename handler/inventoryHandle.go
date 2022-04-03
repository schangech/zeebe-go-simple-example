package handler

import (
	"fmt"

	"github.com/camunda-cloud/zeebe/clients/go/pkg/entities"
	"github.com/camunda-cloud/zeebe/clients/go/pkg/worker"
)

func InventoryHandleJob(client worker.JobClient, job entities.Job) {

	ff := func() {
		fmt.Println("is a inventory handle test func")
	}

	wrapHandle(client, job, ff)
}
