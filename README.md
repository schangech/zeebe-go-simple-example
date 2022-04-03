# zeebe-go-simple-example
zeebe golang simple example demo

## usage
[![N|TOKEN](./images/execute.gif)]

step1：register camunda cloud (free 30 days) or launch a local zeebe service  
https://console.cloud.camunda.io/signup

[![N|TOKEN](./images/create-camunda-cred.png)](https://console.cloud.camunda.io)

step2: get environment variables

```shell
export ZEEBE_ADDRESS=xxx.bru-2.zeebe.camunda.io:443
export ZEEBE_CLIENT_ID=xxx
export ZEEBE_CLIENT_SECRET=xxx
export ZEEBE_AUTHORIZATION_SERVER_URL=https://login.cloud.camunda.io/oauth/token
```

step3: seeting environment variables

dev: setting variables to ide

prod: export variables

## how to develop owner worker

step1: download desktop modeler (create bpmn camunda cloud bpmn pipeline)

[![N|Desktop Modeler]()](https://camunda.com/download/modeler/)

step2: See the example in the Handle directory then register the jobWorker

step3: run main.go

## example
[![N|BPMN1](./images/order-process.png)](./bpmn/order-process.bpmn)
[![N|BPMN2](./images/step1-2-3-click-bpmn-example.png)](./bpmn/step1-2-3-click.bpmn)
