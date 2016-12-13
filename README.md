# chantest
Test passing channel pointer between c and go routines

build go program with:

go build -buildmode=c-archive -o goclient.a goclient.go

and c binary with:

gcc -o cgo cgo.c ./goclient.a  -pthread
