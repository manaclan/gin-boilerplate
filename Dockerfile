# Change go version appropriately
FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY src src
COPY main.go ./

RUN go build -o /docker-gs-ping

EXPOSE 8525

CMD [ "/docker-gs-ping" ]