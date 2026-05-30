# 🧩 Selena Hotels Service

<!--![Docker](https://img.shields.io/badge/Docker-Containers-blue)-->
![CI/CD](https://img.shields.io/badge/CI/CD-GitHub_actions-c01c1c)

---

## 📌 Overview

hotels-service is a backend microservice responsible for managing hotel data within the Selena platform.

It is designed to run in a cloud-native **AWS** environment, fully integrated with:

- Auto Scaling infrastructure
- Internal service networking
- Secure secrets management
- Container-based deployment via Docker

The service exposes REST API endpoints for:

- hotel creation
- hotel retrieval
- hotel management

---

## 🚀 Key Characteristics

- Stateless service (horizontal scaling ready)
- Runs inside Docker containers
- Deployed on EC2 via Auto Scaling Group
- Integrated with AWS Application Load Balancer
- Uses secure database connection (private network)
- Configuration managed via AWS Secrets Manager

---

## ☁️ Cloud Integration (AWS)

<!--This service is tightly coupled with the Selena AWS infrastructure.-->

### 🔗 Networking

- Runs inside private subnets
- No direct public access
- Receives traffic only from:
    -- Public ALB (external requests)
    -- Internal ALB (service-to-service)

<br>

### ⚖️ Load Balancing

#### Public endpoint
- https://hotels-service.selena-aws.com

#### Internal endpoint
- http://hotels.internal.selena

<br>

### 💻 Compute

- EC2 instances managed by Auto Scaling Group

#### Scaling configuration:
- min: 1
- max: 3

#### Each instance runs:
- Docker container with hotels-service

<br>

### 📦 Containerization

- Built via Dockerfile
- Stored in Amazon ECR
- Pulled dynamically on EC2 startup

<br>

### 🔐 Secrets Management

Secrets are NOT stored in the repository.

#### They are injected at runtime via:

- AWS Secrets Manager

Includes:
- database credentials
- service configuration

<br>

### 🗄️ Database

- Uses CockroachDB (self-managed on EC2)

#### Connection characteristics:

- Private network only
- TLS certificates (stored in db/certs/)
- Access controlled via IAM + SSM

<br>

### ⚙️ Service Bootstrap (EC2 UserData)

On instance startup:
- Docker is initialized
- ECR authentication is performed
- Secrets are fetched
- Container is pulled
- Service starts automatically

---

## 🧱 Project Structure

    hotels-service/
    ├── cmd/                    # Application entrypoints
    ├── internal/
    │   ├── handlers/          # HTTP handlers
    │   ├── services/          # Business logic
    │   ├── repository/        # DB interaction
    │   ├── models/            # Domain models
    │   ├── database/          # DB setup & seeds
    │   └── router/            # Routing
    │
    ├── db/
    │   ├── migrations/        # DB migrations
    │   ├── certs/             # TLS certificates
    │   └── migrate.sh
    │
    ├── _docker/               # Docker entrypoint
    ├── Dockerfile
    ├── main.go

---

## ⚙️ Local Development

#### 1. Run service locally

go run main.go

#### 2. Run with Docker

docker build -t hotels-service .
docker run -p 8080:8080 hotels-service

#### 3. Run migrations

cd db
./migrate.sh

---

## 🔗 Service Communication

#### Inside AWS, services communicate via:

- Internal ALB
- Private DNS

#### Example:

- users-service → hotels-service
  http://hotels.internal.selena

---

## 🔐 Security

- No hardcoded secrets
- Private networking only
- IAM roles for access control
- TLS for database connections

---

## 📊 Health Checks

Handled by AWS Application Load Balancer:

- Endpoint: /health

<!--
---

## 🚀 Deployment Flow

- Code pushed to GitHub
- CI/CD builds Docker image
- Image pushed to Amazon ECR
- EC2 instances pull latest image
- The service is updated automatically
-->

---

## ⚠️ Notes

- Service is designed for cloud-first deployment
- Local run is only for development
- Production configuration comes from AWS