FROM golang:1.14-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /api-vade

EXPOSE 8080

CMD [ "/api-vade" ]