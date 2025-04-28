# golog

A simple, high-performance, SLF4J-style logging utility for Go — with colorized console output, clean file logging, strict placeholder matching, and multiple log levels.

---

## Features

- 📅 Timestamps on every log entry
- 🎨 Colorized output in console (info, warn, debug, error)
- 📝 Clean log files (no ANSI color escape sequences)
- 🔥 Simple `{}` placeholders with strict argument checking
- 🚨 Panic if placeholder `{}` count and argument count mismatch
- 📂 Write logs simultaneously to console and a file
- 📦 Lightweight and no external dependencies
- ✍️ Easy-to-use `Info`, `Warn`, `Debug`, and `Error` methods

---

## Installation

```bash
go get github.com/kashari/golog
```

# Usage
## 1. Initialize the logger

```go
import "github.com/kashari/golog"

func main() {
    err := golog.Init("app.log")
    if err != nil {
        panic(err)
    }
    defer golog.Close()

    golog.Info("Application started with user {} and age {}", "Alice", 30)
    golog.Warn("Memory usage high: {}%", 85)
    golog.Debug("Fetching data for user ID {}", 12345)
    golog.Error("Failed to connect to {}:{}", "database", 5432)
}
```

## 2. Example Output
### Console (with colorized levels)
```bash
2025-04-28 21:00:01 [INFO] Application started with user Alice and age 30
2025-04-28 21:00:02 [WARN] Memory usage high: 85%
2025-04-28 21:00:03 [DEBUG] Fetching data for user ID 12345
2025-04-28 21:00:04 [ERROR] Failed to connect to database:5432
```

Method | Description
---- | ----
golog.Init(path) | Initialize logger with a file path
golog.Close() | Close the logger file safely
golog.Info(msg, args...) | Log an INFO level message
golog.Warn(msg, args...) | Log a WARN level message
golog.Debug(msg, args...) | Log a DEBUG level message
golog.Error(msg, args...) | Log an ERROR level message