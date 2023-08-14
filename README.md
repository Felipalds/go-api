# Go HTTP APIs and deploy with Docker
- To production, Go needs only the > binary <
- CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o main ./main.go