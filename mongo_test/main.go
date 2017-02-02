package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type Student struct {
	Name   string
	School float64
	Age    float64
}

type A struct {
	Zipcode  string
	Students []Student
}

func main() {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("t")
	var data []bson.M
	err = c.Find(bson.M{}).All(&data)
	if err != nil {
		log.Fatal(err)
	}

	// Print all data in mongo -> test -> t
	for _, doc := range data {
		for key, value := range doc {
			fmt.Println(key, value)
		}
	}

	var a A
	err = c.Find(bson.M{"_id": 1}).One(&a)
	if err != nil {
		fmt.Println("err1", err)
	} else {
		fmt.Printf("%+v\n", a)
	}

	err = c.Find(bson.M{"_id": "tt"}).One(&a)
	if err != nil {
		fmt.Println("err2", err)
		fmt.Printf("%+v\n", err)
	} else {
		fmt.Printf("%+v\n", a)
	}

	var aAry []A
	err = c.Find(bson.M{"_id": 1}).All(&aAry)
	if err != nil {
		fmt.Println("err3", err)
	} else {
		fmt.Printf("%+v\n", aAry)
	}
	err = c.Find(bson.M{"_id": "tt"}).All(&aAry)
	if err != nil {
		fmt.Println("err4", err)
	} else {
		fmt.Printf("%+v\n", aAry)
	}
}
