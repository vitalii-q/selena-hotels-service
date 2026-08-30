# Current State: Hotels Service

## Local Runtime

- Technology: Go, Gin, GORM, and SQL migrations.
- Development HTTP port: `9064`.
- Development database: CockroachDB, exposed locally on SQL port `9264` and
  reachable in Docker as `hotels-db:26258`.
- The Docker image is built with:

  ```bash
  docker build --no-cache --platform=linux/amd64 -t selena-hotels-service:amd64 .
  ```

- The development container expects the shared Docker network
  `selena-dev_app_network` and a local `.env` file.
- The entrypoint waits for CockroachDB, creates the database/user, applies SQL
  migrations, runs seed data, and starts the service with Air.

## Health and Observability

- Liveness endpoint: `/health`.
- Database readiness endpoint: `/ready`.
- Request ID, logging, and recovery middleware are configured.

## Database and Migrations

- The service owns a CockroachDB database.
- SQL migrations are stored in `db/migrations`.
- Existing schema covers countries, cities, and hotels.

## Existing HTTP API

- `POST /api/v1/hotels`
- `GET /api/v1/hotels`
- `GET /api/v1/hotels/{id}`
- `PUT /api/v1/hotels/{id}`
- `DELETE /api/v1/hotels/{id}`
- `GET /api/v1/locations`

## Tests

- No Go unit or integration test files were found in the repository.
- Intended standard command after tests are introduced: `go test ./...`.

## Known MVP Gaps

- The Room domain model, repository, service layer, and database migration are
  present, but no room HTTP API exists yet.
- There is no availability search by date/guest count.
- There is no room reservation or release API for bookings-service.
- Hotel price currently uses a floating-point field and must not be used as the
  payment source of truth.
- Input validation, error responses, pagination, and sorting need
  standardization.
- Existing Kubernetes manifests are service-local but need review and alignment
  with the planned Minikube structure.

## Working Tree

The repository worktree was clean when this document was created.
