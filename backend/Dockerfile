FROM golang:1.24-alpine

WORKDIR /app

# Install required packages
RUN apt update && apt install -y \
    gcc \
    musl-dev \
    git \
    && rm -rf /var/lib/apt/lists/*  # Clean up

# Set mirror for better download speeds (optional)
ENV GOPROXY=https://proxy.golang.org,direct

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main .

EXPOSE 8080
CMD ["./main"]
