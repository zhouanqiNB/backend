

初始简易框架来自neo4j的官方例子：https://github.com/neo4j-examples/golang-neo4j-realworld-example

## Run

```
go run ./cmd/conduit/
```

在跑之前需要在.zshrc设置环境变量。to target your specific Neo4j instance.All settings are mandatory.

| Environment variable  | Description |
| --------------------- | ----------- |
| NEO4J_URI             | [Connection URI](https://neo4j.com/docs/driver-manual/current/client-applications/#driver-connection-uris) of the instance (e.g. `bolt://localhost`, `neo4j+s://example.org`) |
| NEO4J_USERNAME        | Username of the account to connect with (must have read & write permissions) |
| NEO4J_PASSWORD        | Password of the account to connect with (must have read & write permissions)|

然后开放的接口位于本地3000端口。