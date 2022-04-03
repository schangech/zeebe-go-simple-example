package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/camunda-cloud/zeebe/clients/go/pkg/entities"
	"github.com/camunda-cloud/zeebe/clients/go/pkg/worker"
	"github.com/camunda-cloud/zeebe/clients/go/pkg/zbc"
	"github.com/schangech/zeebe-go-simple-example/actions"
	"github.com/schangech/zeebe-go-simple-example/client"
	"github.com/schangech/zeebe-go-simple-example/handler"
)

func main() {

	ctx := context.Background()

	zbClient := client.NewZeeBeGWClient()

	// 如果采用本地提交上传BPMN流程文件
	actions.DeployProcess(ctx, zbClient, "bpmn/order-process.bpmn")

	//采用程序启动Process实例
	actions.CreateProcess(ctx, zbClient, "order-process")

	// 将我们自己实现的job注册到zeeBe中，用于watch zeeBe中job
	// 一旦有新的job，如果是采用watch函数，则会立即执行，并持续监听
	registerJobs := handler.RegisterJobs

	for jobName, f := range registerJobs {
		go wrapWatchingService(jobName, f, zbClient)
		//go wrapOnceService(jobName, f, zbClient)
	}

	exist()
}

// TODO: 下面函数使用场景：某些进程只使用一次就退出
// 目前还存在问题，因为jobWorker Open 是协程，如果立即close，会造成handler没有执行完
// 但如果在handler中嵌入channel，则需要handler支持参数才能支持
// wrapOnceService 采用这个函数封装，表示对应的handle只会执行一次，即退出
func wrapOnceService(jobTypeName string, handleFunc func(worker.JobClient, entities.Job), zbClient zbc.Client) {

	handler.OnceReadyClose[jobTypeName] = make(chan struct{})

	fmt.Printf("start watch jobName %s \n", jobTypeName)
	jobWorker := zbClient.NewJobWorker().JobType(jobTypeName).Handler(handleFunc).Open()

	closeOnce(jobTypeName)
	jobWorker.Close()
	jobWorker.AwaitClose()
}

func closeOnce(key string) {
	<-handler.OnceReadyClose[key]
	//<-readyClose
}

// wrapWatchingService 采用这个函数封装，表示对应的handle会一直执行，除非认为ctrl+c退出
func wrapWatchingService(jobTypeName string, handleFunc func(worker.JobClient, entities.Job), zbClient zbc.Client) {

	fmt.Printf("start watch jobName %s \n", jobTypeName)
	zbClient.NewJobWorker().JobType(jobTypeName).Handler(handleFunc).Open()

}

func exist() {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	sig := <-sc
	fmt.Printf("exit: signal=<%d>.", sig)

}
