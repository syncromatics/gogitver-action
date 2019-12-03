FROM golang:1.13.1 as build

WORKDIR /build

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go vet ./...

RUN go get -u golang.org/x/lint/golint

RUN golint -set_exit_status ./...

RUN go build -o action ./

FROM ubuntu:18.04 as final

WORKDIR /app

COPY --from=0 /build/action /app/

CMD ["/app/action"]
