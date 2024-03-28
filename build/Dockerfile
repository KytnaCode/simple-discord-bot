FROM golang:1.22.1-bookworm
WORKDIR /app

# All host paths are relative to project's root.

# Install `air` for live reloading.
RUN go install github.com/cosmtrek/air@latest

# Dowloand go modules.

# We wont copy `go.sum` along with `go.mod` since it isn't generated because we haven't installed modules yet.
# TODO: Copy `go.sum` when modules are installed.
COPY go.mod ./
RUN go mod download

# Copy files
COPY . .

# Start `air`.
CMD ["air", "-c", ".air.toml"]
