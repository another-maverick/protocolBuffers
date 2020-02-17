package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	pb "github.com/another-maverick/protocolBuffers/testProtocolBuffer"
	"io/ioutil"
	"net/http"
	"os"
)


func main() {
	res, err := http.Get("http://localhost:12345/")
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}

	var myUser pb.User

	err = proto.UnmarshalMerge(b, &myUser)
	if err != nil{
		fmt.Println("Cannot read the body into the struct", err)
		os.Exit(1)
	}

	fmt.Println(myUser.GetFirstname())
	fmt.Println(myUser.GetLastname())
	fmt.Println(myUser.GetId())
	fmt.Println(myUser.GetCity())



}
