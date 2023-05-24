from golang as base

run mkdir -p /go/src/

copy . /go/src/go-library-API-pr3

workdir /go/src/go-library-API-pr3/cmd/main

run go build .

cmd "./main"
