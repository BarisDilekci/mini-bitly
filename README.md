# mini-bitly

**mini-bitly** is a simple, high-performance URL shortener written in Go.

It is designed to:
- be easy to understand
- be easy to run
- be easy to extend

It solves one problem well: **making long URLs short.**

---

## ⚠️ Important Notes (Read First)

- This is an **API-first** application.
- It does **not** serve an HTML homepage.
- Visiting `http://localhost:8080/` will return an error (`Missing code`) — this is **expected behavior**.
- Usage flow:
  1. Create a short URL using `POST /shorten`
  2. Visit the returned short URL in your browser to get redirected

---

## 🚀 How to Run

### Using Docker (Recommended)

You do not need Go installed.

```bash
docker-compose up --build
# mini-bitly
