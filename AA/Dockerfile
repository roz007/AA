FROM golang:1.13

WORKDIR /app


COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8001
EXPOSE 8002


ENTRYPOINT ["./main"]

CMD ["start-server1"]