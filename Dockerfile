FROM golang:1.21.0
WORKDIR /app
# Download Go modules
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /workflow-run
EXPOSE 8000
# Run
CMD ["/workflow-run"]