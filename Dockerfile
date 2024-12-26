FROM golang:1.23.4 AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go install github.com/swaggo/swag/cmd/swag@latest && \
    swag init -g ./main.go -o ./api
    
RUN go build -o /build/app .

FROM golang:1.23.4 as final

RUN mkdir -p /app

COPY --from=builder /build/app /app

CMD ["/app/app"]
