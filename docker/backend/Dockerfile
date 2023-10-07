FROM golang:1.20-alpine

WORKDIR ./backend

COPY ./backend/go.mod ./
COPY ./backend/go.sum ./

RUN go mod tidy
RUN go mod download


COPY . .

RUN go build -o backend ./cmd/app/main.go

ENV GIN_MODE release

EXPOSE 8090

CMD [ "./backend" ]