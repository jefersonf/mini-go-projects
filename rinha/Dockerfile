FROM golang:1.22

WORKDIR /rinha
COPY . /rinha

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build cmd/main.go

CMD ["./main"]