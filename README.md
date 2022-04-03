# zeebe-go-simple-example
zeebe golang simple example demo

第一步：注册登录
https://console.cloud.camunda.io/signup

第二步：获取对应的环境变量信息
```shell
export ZEEBE_ADDRESS=xxx.bru-2.zeebe.camunda.io:443
export ZEEBE_CLIENT_ID=xxx
export ZEEBE_CLIENT_SECRET=xxx
export ZEEBE_AUTHORIZATION_SERVER_URL=https://login.cloud.camunda.io/oauth/token
```

第三步：
dev: 将对应的环境变量设置在对应的IDE中，然后运行main.go
prod: source env.sh, 然后运行二进制

