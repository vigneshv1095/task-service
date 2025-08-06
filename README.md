# Task Management Service

A simple Go microservice for managing tasks (CRUD), with pagination & status-based filtering and Swagger documentation.

## 📦 Project Structure
- **cmd/server**: entrypoint
- **internal/handler**: HTTP layer
- **internal/service**: business logic
- **internal/repository**: data access
- **internal/model**: domain model
- **pkg/pagination**: parse limit/offset
- **docs**: Swagger docs (generated)

## ⚙️ Design Decisions
- **Clean separation**: each layer has a single responsibility.
- **Stateless API**: scales horizontally behind a load balancer.
- **In-memory repo** by default; swap in PostgreSQL by implementing `TaskRepository`.
- **Pagination**: `?limit=&offset=`
- **Filtering**: `?status=Pending|InProgress|Completed`
- **Swagger UI**: served at `/swagger/index.html`

## 🚀 Getting Started

### Prerequisites
- Go 1.21+

### Run Locally Using Docker
```bash
git clone git@github.com:vigneshv1095/task-service.git
cd task-service
docker build -t task-service:latest .
docker run -p 8080:8080 task-service:latest
```

Then visit [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

## 📈 Scalability & Extensibility
- **Horizontal scaling**: multiple pods + load balancer.
- **Inter-service**: use gRPC/REST to talk to a User Service, or produce Kafka events for async flows.
