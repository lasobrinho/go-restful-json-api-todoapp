package main 

import (
	"time"
	"gopkg.in/mgo.v2/bson"
)

type Todo struct {
	Id			bson.ObjectId	`json:"id" bson:"_id,omitempty"`
	Name 		string			`json:"name"`
	Completed 	bool			`json:"completed"`
	Due			time.Time 		`json:"due"`
}
type Todos []Todo

func NewTodo(name string, completed bool, due time.Time) Todo {
	return Todo {
		Id:			bson.NewObjectId(),
		Name:		name,
		Completed:	completed,
		Due:		due,
	}
}