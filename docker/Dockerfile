FROM golang:1.20

WORKDIR ./golobe

COPY ./go.mod ./
COPY ./go.sum ./

RUN go mod download

COPY . .

RUN go build -o cmd/golobe ./cmd/main.go

ENV GIN_MODE release

EXPOSE 8090

CMD [ "./cmd/golobe" ]