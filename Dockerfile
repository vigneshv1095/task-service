FROM golang:1.21-alpine

WORKDIR /app

# 1. Install git for VCS-backed modules
RUN apk add --no-cache git

# 2. Install swag CLI (version-pinned)
RUN go install github.com/swaggo/swag/cmd/swag@v1.8.12

# 3. Copy go.mod & go.sum and download deps
COPY go.mod go.sum ./
RUN go mod download

# 4. Copy all source
COPY . .

# 5. Generate Swagger docs
RUN swag init -g cmd/server/main.go -o docs

# 6. Build the service
RUN go build -o task-service cmd/server/main.go

EXPOSE 8080

CMD ["./task-service"]
