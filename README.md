# mini-bitly

**mini-bitly** is a simple, high-performance URL shortener written in Go.

It is designed to make long URLs manageable, short, and shareable.

## 🌟 Features

*   **RESTful API**: Simple and standard-compliant HTTP API.
*   **Fast**: Uses an in-memory database for millisecond-response times.
*   **Lightweight**: Runs with minimal dependencies.
*   **Docker Support**: Containerized architecture for easy deployment anywhere.
*   **Clean Architecture**: Modular structure with Handler, Service, and Repository layers.

---

## 🛠 Installation & Running

There are two main ways to run the project: using Docker (recommended) or manually with Go.

### 1. Using Docker (Recommended)

You do not need Go installed on your machine. Docker is sufficient.

```bash
docker-compose up --build
```

This command builds the application and starts it on port `:8080`.

### 2. Manual Execution with Go

If you have Go (1.25 or later) installed:

```bash
# Download dependencies
go mod download

# Start the application
go run cmd/server/main.go
```

The server will start at: `http://localhost:8080`

---

## 🚀 Usage (API Documentation)

**Note:** This project does not provide a web interface (HTML). It is purely API-based. Visiting `http://localhost:8080` in a browser will result in an error. You should use `curl`, `Postman`, or similar tools for interaction.

### 🔗 1. Shorten URL

Used to create a new short link.

*   **Endpoint:** `POST /shorten`
*   **Content-Type:** `application/json`

**Example Request:**

```bash
curl -X POST http://localhost:8080/shorten \
     -H "Content-Type: application/json" \
     -d '{"url": "https://www.google.com/search?q=golang"}'
```

**Success Response (200 OK):**

```json
{
  "short_url": "http://localhost:8080/a1b2c3"
}
```

### 🔀 2. Redirect (Open Short Link)

Redirects the short link to the original address.

*   **Endpoint:** `GET /:code`

**Example Usage:**

Paste the `short_url` you received in the previous step into your browser or test with curl:

```bash
# -I flag shows headers only (to see the redirect)
curl -I http://localhost:8080/a1b2c3
```

**Expected Result:**
The server returns a `302 Found` status code with the `Location` header set to the original URL, redirecting you there.

---

## 📂 Project Structure

```
mini-bitly/
├── cmd/
│   └── server/       # Application entry point (main.go)
├── internal/
│   ├── handler/      # HTTP request handling layer
│   ├── service/      # Business Logic
│   ├── repository/   # Data storage layer (In-Memory storage)
│   └── domain/       # Data models and errors
├── Dockerfile        # Docker image file
├── docker-compose.yml
└── README.md
```
