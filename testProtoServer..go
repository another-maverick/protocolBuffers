package main

import (
	"net/http"
	"github.com/golang/protobuf/proto"
	"github.com/another-maverick/protocolBuffers/"

)


func main() {
	http.HandlerFunc("/", testHandler)
	http.ListenAndServe(":12345", nil)
}
