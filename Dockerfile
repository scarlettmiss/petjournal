# Step 1: Build the Next.js frontend
FROM node:18 as frontend

WORKDIR /app

# Copy the entire application code
COPY ui .

RUN yarn install

# Build the Next.js application and output to the "build" directory
RUN yarn build

FROM golang:1.21.3 as builder

# Set the working directory to /go/src/app
WORKDIR /go/src/app

# Copy the local package files to the container's workspace
COPY server .

# Copy the "build" directory from the frontend build
COPY --from=frontend /app/build ./cmd/server/public

# Set environment variables
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

# Build the Go application
RUN go build -o petJournal ./cmd/server

# Start a new stage to reduce the final image size
FROM alpine:latest

# Set the working directory to /app
WORKDIR /app

# Copy the binary from the builder stage to the final stage
COPY --from=builder /go/src/app/petJournal .

# Set environment variable for your application
ENV SECRET_KEY=$SECRET_KEY
ENV DB_URL=$DB_URL
ENV DB_NAME=$DB_NAME


# Expose port 8080 for the application
EXPOSE 8080

# Define the command to run the application
ENTRYPOINT ["./petJournal"]
