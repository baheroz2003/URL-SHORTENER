🚀 URL Shortener API
A simple URL shortener API built with Go, Fiber, Redis, and Docker. It shortens long URLs and redirects short links back to the original URL.

🔧 Tech Stack
🧠 Go – backend logic

⚡ Fiber – web framework

🐳 Docker & Docker Compose – containerization

🔴 Redis – fast key-value store for short URL mappings

🌱 Go Modules – dependency management

📦 Features
✅ Shorten any valid URL

✅ Redirect from short link to original URL

✅ Redis TTL support (links expire after 24h)

✅ Dockerized for easy setup

-----------TO RUN-----------
type docker-compose up -d
------------to stop---
docker-compose down

-------------------------------------------------FLOW-------------------------------------------------------------------------------------------------------------
🧭 ORDER TO UNDERSTAND THE FLOW
main.go
🔹 Entry point of the application.
🔹 Sets up the server, middleware, and routes.

routes/ folder
🔹 Check how routes like /shorten and /resolve/:shortURL are registered.
🔹 Links HTTP endpoints to logic functions.

shorten.go
🔹 Handles the creation of short URLs.
🔹 Includes logic for generating, validating, and storing shortened URLs.

resolve.go
🔹 Handles resolving short URLs back to original ones.
🔹 Includes logic for reading from Redis and redirecting.

helpers/helpers.go
🔹 Contains reusable utility functions.
🔹 Example: URL validators, random string generators, etc.

database/database.go
🔹 Redis connection logic.
🔹 See how Redis is initialized, accessed, and used by the app.

go.mod and go.sum
🔹 See which Go modules and libraries are used (like fiber, redis, etc).

Dockerfile
🔹 Defines how the Go app is built inside a Docker image.

docker-compose.yml
🔹 Brings up both the Go app and Redis together.
🔹 Defines networking and service dependencies.

README.md
🔹 Often provides usage, setup instructions, and examples.
🔹 Follow it for running and testing.

db/ folder (if used)
🔹 Might contain schema logic or data dumps for Redis, or be empty.


