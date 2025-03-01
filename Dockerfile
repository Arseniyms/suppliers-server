FROM golang:1.24

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download 

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

CMD ["./main"]

EXPOSE 9090