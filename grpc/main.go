package main

import (
	"fmt"
	"log"
	"github.com/golang/protobuf/proto"
)

func main() {

	elliot:=&Person{
		Name:"Elliot",
		Age:24,
		Id:01,
		Add:"chenai",
	}

	data, err := proto.Marshal(elliot)
	if err != nil {
		log.Fatal("marshalling error:", err)
	}

	fmt.Println(data)

	newElliot := &Person{}

	err = proto.Unmarshal(data, newElliot)
	if err !=nil {
		log.Fatal("Unmarshalling Error:", err)

	}

	fmt.Println(newElliot.GetAge())
	fmt.Println(newElliot.GetName())
	fmt.Println(newElliot.GetId())
	fmt.Println(newElliot.GetAdd())

}