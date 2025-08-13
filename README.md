# Weather API Wrapper

A simple efficient wrapper for a third-party weather API includes caching layer with Redis.

---

## Features

- **External API Integration:** Connects to the Visual Crossing weather API to fetch real-time data.
- **Caching Strategy:** Implements a cache-aside pattern with Redis to serve recent requests faster.
- **Simplified Endpoint:** Provides a simple `GET /weather/{city}` endpoint.
- **API Key Security:** Hides the sensitive third-party API key on the backend, never exposing it to the client.

---

## Tech Stack

- **Language**: Go
- **Web Framework**: Fiber
- **Cache**: Redis (w/ `github.com/redis/go-redis/v9`)
- **Dependencies**: Docker
- **Environment Config**: `.env` w/ `github.com/joho/godotenv`

---

## Folder Structure

```
weather-api-wrapper/
├── cache/          # Redis connection & logic
├── client/         # External API client logic
├── controller/     # HTTP request handlers
├── model/          # Data models (structs)
├── service/        # Business logic
├── .env            # config file
├── .gitignore      # place for ignored file
├── go.mod          # go module
├── go.sum
├── main.go         # Main func (router, dependency injection)
├── openapi.json    # API spec documentation
└── README.md       # Project info
```

---

## Getting Started

### 1. Clone the repo
```bash
git clone https://github.com/pisondev/weather-api-wrapper.git
cd weather-api-wrapper
```

### 2. Create a `.env` file
Create a file named `.env` in the root of the project and add your credentials.

```env
WEATHER_API_KEY="YOUR_VISUAL_CROSSING_API_KEY"
REDIS_ADDRESS="localhost:6379"
REDIS_PASSWORD=""
REDIS_DB=0
```

> ⚠️ Don't commit your `.env` file — it should be in your `.gitignore`.

### 3. Start Dependencies
This project requires a running Redis instance. Use Docker to start one easily.
```bash
docker run -d --name redis-weather -p 6379:6379 redis:latest
```
or
```bash
docker create --name redis-weather -p 6379:6379 redis:latest
docker start redis-weather
```

### 4. Install Go Modules
```bash
go mod tidy
```

### 5. Run the app
```bash
go run main.go
```

## API Endpoints

| Method | Endpoint | Description |
|---|---|---|
| GET | `/weather/{city}` | Get current weather by city. Results are cached for 1 hour. |

> The API is documented in more detail in the `openapi.json` file.

---

## TODO / Roadmap

- [ ] **Advanced Caching:** Implement a "stale-while-revalidate" or "serve stale on error" strategy to improve reliability if the external API is down.
- [ ] **Concurrent Fan-Out:** Add an endpoint like `/weather?cities=jakarta,london` to fetch data for multiple cities in parallel using goroutines.
- [ ] **Structured Error Handling:** Refine the `client` package to return specific error types (e.g., `ErrCityNotFound`) for more precise error responses.
- [ ] **Tests and CI/CD:** Add unit and integration tests for the service and handler layers.
- [ ] **Deployment:** Containerize the Go application with Docker for easy deployment.

---

## License

MIT — free to use and modify.