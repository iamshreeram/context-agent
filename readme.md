## Context-Agent
This application publishes events to NATS messaging system

## How to Use:
Start the NATS server by running the following commands:
```
docker pull nats
docker run -d --name nats-main -p 4222:4222 -p 6222:6222 -p 8222:8222 nats
```

Run the Go application to publish events to the NATS server:
``` 
go run cmd/main.go
```