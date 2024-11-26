ARG GO_VERSION=1.22.2-bookworm
FROM golang:${GO_VERSION}

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy the rest of the required files
COPY src src
COPY internal internal
COPY main.go main.go

CMD go run main.go
