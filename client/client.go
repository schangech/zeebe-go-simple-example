package client

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/camunda-cloud/zeebe/clients/go/pkg/zbc"
)

var (
	ID                     = os.Getenv(zbc.OAuthClientIdEnvVar)
	Secret                 = os.Getenv(zbc.OAuthClientSecretEnvVar)
	Audience               = os.Getenv(zbc.GatewayAddressEnvVar)
	ZeeBeGateway           = os.Getenv(zbc.GatewayAddressEnvVar)
	AuthorizationServerURL = os.Getenv(zbc.OAuthAuthorizationUrlEnvVar)
)

// NewZeeBeAuthClient 自己实现基于认证的客户端
func NewZeeBeAuthClient() zbc.Client {

	var audience string
	index := strings.LastIndex(Audience, ":")
	if index > 0 {
		audience = Audience[0:index]
	}

	credProvider, err := zbc.NewOAuthCredentialsProvider(&zbc.OAuthProviderConfig{
		ClientID:               ID,
		ClientSecret:           Secret,
		AuthorizationServerURL: AuthorizationServerURL,
		Audience:               audience,
	})
	if err != nil {
		panic(err)
	}

	client, err := zbc.NewClient(&zbc.ClientConfig{
		GatewayAddress:      ZeeBeGateway,
		CredentialsProvider: credProvider,
	})
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	response, err := client.NewTopologyCommand().Send(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.String())
	return client
}

// NewZeeBeGWClient source env.sh
// NewZeeBeGWClient 采用环境变量方式
func NewZeeBeGWClient() zbc.Client {
	gatewayAddr := ZeeBeGateway
	plainText := false

	if gatewayAddr == "" {
		gatewayAddr = fmt.Sprintf("%s:%s", zbc.DefaultAddressHost, zbc.DefaultAddressPort)
		plainText = true
	}

	zbClient, err := zbc.NewClient(&zbc.ClientConfig{
		GatewayAddress:         gatewayAddr,
		UsePlaintextConnection: plainText,
	})

	if err != nil {
		panic(err)
	}

	return zbClient
}
