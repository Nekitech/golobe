FROM golang:1.20-alpine

WORKDIR /backend

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o backend

ENV GIN_MODE release

EXPOSE 8090

CMD [ "./backend" ]