# Multi stage build for the route portal

# Frontend builder stage
FROM node:18-alpine AS frontend-builder
# Set working directory for frontend
WORKDIR /app/frontend
# Copy package json files
COPY frontend/package*.json ./
# Install dependencies
RUN npm ci
# Copy frontend source
COPY frontend/ ./
# Build the frontend
RUN npm run build

# Backend builder stage
FROM golang:1.24-alpine AS backend-builder
# Set working directory for backend
WORKDIR /app/backend
# Copy go mod files
COPY backend/go.mod backend/go.sum ./
# Download dependencies
RUN go mod download
# Copy backend source
COPY backend/ ./
# Build the backend binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/api

# Final stage - runtime
FROM alpine:latest
# Install nodejs for runtime
RUN apk --no-cache add nodejs npm
# Set workdir
WORKDIR /app

# Now Create dir for yaml storage
RUN mkdir -p /.data
# Copy backend from builder
COPY --from=backend-builder /app/backend/main ./backend

# Now copy the frontend from builder
COPY --from=frontend-builder /app/frontend/.next ./frontend/.next
COPY --from=frontend-builder /app/frontend/public ./frontend/public
COPY --from=frontend-builder /app/frontend/package*.json ./frontend/
COPY --from=frontend-builder /app/frontend/next.config.* ./frontend/

# Install production dependencies for the nextJs app
WORKDIR /app/frontend
RUN npm ci --only=production && npm cache clean --force

# Back to app to start the whole app
WORKDIR /app
# Expose ports
EXPOSE 8080 3000

# Create startup script
RUN echo '#!/bin/sh' > start.sh && \
    echo 'cd /app && ./backend &' >> start.sh && \
    echo 'cd /app/frontend && npm run start' >> start.sh && \
    chmod +x start.sh

# Start the app
CMD ["./start.sh"]
