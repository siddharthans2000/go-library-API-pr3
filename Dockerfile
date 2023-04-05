from golang as base

run mkdir -p /go/src/go-library-API-pr3

copy . /go/src/go-library-API-pr3

workdir /go/src/go-library-API-pr3/cmd/main

run go build .

from mysql

copy --from=base /go/src/go-library-API-pr3 /go/src/go-library-API-pr3

workdir /go/src/go-library-API-pr3/cmd/main
