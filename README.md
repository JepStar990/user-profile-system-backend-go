# User Profile System Backend (Go)

Go microservice managing user profiles, authentication, settings, favorites, downloads, listening history, and content recommendations for the RadioAI platform.

Built with [Fiber](https://gofiber.io), [GORM](https://gorm.io), MySQL 8.0, Prometheus metrics, and OpenTelemetry tracing.

## Prerequisites

| Tool | Minimum Version | Install |
|---|---|---|
| Go | 1.25.0 | `https://go.dev/dl/` |
| Docker | 24+ | `https://docs.docker.com/engine/install/` |
| Docker Compose | v2+ | (bundled with Docker) |
| MySQL | 8.0 | (or use the `docker-compose` MySQL) |

## Quick Start

```bash
# Clone and cd
git clone <repo-url> && cd user-profile-system-backend-go

# Start everything (MySQL + API)
docker compose up -d

# Check it's alive
curl http://localhost:8080/api/health
# → {"status":"ok"}
```

The API will be available at **http://localhost:8080**. MySQL is exposed on **localhost:3306**.

## Running Without Docker

If you already have MySQL running locally, you can build and run the Go binary directly:

```bash
# Copy and edit env
cp .env.example .env
# Update DB_DSN with your MySQL credentials

# Install dependencies
go mod download

# Run
go run ./cmd/server
```

## Testing the API

### 1. Register a user

```bash
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{"email":"user@example.com","username":"test","password":"Secure123!"}'
# → {"message":"User registered successfully","user_id":"..."}
```

### 2. Login (get tokens)

```bash
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"user@example.com","password":"Secure123!"}'
# → {"access_token":"<access>","refresh_token":"<refresh>"}
```

### 3. Access a protected endpoint

```bash
TOKEN="<access_token from login>"

# Get profile
curl http://localhost:8080/api/private/profile \
  -H "Authorization: Bearer $TOKEN"

# Get settings
curl http://localhost:8080/api/private/settings \
  -H "Authorization: Bearer $TOKEN"

# Add a favorite
curl -X POST http://localhost:8080/api/private/favorites \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"content_id":"track-001","content_type":"music"}'
```

### 4. Admin endpoints

```bash
curl http://localhost:8080/api/admin/health \
  -H "X-Admin-Key: supersecretadminapikey"
```

## API Reference

### Public routes (`/api/auth`)

| Method | Path | Description |
|---|---|---|
| POST | `/api/auth/register` | Create new user account |
| POST | `/api/auth/login` | Authenticate, returns JWT pair |
| POST | `/api/auth/refresh` | Refresh access/refresh tokens |
| POST | `/api/auth/logout` | Revoke refresh token |

### Protected routes (`/api/private`) — requires `Authorization: Bearer <token>`

| Method | Path | Description |
|---|---|---|
| GET | `/profile` | Get user profile |
| PUT | `/profile` | Update profile (fullName, bio) |
| POST | `/profile/change-password` | Change password |
| POST | `/profile/avatar` | Upload avatar image |
| GET | `/settings` | Get all settings |
| PUT | `/settings/audio` | Update audio settings |
| PUT | `/settings/voice` | Update voice settings |
| PUT | `/settings/live` | Update live radio settings |
| PUT | `/settings/notifications` | Update notification settings |
| PUT | `/settings/appearance` | Update appearance settings |
| PUT | `/settings/privacy` | Update privacy settings |
| POST | `/favorites` | Add favorite |
| GET | `/favorites` | List favorites |
| DELETE | `/favorites` | Remove favorite |
| POST | `/downloads` | Register download |
| GET | `/downloads` | List downloads |
| DELETE | `/downloads` | Remove download |
| GET | `/downloads/url` | Get presigned S3 URL |
| GET | `/history` | Get listening history |
| POST | `/history/progress` | Update progress on a track |
| DELETE | `/history` | Clear listening history |
| GET | `/history/stats` | Get listening stats |

### Admin routes (`/api/admin`) — requires `X-Admin-Key` header

| Method | Path | Description |
|---|---|---|
| GET | `/admin/health` | System health check |
| GET | `/admin/metrics` | App metrics |
| GET | `/admin/version` | Build version info |

### Monitoring

| Path | Description |
|---|---|
| `GET /metrics` | Prometheus metrics endpoint |
| `GET /api/health` | Public health check |

## Environment Variables

| Variable | Default | Description |
|---|---|---|
| `APP_ENV` | `development` | Environment (development/production) |
| `APP_PORT` | `8080` | API listen port |
| `DB_DSN` | — | MySQL connection string |
| `JWT_ACCESS_SECRET` | — | Secret for signing access tokens |
| `JWT_REFRESH_SECRET` | — | Secret for signing refresh tokens |
| `ACCESS_TOKEN_EXPIRES` | `15m` | Access token TTL |
| `REFRESH_TOKEN_EXPIRES` | `720h` | Refresh token TTL (30 days) |
| `ADMIN_API_KEY` | — | Key for admin endpoints |
| `CORS_ALLOW_ORIGINS` | — | Allowed CORS origins |
| `AWS_REGION` | — | AWS region for S3 |
| `AWS_S3_BUCKET` | — | S3 bucket name |
| `OTEL_EXPORTER_OTLP_ENDPOINT` | — | OTLP traces endpoint |

## Project Structure

```
.
├── cmd/server/main.go          # Entry point
├── internal/
│   ├── controllers/            # HTTP handlers
│   ├── db/                     # Database connection
│   ├── dto/                    # Request/response types
│   ├── models/                 # GORM entities
│   ├── recommendations/        # Content recommendation engine
│   ├── repositories/           # Database queries
│   ├── security/               # JWT, password hashing, auth middleware
│   ├── server/                 # App bootstrap + route setup
│   ├── services/               # Business logic
│   ├── storage/                # S3 uploader
│   ├── telemetry/              # Metrics + tracing middleware
│   └── utils/                  # Validator, error handler, middleware
├── Dockerfile
├── docker-compose.yml
├── go.mod
└── go.sum
```

## Development

```bash
# Run all tests
go test ./...

# Format code
go fmt ./...

# Vet for issues
go vet ./...

# Build
go build -o server ./cmd/server
```
