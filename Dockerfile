FROM golang:1.17.9-alpine3.15

ENV ROOT=/go/src/app

WORKDIR ${ROOT}

RUN apk update && apk add git

COPY cmd/main.go cmd/main.go
COPY go.mod .

RUN go mod tidy

CMD ["go", "run", "cmd/main.go"]
