package main 

import (
	"fmt"
	"gopkg.in/mgo.v2"
)

var dbSession *mgo.Session

func NewDBConnection() {
	url := "localhost"
	
	var err error

	dbSession, err = mgo.Dial(url)

	if err != nil {
		fmt.Println("Unable to establish a connection to MongoDB at", url)
		panic(err)
	} else {
		fmt.Println("Connection established to MongoDB at", url)
	}
}