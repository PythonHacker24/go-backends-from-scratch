FROM golang:1.24.2

WORKDIR /app

COPY go.mod ./
COPY main.go ./

RUN go mod tidy 
RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
