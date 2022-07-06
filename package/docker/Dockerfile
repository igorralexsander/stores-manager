FROM golang:1.18-alpine as builder

WORKDIR /build

# Git must be installed because go mod requires it to download dependencies
# After its used we may remove it to reduce docker image size
RUN apk --no-cache add git tzdata gcc libc-dev

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -tags musl -o stores-manager cmd/stores-manager/main.go

# Execution container
FROM alpine
RUN apk --no-cache add tzdata
WORKDIR /app
USER 1000
COPY --from=builder /build/stores-manager /app/
COPY ../../config.yml ./
EXPOSE 8080
CMD ["/app/stores-manager"]