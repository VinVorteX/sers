FROM golang:1.24-alpine

WORKDIR /app

# Add DNS settings to resolve connectivity issues
RUN echo "nameserver 8.8.8.8" > /etc/resolv.conf && \
    echo "nameserver 8.8.4.4" >> /etc/resolv.conf

# Install required packages
RUN apk add --no-cache \
    gcc \
    musl-dev \
    git

# Set mirror for better download speeds
ENV GOPROXY=https://goproxy.cn,direct

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
