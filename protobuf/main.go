package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/golang/protobuf/proto"
)

func main() {
	test := &Test{
		Label: "hello",
		Type:  17,
		Reps:  []int64{1, 2, 3},
	}
	data, err := proto.Marshal(test)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	b, err := json.Marshal(test)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v\n", test)
	fmt.Printf("Protobuf: %d\n%+v\n", len(data), data)
	fmt.Printf("json: %d\n%+v\n", len(b), b)

	newTest := &Test{}
	err = proto.Unmarshal(data, newTest)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}
	fmt.Printf("%+v\n", newTest)
}
