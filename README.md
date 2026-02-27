# 🔧 steam-api - Backend Microservice

Go microservice for Steam API integration, game comparison, and Redis caching.

**Part of:** [Steam Games Comparison App](../README.md)

## 📦 Service Overview

- **Language:** Go 1.24
- **Framework:** Gin Web Framework
- **Cache:** Redis (optional, for production)
- **Port:** 8080
- **Health Check:** GET `/health`

## 🚀 Quick Start

### Prerequisites
- Go 1.24+
- Redis (optional, but recommended)
- Steam API key from [steamcommunity.com/dev](https://steamcommunity.com/dev/registerkey)

### Run Locally

```bash
# Setup
cp .env.example .env
# Edit .env with your STEAM_API_KEY

go mod download

# Start server
make web
# Available at http://localhost:8080

# Health check
curl http://localhost:8080/health
```

### Run with Docker

```bash
# Build
docker build -t steam-api:latest .

# Run standalone (no cache)
docker run -p 8080:8080 \
  -e STEAM_API_KEY=your_key_here \
  steam-api:latest

# Run with Redis
docker-compose up -d
```

### Run with full stack
```bash
# From root directory
docker-compose up -d steam-api
```

## 📚 API Endpoints

### Health Check
```
GET /health
```
**Response:**
```json
{
  "status": "ok"
}
```

### Compare User Libraries
```
GET /user?user_id_1=<steamid64>&user_id_2=<steamid64>
```

**Parameters:**
- `user_id_1` - First player's 64-bit Steam ID
- `user_id_2` - Second player's 64-bit Steam ID

**Response:**
```json
{
  "user1": {
    "steamId": "76561198...",
    "personaName": "Player Name",
    "profileUrl": "..."
  },
  "user2": { ... },
  "shared_games": [
    {
      "appId": 570,
      "name": "Dota 2",
      "playtime": 120,
      "tags": ["moba", "free-to-play"],
      "isCoOp": false,
      "isMultiplayer": true
    }
  ]
}
```

**Caching:**
- Results cached for **1 hour**
- Cache Key: `comparison:<user_id_1>:<user_id_2>`
- Automatic hit/miss handling

## 🏗️ Architecture

```
Request
   ↓
HTTP Handler (CORS enabled)
   ↓
Redis Cache (check)
   ↓ (miss)
Steam API Client
   ↓
Steam Service Layer
   ↓
Game Comparator
   ↓
Redis Cache (store)
   ↓
Response (JSON)
```

## 🔄 Development

### Common Commands
```bash
make help              # Show all commands
make web              # Start dev server
make test             # Run unit tests
make test-coverage    # Generate coverage report
make lint             # Run linters (golangci-lint)
make fmt              # Check formatting
make fmt-fix          # Auto-fix formatting
make security         # Run security scan (gosec)
make ci               # Run all checks
make docker-build     # Build Docker image
make docker-run       # Start with docker-compose
make clean            # Clean up artifacts
```

### Project Structure
```
steam-api/
├── cmd/
│   ├── compare_steam_libs/    # CLI comparison tool
│   ├── get_steam_api_list/    # List Steam APIs
│   ├── list_steam_apps/       # List Steam apps
│   └── web/
│       └── main.go            # HTTP server entry
├── internal/
│   ├── cache/
│   │   └── redis.go           # Redis caching layer
│   ├── config/
│   │   └── config.go          # Configuration management
│   ├── httpserver/
│   │   ├── server.go          # Server setup + CORS
│   │   └── handlers/
│   │       ├── get_user_info.go      # Compare endpoint
│   │       └── health.go             # Health check
│   ├── steamclient/           # Direct Steam API client
│   ├── steamservice/          # Business logic layer
│   └── steamgamecomparator/   # Comparison logic
├── .golangci.yml              # Linting config
├── Dockerfile                 # Multi-stage build
├── docker-compose.yml         # Local compose
├── go.mod / go.sum           # Dependencies
└── makefile                  # Development commands
```

## 🧪 Testing

### Run Tests
```bash
make test                 # Run all tests verbosely
make test-coverage        # Generate HTML coverage report
```

### Coverage Report
```bash
make test-coverage
open coverage.html        # macOS
xdg-open coverage.html   # Linux
start coverage.html      # Windows
```

### Write Tests
```go
// internal/cache/redis_test.go
package cache_test

import (
    "context"
    "testing"
)

func TestRedisCache_Get(t *testing.T) {
    // Test logic
}
```

## 🔍 Linting & Code Quality

### Code Linting
```bash
make lint                 # Check with golangci-lint
make fmt                  # Check formatting
make fmt-fix              # Auto-fix formatting
make vet                  # Run go vet
make security             # Run gosec security scanner
```

### Configuration
- **golangci-lint:** `.golangci.yml`
- **go fmt:** Built-in (no config needed)
- **go vet:** Built-in (no config needed)
- **gosec:** Built-in (no config needed)

## 🔐 Security

### Built-in Security Features
1. **CORS Middleware** - Allows frontend communication securely
2. **Redis Connection** - Graceful degradation if cache unavailable
3. **Error Handling** - No sensitive info leaked in responses
4. **Input Validation** - Steam IDs validated

### Security Scanning
```bash
make security             # Run gosec
# Checks for:
# - Hardcoded credentials
# - Weak cryptography
# - SQL injection risks
# - Memory/concurrency issues
```

## 📊 Performance

### Caching Benefits
- **Without cache:** ~2-3 seconds per request (Steam API call)
- **With cache:** ~50-100ms per request (Redis hit)
- **TTL:** 1 hour per comparison

### Optimization Tips
1. Redis should run on same network for best latency
2. Connection pooling handled by redis client
3. Multi-stage Docker build keeps image small (~150MB)

## 🚢 Deployment

### Docker Image
```bash
docker build -t steam-api:latest .
```

**Image size:** ~150MB (optimized multi-stage build)

### Run Standalone
```bash
docker run -p 8080:8080 \
  -e STEAM_API_KEY=your_key \
  -e REDIS_URL=redis://localhost:6379 \
  -e API_HOST=0.0.0.0 \
  -e API_PORT=8080 \
  steam-api:latest
```

### Run with Docker Compose (includes Redis)
```bash
docker-compose up -d steam-api
```

### Kubernetes Ready
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: steam-api
spec:
  containers:
  - name: steam-api
    image: steam-api:latest
    ports:
    - containerPort: 8080
    env:
    - name: STEAM_API_KEY
      valueFrom:
        secretKeyRef:
          name: steam-secrets
          key: api-key
    livenessProbe:
      httpGet:
        path: /health
        port: 8080
      initialDelaySeconds: 5
      periodSeconds: 10
```

## 🔗 Environment Variables

| Variable | Required | Default | Description |
|----------|----------|---------|-------------|
| STEAM_API_KEY | ✅ Yes | - | Steam Web API key |
| TEST_USER_1 | ❌ No | - | Test user Steam ID |
| TEST_USER_2 | ❌ No | - | Test user Steam ID |
| REDIS_URL | ❌ No | redis://localhost:6379 | Redis connection URL |
| API_HOST | ❌ No | 0.0.0.0 | Server host |
| API_PORT | ❌ No | 8080 | Server port |

## 📝 Configuration

### .env Example
```env
STEAM_API_KEY=your_api_key_here
REDIS_URL=redis://localhost:6379
API_HOST=0.0.0.0
API_PORT=8080
TEST_USER_1=76561198123456789
TEST_USER_2=76561198987654321
```

## ✅ Health Checks

### Docker Health Check
```
GET /health
Expected: 200 OK with { "status": "ok" }
Interval: 30 seconds
Timeout: 10 seconds
```

### Manual Test
```bash
curl -v http://localhost:8080/health
```

## 🐛 Troubleshooting

### Redis Connection Failed
```
Warning: Redis cache connection failed (continuing without cache)
```
✅ **Normal** - App works without cache, just slower
- **Fix:** Ensure Redis is running if you want caching
- **Check:** `docker-compose ps`

### Steam API Errors
- **"401 Unauthorized"** - Invalid STEAM_API_KEY
- **"Invalid user"** - Steam ID not found or profile private
- **"Rate limited"** - Too many API calls (Steam throttling)

**Solutions:**
1. Verify STEAM_API_KEY in .env
2. Use 64-bit Steam IDs (not vanity URLs)
3. Profiles must be public
4. Redis helps reduce rate limiting by caching results

### Port Already in Use
```bash
# Find what's using port 8080
lsof -i :8080
# Kill the process or use different port
API_PORT=9090 make web
```

## 📖 Dependencies

### Direct
- `github.com/gin-gonic/gin` - Web framework
- `github.com/joho/godotenv` - .env loading
- `github.com/redis/go-redis` - Redis client

### Transitive
- Checked in `go.sum`
- Updated via `go get -u ./...`

## 🚀 Next Steps

- Read [main README](../README.md) for full stack setup
- Check [CI_CD_GUIDE.md](../CI_CD_GUIDE.md) for testing & security
- Deploy frontend: [steam-app README](../steam-app/README.md)

## 📞 Support

- Backend issues? Check logs: `docker logs steam-api`
- API contract questions? See [API Endpoints](#-api-endpoints) above
- CI/CD help? Read [CI_CD_GUIDE.md](../CI_CD_GUIDE.md)

## 📄 License

MIT
