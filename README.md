# Netwatcher

Netwatcher is a lightweight Go-based tool for periodically monitoring a list of URLs using a custom worker pool pattern with graceful shutdown support.

## Description
Netwatcher is a pet project designed to demonstrate Go concurrency patterns, such as worker pools, channels, and graceful shutdowns, while following clean architecture principles.

A list of URLs is defined in jobs/generator.go. Jobs are periodically generated and pushed into a worker pool. The worker pool concurrently processes the jobs (HTTP GET requests).
Results are sent into a result channel. A result processor prints the outcome of each check (status code, response time, errors).

## Usage
### Run the application
```bash
go run ./cmd/netwatcher
```
### Sample output
```bash
[SUCCESS] - [https://github.com] - Status: 200, Response Time: 204.818706ms
[SUCCESS] - [https://google.com] - Status: 200, Response Time: 464.531915ms
[ERROR] - [https://golang.org] - Get "https://golang.org": context deadline exceeded (Client.Timeout exceeded while awaiting headers)
[SUCCESS] - [https://api.coincap.io/v2/assets] - Status: 429, Response Time: 1.988438336s
[SUCCESS] - [https://github.com] - Status: 200, Response Time: 63.708308ms
[SUCCESS] - [https://google.com] - Status: 200, Response Time: 174.389ms
Shutdown signal received...
2025/03/23 14:42:56 [worker ID 1] finished proccesing
[SUCCESS] - [https://api.coincap.io/v2/assets] - Status: 200, Response Time: 100.638527ms
2025/03/23 14:42:57 [worker ID 2] finished proccesing
2025/03/23 14:42:57 [worker ID 0] finished proccesing
[SUCCESS] - [https://golang.org] - Status: 200, Response Time: 1.508642526s
Graceful shutdown complete
```
## Features
- Periodic HTTP checks for a configurable list of URLs.
- Concurrent requests handling using a worker pool.
- Asynchronous result processing with logging.
- Graceful shutdown on `SIGINT` / `SIGTERM` signals.
- Clean architecture with separated packages: `jobs`, `pool`, `processor`, and `models`.

## Dependencies
Go 1.20+
Standard Go library (net/http, encoding/json, etc.)
