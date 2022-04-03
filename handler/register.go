package handler

import (
	"github.com/camunda-cloud/zeebe/clients/go/pkg/entities"
	"github.com/camunda-cloud/zeebe/clients/go/pkg/worker"
)

var (
	registerStep1JobName     = "step1"
	registerStep2JobName     = "step2"
	registerStep3JobName     = "step3"
	registerPaymentJobName   = "payment-service"
	registerInventoryJobName = "inventory-service"
	registerShipmentJobName  = "shipment-service"

	RegisterJobs = map[string]func(worker.JobClient, entities.Job){
		registerStep1JobName:     Step1HandleJob,
		registerStep2JobName:     Step2HandleJob,
		registerStep3JobName:     Step3HandleJob,
		registerPaymentJobName:   PaymentHandleJob,
		registerInventoryJobName: InventoryHandleJob,
		registerShipmentJobName:  ShipmentHandleJob,
	}
)
