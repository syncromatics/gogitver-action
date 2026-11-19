FROM golang:1.26.3 AS build

WORKDIR /build

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go vet ./...

RUN go install golang.org/x/lint/golint@v0.0.0-20190930215403-16217165b5de

RUN golint -set_exit_status ./...

RUN go build -o action ./

CMD ["/build/action"]
