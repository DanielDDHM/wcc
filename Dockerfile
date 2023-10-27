FROM golang:1.18-buster as builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy local code to the container image.
COPY . ./

# Build the binary.
RUN go build -o app

# Use the official Debian slim image for a lean production container.
# https://hub.docker.com/_/debian
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM debian:buster-slim
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/app /app/app
CMD chmod +x /app/app

COPY --from=builder /app/placeholder.env /.env

# Run the web service on container startup.
ENTRYPOINT ["/app/app"]