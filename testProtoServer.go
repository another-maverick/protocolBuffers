package main

import (
	"net/http"
	"github.com/golang/protobuf/proto"
	pb "github.com/another-maverick/protocolBuffers"

)


func main() {
	http.HandleFunc("/", testHandler)
	http.ListenAndServe(":12345", nil)
}

func testHandler(res http.ResponseWriter, req *http.Request) {
	u := &pb.User{
		Firstname: proto.String("Steph"),
		Lastname: proto.String("Curry"),
		Id: proto.Int32(30),
		City: proto.String("Oakland"),
	}

	body, err := proto.Marshal(u)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	res.Header().Set("Content-Type", "application/x-protobuf")
	res.Write(body)
}

