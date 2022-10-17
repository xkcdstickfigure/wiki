FROM golang:1.19
WORKDIR /app
COPY . .
RUN go mod download
CMD go run .