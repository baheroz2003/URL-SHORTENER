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
------------to stop---------
docker-compose down
----------------------------
🧭 ORDER TO UNDERSTAND THE FLOW
main.go
routes/ folder
shorten.go
resolve.go
helpers/helpers.go
database/database.go
go.mod and go.sum
Dockerfile
docker-compose.yml
README.md
db/ folder (if used)
🔹 Might contain schema logic or data dumps for Redis, or be empty.


