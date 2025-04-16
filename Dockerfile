FROM oven/bun:latest AS frontend-builder

WORKDIR /app/frontend

# Copy only files needed for Vite install/build
COPY frontend/package*.json ./
RUN bun install

COPY frontend ./
RUN bun run build

# ----------- Build Go Backend with Embed ----------
FROM golang:1.24-alpine AS backend-builder
RUN apk add --no-cache gcc musl-dev
WORKDIR /app

# Install necessary packages
RUN apk add --no-cache git

# Copy go.mod/sum and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy backend code and frontend build
COPY . .
# Make sure the dist folder is present before build
RUN ls -la frontend/dist

# Build the Go app with embedded frontend
RUN CGO_ENABLED=1 GOOS=linux go build -o app .

# ----------- Final Clean Image --------------------
FROM alpine:latest

WORKDIR /app

# Copy binary from builder
COPY --from=backend-builder /app/app .

# Run it
EXPOSE 8080
CMD ["./app"]
